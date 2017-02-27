package redis


import (
	"github.com/garyburd/redigo/redis"
)

const(
	redisNetwork = "tcp"
	redisAddr = "127.0.0.1:6379"
)

var(
	connection redis.Conn
)

func Init(){
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

func ExistKey(key interface{}) bool{
	value,err := connection.Do("EXISTS",key)
	return err == nil && value.(int64) == 1
}

func ExistField(key, field interface{}) bool{
	value,err := connection.Do("HEXISTS",key, field)
	return err == nil && value.(int64) == 1
}

func Close(){
	connection.Close()
}