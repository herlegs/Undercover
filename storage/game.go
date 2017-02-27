/***
Table structures:

Room: roomKey ~ {
	status: room (GameState)
	admin: userID
	majorityNum: int
	majorityWord: string
	minorityNum: int
	minorityWord: string
	playerCounter: int
}

UserWordTable: roomKey + "word" ~ {
	userID: word
}

UserNameTable: roomKey + "name" ~ {
	userID: name
}
 */
package storage

import (
	"time"

	redigo "github.com/garyburd/redigo/redis"

	"github.com/herlegs/Undercover/redis"
)

/***

 */

type GameState int

const (
	NotExist GameState = iota
	//admin created room
	Created
	//waiting for admin to start game with config
	Configuring
	//waiting for players to join
	Waiting
	//started
	Started
	//ended
	Ended
)

const(
	RoomTTL = int(time.Hour / time.Second)
	Status = "status"
	Counter = "playerCounter"
	UserTableSuffix = "user"
)

func CreateNewPlayer(room, userID, userName string){
	if !IsRoomExist(room) || IsPlayerExist(room, userID){
		return
	}
	//create player ID from counter
	player := &Player{}
	player.ID = GetPlayerCounter(room)
	player.Name = userName
	player.Word = ""
	SetPlayerInfo(room, userID, player)
}

func GetPlayerCounter(room string) int{
	roomLocks[room].Lock()
	counter,_ := redigo.Int(redis.HGet(room, Counter))
	redis.HSet(room, Counter, counter + 1)
	roomLocks[room].Unlock()
	return counter
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

func setExpire(key string){
	redis.ExpireKey(key, RoomTTL)
}


