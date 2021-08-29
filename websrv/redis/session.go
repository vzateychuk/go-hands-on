package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/garyburd/redigo/redis"
)

type Session struct {
	Login     string
	Useragent string
}

type SessionID struct {
	ID string
}

const sessKeyLen = 10

type SessionManager struct {
	redisConn redis.Conn
}

func NewSessionManager(conn redis.Conn) *SessionManager {
	return &SessionManager{
		redisConn: conn,
	}
}

// Создание новой сессии
func (sm *SessionManager) Create(in *Session) (*SessionID, error) {
	// Получили случайный session ID
	id := SessionID{RandStringRunes(sessKeyLen)}
	// Сериализуем в JSON
	sessJson, _ := json.Marshal(in)
	mkey := "sessions:" + id.ID
	// Сохраняем сессию в Redis выполняя коману SET
	data, err := sm.redisConn.Do("SET", mkey, sessJson, "EX", 86400)
	// Преобразуем ответ Redis к строке
	result, err := redis.String(data, err)
	if err != nil {
		return nil, err
	}
	if result != "OK" {
		return nil, fmt.Errorf("result not OK")
	}
	return &id, nil
}

// Выборка сессии по ID
func (sm *SessionManager) Check(in *SessionID) *Session {
	mkey := "sessions:" + in.ID
	data, err := redis.Bytes(sm.redisConn.Do("GET", mkey))
	if err != nil {
		log.Println("cant get data:", err)
		return nil
	}
	sess := &Session{}
	err = json.Unmarshal(data, sess)
	if err != nil {
		log.Println("cant unpack session data:", err)
		return nil
	}
	return sess
}

// Удалить значение из Redis
func (sm *SessionManager) Delete(in *SessionID) {
	mkey := "sessions:" + in.ID
	_, err := redis.Int(sm.redisConn.Do("DEL", mkey))
	if err != nil {
		log.Println("redis error:", err)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
