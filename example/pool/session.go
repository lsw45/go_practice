package pool

import (
	"encoding/json"
	"fmt"
	"github.com/CardInfoLink/tasty/models"
	"github.com/CardInfoLink/tasty/redis"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// Session 用户的 websocket 连接信息
type Session struct {
	sync.Mutex
	ShopID        string          `json:"shopId"`    // 门店号
	TableNo       string          `json:"tableNo"`   // 桌台号
	UserID        string          `json:"userID"`    // 用户id
	PayMode       string          `json:"payMode"`   // 支付模式
	RandomKey     string          `json:"randomKey"` // 随机数 当同一个用户的连接断开并很快重连时, 可以用这个随机数来区分连接
	Conn          *websocket.Conn `json:"-"`         // 用户的 websocket 连接信息
	sendBuf       chan []byte     // writePump 从此 chan 中, 读取消息, 发送给前端
	bufClosed     bool
	close         chan struct{}
	readCloseFunc func(*Session)
}

func NewSession(shopId string, tableNo string, userId string, payMode string, randomKey string, conn *websocket.Conn) *Session {
	return &Session{
		ShopID:    shopId,
		TableNo:   tableNo,
		UserID:    userId,
		PayMode:   payMode,
		RandomKey: randomKey,
		Conn:      conn,
		sendBuf:   make(chan []byte, 1024),
		close:     make(chan struct{}),
	}
}

func (uc *Session) SetReadCloseFunc(f func(session *Session)) {
	uc.readCloseFunc = f
}

func (uc *Session) UserKey() string {
	return uc.UserID + "_" + uc.RandomKey
}

func (uc *Session) TableKey() string {
	return uc.ShopID + "_" + uc.TableNo
}

func (uc *Session) CloseBuf() {
	uc.Lock()
	if uc.bufClosed {
		uc.Unlock()
		return
	}
	uc.bufClosed = true
	close(uc.sendBuf)
	uc.Unlock()
}

func (uc *Session) SendMessageToChan(data []byte) bool {
	if data == nil {
		return true
	}
	uc.Lock()
	if uc.bufClosed {
		uc.Unlock()
		return false
	}
	select {
	case uc.sendBuf <- data:
	default:
		uc.CloseBuf()
		uc.Unlock()
		return false
	}
	uc.Unlock()
	return true
}

var pongWait = time.Duration(10)
var pingPeriod = (pongWait * 9) / 10

// ReadPump 用户 websocket 连接的读取函数
func (uc *Session) ReadPump() {
	defer func() {
		logrus.Debugf("websocket user disconnecting from read pump, session: %s", uc)
		DefaultSessionPool.Remove(uc)
		uc.CloseBuf()
		uc.Conn.Close()
		close(uc.close)
		if uc.readCloseFunc != nil {
			uc.readCloseFunc(uc)
		}
	}()

	// 设置此次读取的过期时间
	if err := uc.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		logrus.Errorf("websocket set write deadline error: %v, session: %s", err, uc)
		return
	}
	uc.Conn.SetReadLimit(1024)
	uc.Conn.SetPongHandler(func(string) error { uc.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, reqMsg, err := uc.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				logrus.Warnf("websocket closed unexpected error: %v, session: %s", err, uc)
			}
			break
		}

		logrus.Infof("receive websocket message, session: %s, message: %s", uc, string(reqMsg))
		var wsReq models.WsMessage
		if err = json.Unmarshal(reqMsg, &wsReq); err != nil {
			logrus.Warnf("websocket unmarshal message error: %v, sesion: %s request message is %v", err, uc, string(reqMsg))
			continue
		}

		var replyMsg string
		switch wsReq.Type {
		case "MsgNeedPush":
			var bs []*models.Basket
			bs, err = redis.BasketCache.GetAll(uc.ShopID, uc.TableNo, nil)
			if err != nil {
				logrus.Warnf("get all basket cache error: %v, shpID is %v, tableNo is %v", err, uc.ShopID, uc.TableNo)
				continue
			}
			var bsByte []byte
			bsByte, err = json.Marshal(bs)
			if err != nil {
				logrus.Warnf("marshal json error: %v, shpID is %v, tableNo is %v, baskets are %+v", err, uc.ShopID, uc.TableNo, bs)
				continue
			}
			fmt.Println(bsByte)
			continue
		default:
			logrus.Warnf("websocket unknown message type error, session: %s, request message is %v", uc, string(reqMsg))
			replyMsg = "unknown message type"
		}

		if replyMsg == "" {
			replyMsg = "ok"
		}

		wsReq.Data = replyMsg
		if !uc.SendMessageToChan([]byte(wsReq.Data)) {
			return
		}
	}
}

// WritePump 用户 websocket 连接的写入函数
func (uc *Session) WritePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		logrus.Debugf("websocket user disconnecting from write pump, session: %x", uc)
	}()

	for {
		select {
		case message, ok := <-uc.sendBuf:
			if !ok {
				logrus.Warnf("websocket is closed, reply message from write pump error, session: %s", uc)
				uc.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := uc.sendMessage(message)
			if err != nil {
				return
			}
			logrus.Debugf("websocket send message success, session: %+v, message is %s", uc, string(message))
		case <-ticker.C:
			err := uc.sendMessage([]byte(`{"type": "ping", "data": "ping"}`))
			if err != nil {
				return
			}
		case <-uc.close:
			return // connection 从 readPump 断开了, writePump 中直接 return 即可
		}
	}
}

func (uc *Session) sendMessage(data []byte) error {
	if data == nil {
		return nil
	}
	var err error
	if err = uc.Conn.SetWriteDeadline(time.Now().Add(time.Duration(5))); err != nil {
		logrus.Errorf("websocket set write deadline error: %v, session: %s,  reply message is %s", err, uc, string(data))
		return err
	}
	if err = uc.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
		logrus.Errorf("websocket write message error: %v, session: %s,  reply message is %s", err, uc, string(data))
		return err
	}
	return nil
}
