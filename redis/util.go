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
	exists,err := redis.Bool(connection.Do("EXISTS",key))
	return err == nil && exists
}

func ExistField(key, field interface{}) bool{
	exists,err := redis.Bool(connection.Do("HEXISTS",key, field))
	return err == nil && exists
}

func Delete(key interface{}) error{
	_,err := connection.Do("DEL",key)
	return err
}

func Close(){
	connection.Close()
}