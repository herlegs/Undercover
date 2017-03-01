package logic

import (
	"sync"
	"errors"
	"github.com/herlegs/Undercover/storage"
	"github.com/herlegs/Undercover/api/dto"
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
	cache.words = shuffle(words)
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
	return "", errors.New("Room is full")
}

//roomName -- wordList
var roomCache = make(map[string]*GameCache)

func CreateNewRoom(req *dto.CreateRoomRequest) (*dto.CreateRoomResponse, error){
	roomID := generateRoomName()
	storage.CreateRoom(roomID, req.AdminID)
	if storage.IsRoomExist(roomID) {
		return &dto.CreateRoomResponse{RoomID:roomID}, nil
	} else{
		return nil, errors.New("Failed creating room")
	}
}

func StartGame(req *dto.StartGameRequest) (*dto.StartGameRequest, error){
	return nil,nil
}

func EndGame(req *dto.EndGameRequest) (*dto.EndGameRequest, error){
	return nil,nil
}

func CloseRoom(req * dto.CloseRoomRequest) {
	storage.CloseRoom(req.RoomID)
}