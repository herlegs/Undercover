package logic

import (
	"sync"
	"errors"
	dao "github.com/herlegs/Undercover/storage"
	"github.com/herlegs/Undercover/api/dto"
)

const(
	MINORITY int = iota
	MAJORITY
)

type GameCache struct{
	sync.RWMutex
	identities []int
	index int
	majorityNum int
	minorityNum int
	majorityWord string
	minorityWord string
}

func (cache *GameCache) generateWords(majorWord, minorWord string, majorityNum, minorityNum int){
	ids := make([]int, majorityNum + minorityNum)
	for i := 0; i < majorityNum; i++ {
		ids = append(ids, MAJORITY)
	}
	for i := 0; i < minorityNum; i++ {
		ids = append(ids, MINORITY)
	}
	cache.identities = shuffle(ids)
	cache.index = 0
}

func (cache *GameCache) getWord() (string, error){
	cache.Lock()
	defer cache.Unlock()
	if cache.index < len(cache.identities){
		identity := cache.identities[cache.index]
		cache.index++
		word := ""
		if identity == MAJORITY {
			word = cache.majorityWord
		}else{
			word = cache.majorityWord
		}
		return word, nil
	}
	return "", errors.New("Room is full")
}

//roomName -- gameCache
var gameCacheMap = make(map[string]*GameCache)

func CreateNewRoom(req *dto.CreateRoomRequest) (*dto.CreateRoomResponse, error){
	roomID := generateRoomName()
	dao.CreateRoom(roomID, req.AdminID)
	if dao.IsRoomExist(roomID) {
		return &dto.CreateRoomResponse{RoomID:roomID}, nil
	} else{
		return nil, errors.New("Failed creating room")
	}
}

func StartGame(req *dto.StartGameRequest) (*dto.StartGameResponse){
	resp := &dto.StartGameResponse{}
	if !dao.IsRoomAdmin(req.RoomID, req.AdminID) {
		resp.Authorized = false
		return resp
	}
	resp.Authorized = true
	//generate words
	gameCache := &GameCache{}
	gameCache.generateWords(req.MajorityWord, req.MinorityWord, req.MajorityNum, req.MinorityNum)
	gameCacheMap[req.RoomID] = gameCache
	dao.SetRoomStatus(req.RoomID, dao.Waiting)
	resp.RoomStatus = dao.Waiting
	return resp
}

func EndGame(req *dto.EndGameRequest) (*dto.EndGameRequest, error){
	return nil,nil
}

func CloseRoom(req * dto.CloseRoomRequest) {
	dao.CloseRoom(req.RoomID)
}

func IsRoomAdmin(req * dto.UserIdentityRequest) bool{
	return dao.IsRoomAdmin(req.RoomID, req.UserID)
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

func GetGameConfig(req * dto.UserIdentityRequest) *dto.GameConfig{
	roomStatus := dao.GetRoomStatus(req.RoomID)
	isAdmin := IsRoomAdmin(req)
	gameCache := gameCacheMap[req.RoomID]
	gameConfig := &dto.GameConfig{
		MajorityNum: gameCache.majorityNum,
		MinorityNum: gameCache.minorityNum,
		MajorityWord: gameCache.majorityWord,
		MinorityWord: gameCache.minorityWord,
		TotalNum: gameCache.majorityNum + gameCache.minorityNum,
	}
	if isAdmin || roomStatus == dao.Ended{
		return gameConfig
	}else{
		return hideGameConfig(gameConfig)
	}
}


func GetRoomInfo(req * dto.UserIdentityRequest) *dto.RoomInfo {
	roomStatus := dao.GetRoomStatus(req.RoomID)
	players := GetRoomPlayers(req)
	gameConfig := GetGameConfig(req)
	return &dto.RoomInfo{
		RoomStatus: roomStatus,
		GameConfig: gameConfig,
		Players: players,
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

//hide game config from player
func hideGameConfig(gameConfig *dto.GameConfig) *dto.GameConfig{
	return &dto.GameConfig{
		TotalNum: gameConfig.TotalNum,
	}
}