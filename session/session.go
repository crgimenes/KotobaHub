package session

import (
	"crypto/rand"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const cookieName = "kotobahub_session"

type sessionData struct {
	expireAt time.Time
	data     map[string]any
}

var (
	sessionDataMap map[string]*sessionData
	mux            sync.Mutex
)

func Load() {
	sessionDataMap = make(map[string]*sessionData)

	go func() {
		for {
			removeExpired()
			time.Sleep(1 * time.Hour)
		}
	}()
}

func get(r *http.Request) (string, *sessionData, bool) {
	cookies := r.Cookies()
	if len(cookies) == 0 {
		return "", nil, false
	}

	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return "", nil, false
	}

	id := cookie.Value
	mux.Lock()
	s, ok := sessionDataMap[id]
	mux.Unlock()
	if !ok {
		return "", nil, false
	}

	if s.expireAt.Before(time.Now()) {
		mux.Lock()
		delete(sessionDataMap, cookie.Value)
		mux.Unlock()
		return "", nil, false
	}

	return cookie.Value, s, true
}

func Delete(w http.ResponseWriter, id string) {
	mux.Lock()
	delete(sessionDataMap, id)
	mux.Unlock()
	cookie := http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}

func Save(r *http.Request, w http.ResponseWriter) (id string) {
	expireAt := time.Now().Add(3 * time.Hour)
	id, s, ok := get(r)
	if !ok {
		id = RandomID()
		s = &sessionData{
			expireAt: expireAt,
			data:     make(map[string]any),
		}
	}

	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   id,
		Expires: expireAt,
	}

	s.expireAt = expireAt

	mux.Lock()
	sessionDataMap[id] = s
	mux.Unlock()

	http.SetCookie(w, cookie)

	return id
}

func Get(id string) any {
	mux.Lock()
	d := sessionDataMap[id]
	mux.Unlock()
	if d == nil {
		return nil
	}
	return d.data
}

func Set(id string, key string, value any) {
	mux.Lock()
	sessionDataMap[id].data[key] = value
	mux.Unlock()
}

func GetString(id string, key string) string {
	s := Get(id)
	if s == nil {
		return ""
	}

	switch s.(type) {
	case int:
		return strconv.Itoa(s.(int))
	case int64:
		return strconv.FormatInt(s.(int64), 10)
	case float64:
		return strconv.FormatFloat(s.(float64), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(s.(bool))
	case []byte:
		return string(s.([]byte))
	case string:
		return s.(string)
	default:
		log.Printf("session.GetString: key=%s, value=%v is not supported type %T\n", key, s, s)
		return ""
	}
}

func removeExpired() {
	for k, v := range sessionDataMap {
		if v.expireAt.Before(time.Now()) {
			delete(sessionDataMap, k)
		}
	}
}

func RandomID() string {
	const (
		length  = 16
		charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	)
	lenCharset := byte(len(charset))
	b := make([]byte, length)
	rand.Read(b)
	for i := 0; i < length; i++ {
		b[i] = charset[b[i]%lenCharset]
	}
	return string(b)
}
