package storage

import (
	"sync"

	redigo "github.com/garyburd/redigo/redis"

	"github.com/herlegs/Undercover/redis"
)
var roomLocks = make(map[string]*sync.Mutex)

func CreateRoom(room, adminID string){
	defer setExpire(room)
	UpdateRoomStatus(room, Created)
	redis.HSet(room, Counter, 0)
	roomLocks[room] = &sync.Mutex{}
}

func UpdateRoomStatus(room string, status GameState){
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

func IsRoomExist(room string) bool{
	return redis.ExistKey(room)
}