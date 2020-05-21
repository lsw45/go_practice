package pool_test

import (
	"github.com/CardInfoLink/tasty/redis"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout: time.Duration(4),
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin:      checkOrigin, //检查源host是否被允许
}

// 接口
func UpgradeHttp(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("创建websocket错误: %s", err.Error())
		return
	}

	var usess Session
	usess.ShopID = "shopId"
	usess.TableNo = "tableNo"
	usess.UserID = "userId"
	usess.RandomKey = "randomKey"
	usess.PayMode = "payMode"

	tableNo := usess.TableNo
	if usess.TableNo == "" {
		tableNo = usess.UserID
	}

	uc := NewSession(usess.ShopID, tableNo, usess.UserID, usess.PayMode, usess.RandomKey, conn)

	DefaultSessionPool.Save(uc) // 将用户的 websocket 连接加入本机缓存中

	uc.SetReadCloseFunc(func(sess *Session) {
		err := redis.WSSessionCache.Remove(sess.TableKey(), sess.UserKey())
		if err != nil {
			logrus.Errorf("report.error.middle remove websocket session key error: %s, session: %s", err.Error(), sess)
		}
	})

	go uc.ReadPump()
	go uc.WritePump()
}

// 是否允许跨域访问
var AllowAllOrigins = false
var AllowOrigins = "^http[s]?://.*(xunliandata.com|everonet.com)$"

func checkOrigin(r *http.Request) bool {
	if AllowAllOrigins {
		return true
	}
	origin := r.Header.Get("Origin")
	logrus.Debugf("allow origins are %v, request origin is %v", AllowOrigins, origin)
	pattern := AllowOrigins
	allowed, _ := regexp.MatchString(pattern, origin)

	// allow private ip
	if !allowed {
		if strings.Contains(origin, "http://localhost") {
			return true
		}
		allowed, _ := regexp.MatchString("http://[0-9]{1,3}\\.[0-9]{1,3}.\\.[0-9]{1,3}\\.[0-9]{1,3}(:[0-9]*)?", origin)
		return allowed
	}

	if !allowed {
		logrus.Warnf("origins not allowed: %v, %v ", origin, AllowOrigins)
	}

	return allowed
}
