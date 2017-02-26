package storage

import (
	"time"
	"sync"
)

type GameState int

const (
	Created GameState = iota
	Waiting
	Started
	Ended
)

const(
	RoomTTL = int(time.Hour / time.Second)
	RoomName = "name"
	Status = "status"
)

type GameCache struct{
	sync.RWMutex
	words []string
	index int
}

func (cache *GameCache) generateWords(word1, word2 string, majorityNum, minorityNum int){
	words := make([]string, majorityNum + minorityNum)
	for i := 0; i < majorityNum; i++ {
		words = append(words, word1)
	}
	for i := 0; i < minorityNum; i++ {
		words = append(words, word2)
	}
	cache.words = words
	cache.index = 0
}

func (cache *GameCache) getWord() (string, error){
	cache.Lock()
	defer cache.Unlock()
	if cache.index < len(cache.words){
		word := cache.words[cache.index]
		cache.index++
		return word, nil
	}
	return "", error("Room is full")
}

//roomName -- wordList
var roomCache = make(map[string]*GameCache)

func CreateRoom(room string){
	userTable := room + "user"
	defer setExpire(room)
	defer setExpire(userTable)
	HSet(room, RoomName, room)
	HSet(room, Status, Created)
}

func setExpire(key string){
	ExpireKey(key, RoomTTL)
}


