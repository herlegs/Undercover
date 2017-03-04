package storage

import (
	redigo "github.com/garyburd/redigo/redis"
	"github.com/herlegs/Undercover/redis"
	"encoding/json"
)

type Player struct{
	UserID string
	RoomID string
	//player ID in room
	ID int
	Name string
	Word string
	IsMinority bool
	InGame bool
	Alive bool
	HasVoted int
}

func IsPlayerExist(room, userID string) bool{
	userTable := room + RoomUserTableSuffix
	return redis.ExistField(userTable, userID)
}

func GetPlayerInfo(room, userID string) *Player{
	if !IsPlayerExist(room, userID){
		return nil
	}
	userTable := room + RoomUserTableSuffix
	obj,err := redis.HGet(userTable, userID)
	if err != nil {
		return nil
	}
	player := &Player{}
	err = json.Unmarshal(obj.([]byte), player)
	if err != nil {
		return nil
	}
	return player
}

func SetPlayerInfo(room, userID string, player *Player){
	defer addToSession(player)
	userTable := room + RoomUserTableSuffix
	bytes,_ := json.Marshal(player)
	redis.HSet(userTable, userID, bytes)
}

func GetAllPlayer(room string) []*Player{
	userTable := room + RoomUserTableSuffix
	players := make([]*Player,0)
	playerMap,err := redigo.StringMap(redis.HGetAll(userTable))
	if err != nil {
	   return players
	}
	for userID, objStr := range playerMap{
		player := &Player{}
		err := json.Unmarshal([]byte(objStr), player)
		if err == nil {
			if(player.Name == ""){
				player.Name = userID
			}
			players = append(players, player)
		}
	}
	return players
}


func CreateNewPlayer(room, userID, userName string) *Player{
	if !IsRoomExist(room) || IsPlayerExist(room, userID){
		return nil
	}
	//create player ID from counter
	player := &Player{
		UserID : userID,
		RoomID : room,
		ID : GeneratePlayerCounter(room),
		Name : userName,
		Word : "",
		InGame : false,
		Alive : true,
		HasVoted : -1,
	}
	SetPlayerInfo(room, userID, player)
	return player
}

func UpdateUserName(room, userID, userName string){
	player := GetPlayerInfo(room, userID)
	if player == nil{
		return
	}
	player.Name = userName
	SetPlayerInfo(room, userID, player)
}

func GetUserName(room, userID string) string{
	player := GetPlayerInfo(room, userID)
	if player == nil || player.Name == ""{
		return userID
	}
	return player.Name
}

func UpdateUserWord(room, userID, word string){
	player := GetPlayerInfo(room, userID)
	if player == nil {
		return
	}
	player.Word = word
	SetPlayerInfo(room, userID, player)
}

func GetUserWord(room, userID string) string{
	player := GetPlayerInfo(room, userID)
	if player == nil {
		return ""
	}
	return player.Word
}

func GetAllInGamePlayer(room string) []*Player{
	filtered := make([]*Player, 0)
	allPlayers := GetAllPlayer(room)
	for _, player := range allPlayers{
		if player.InGame {
			filtered = append(filtered, player)
		}
	}
	return filtered
}

func ResetAllPlayerInGameStatus(room string){
	allPlayers := GetAllPlayer(room)
	for _, player := range allPlayers{
		player.InGame = false;
		SetPlayerInfo(room, player.UserID, player)
	}
}

func SetPlayerInGameStatus(room, userID string){
	player := GetPlayerInfo(room, userID)
	player.InGame = false;
	SetPlayerInfo(room, player.UserID, player)
}