package pool

import (
	"sync"
)

// SessionPool 包装一下用户 websocket 连接的 map, 并发安全
type SessionPool struct {
	sync.RWMutex
	m map[string]map[string]*Session
}

var DefaultSessionPool = NewSessionPool()

func NewSessionPool() *SessionPool {
	return &SessionPool{
		m: make(map[string]map[string]*Session),
	}
}

func (sm *SessionPool) Save(s *Session) {
	sm.Lock()
	if _, ok := sm.m[s.TableKey()]; ok {
		sm.m[s.TableKey()][s.UserKey()] = s
	} else {
		sm.m[s.TableKey()] = map[string]*Session{s.UserKey(): s}
	}
	sm.Unlock()
}

func (sm *SessionPool) Get(key string) (m map[string]*Session) {
	sm.RLock()
	m, _ = sm.m[key]
	sm.RUnlock()
	return
}

func (sm *SessionPool) UnexistUser(tableKey string, keys []string) (ks []string) {
	sm.RLock()
	if m, ok := sm.m[tableKey]; ok {
		for _, key := range keys {
			if _, b := m[key]; !b {
				ks = append(ks, key)
			}
		}
	} else {
		ks = keys
	}
	sm.RUnlock()
	return
}

func (sm *SessionPool) Sessions(key string) (ss []*Session) {
	sm.Lock()
	if m, ok := sm.m[key]; ok {
		for _, s := range m {
			ss = append(ss, s)
		}
	}
	sm.Unlock()
	return
}

func (sm *SessionPool) Remove(s *Session) {
	sm.Lock()
	if _, ok := sm.m[s.TableKey()]; ok {
		delete(sm.m[s.TableKey()], s.UserKey())
	}
	sm.Unlock()
}

func (sm *SessionPool) ConnectCount() (count int) {
	sm.RLock()
	for _, v := range sm.m {
		count += len(v)
	}
	sm.RUnlock()
	return
}
