package storage


import (
	"github.com/garyburd/redigo/redis"
	"math/rand"
)

const(
	redisNetwork = "tcp"
	redisAddr = "127.0.0.1:6379"
)

var(
	connection redis.Conn
)

func RedisInit(){
	connection,_ = redis.Dial(redisNetwork, redisAddr)
}

func Get(key interface{}) (interface{}, error){
	return connection.Do("GET",key)
}

func Set(key, value interface{}) error{
	_,err := connection.Do("SET",key,value)
	return err
}

func HSet(key, field, value interface{}) error{
	_,err := connection.Do("HSET",key,field,value)
	return err
}

func HGet(key, field interface{}) (interface{}, error){
	return connection.Do("HGET",key,field)
}

func HGetAll(key interface{}) (interface{}, error){
	return connection.Do("HGETALL",key)
}

func ExpireKey(key interface{}, ttl int) error{
	_,err := connection.Do("EXPIRE",key,ttl)
	return err
}

func RedisClose(){
	connection.Close()
}

func shuffle(src []string) []string{
	dest := make([]string, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}