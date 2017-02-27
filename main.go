package main

import (
	//"github.com/herlegs/Undercover/server"
	"github.com/herlegs/Undercover/redis"
	//redigo "github.com/garyburd/redigo/redis"

	"github.com/herlegs/Undercover/storage"
	"fmt"
)

func init(){
	redis.Init()
}

func main(){
	defer cleanup()
	//server.Start()
	storage.UpdateRoomStatus("user", storage.NotExist)
	stat := storage.GetRoomStatus("user")
	fmt.Println(stat == storage.NotExist)

}

func cleanup(){
	redis.Close()
}