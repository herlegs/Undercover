package logic

import (
	"sync"
	"errors"
	dao "github.com/herlegs/Undercover/storage"
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
	dao.CreateRoom(roomID, req.AdminID)
	if dao.IsRoomExist(roomID) {
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
	dao.CloseRoom(req.RoomID)
}

func IsRoomAdmin(req * dto.UserIdentityRequest) bool{
	roomAdmin := dao.GetAdmin(req.RoomID)
	return req.UserID == roomAdmin
}

func GetRoomPlayers(req * dto.UserIdentityRequest) []*dao.Player {
	isAdmin := IsRoomAdmin(req)
	roomStatus := dao.GetRoomStatus(req.RoomID)
	players := dao.GetAllInGamePlayer(req.RoomID)
	if isAdmin || roomStatus == dao.Ended{
		return players
	}else{
		return hidePlayerInfo(players, req.UserID)
	}
}

//hide other players information from current player
func hidePlayerInfo(players []*dao.Player, userID string) []*dao.Player{
	maskedPlayers := make([]*dao.Player, len(players))
	for _, player := range players {
		maskedPlayer := &dao.Player{
			ID: player.ID,
			Name: player.Name,
			Alive: player.Alive,
		}
		if player.UserID == userID {
			maskedPlayer.UserID = userID
			maskedPlayer.Word = player.Word
		}
		maskedPlayers = append(maskedPlayers, maskedPlayer)
	}
	return maskedPlayers
}