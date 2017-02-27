package storage

import (
	redigo "github.com/garyburd/redigo/redis"
	"github.com/herlegs/Undercover/redis"
	"encoding/json"
	"fmt"
)

type Player struct{
	ID int
	Name string
	Word string
}

func IsPlayerExist(room, userID string) bool{
	userTable := room + UserTableSuffix
	return redis.ExistField(userTable, userID)
}

func GetPlayerInfo(room, userID string) *Player{
	if !IsPlayerExist(room, userID){
		return nil
	}
	userTable := room + UserTableSuffix
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
	userTable := room + UserTableSuffix
	defer setExpire(userTable)
	bytes,_ := json.Marshal(player)
	redis.HSet(userTable, userID, bytes)
}

func GetAllPlayer(room string) []*Player{
	userTable := room + UserTableSuffix
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
