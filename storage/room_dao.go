package storage

import (
	"sync"

	redigo "github.com/garyburd/redigo/redis"

	"github.com/herlegs/Undercover/redis"
)
var roomLocks = make(map[string]*sync.Mutex)

func CreateRoom(room, adminID string){
	SetRoomStatus(room, Created)
	setPlayerCounter(room, 0)
	redis.HSet(room, Admin, adminID)
	roomLocks[room] = &sync.Mutex{}
	CreateNewPlayer(room, adminID, Admin)
}

func SetRoomStatus(room string, status GameState){
	redis.HSet(room, Status, status)
}

func GetRoomStatus(room string) GameState{
	if !IsRoomExist(room){
		return NotExist
	}
	status, err := redigo.Int(redis.HGet(room, Status))
	if err != nil {
		return NotExist
	}
	return GameState(status)
}

func setPlayerCounter(room string, counter int){
	redis.HSet(room, Counter, counter)
}

func getPlayerCounter(room string) int{
	counter,_ := redigo.Int(redis.HGet(room, Counter))
	return counter
}

func GeneratePlayerCounter(room string) int{
	roomLocks[room].Lock()
	counter := getPlayerCounter(room)
	setPlayerCounter(room, counter + 1)
	roomLocks[room].Unlock()
	return counter
}

func setAdmin(room, userID string){
	redis.HSet(room, Admin, userID)
}

func GetAdmin(room string) string{
	admin,_ := redigo.String(redis.HGet(room, Admin))
	return admin
}

func IsRoomAdmin(room, userID string) bool{
	admin := GetAdmin(room)
	return admin == userID
}

func IsRoomExist(room string) bool{
	return redis.ExistKey(room)
}

func CloseRoom(room string){
	redis.Delete(room)
	redis.Delete(room + RoomUserTableSuffix)
}