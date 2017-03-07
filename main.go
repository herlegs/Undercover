package main

import (
	"github.com/herlegs/Undercover/server"
	"github.com/herlegs/Undercover/redis"
)

func main(){
	defer cleanup()
	server.Start()
}

func cleanup(){
	redis.Close()
}