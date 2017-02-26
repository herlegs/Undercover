package Undercover

import (
	"github.com/herlegs/Undercover/server"
	"github.com/herlegs/Undercover/storage"
)

func init(){
	storage.RedisInit()
}

func main(){
	defer cleanup()
	server.Start()
}

func cleanup(){
	storage.RedisClose()
}