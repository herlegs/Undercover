package redis


import (
	"github.com/garyburd/redigo/redis"
)

func Get(key interface{}) (interface{}, error){
	connection := pool.Get()
	defer connection.Close()
	return connection.Do("GET",key)
}

func Set(key, value interface{}) error{
	connection := pool.Get()
	defer connection.Close()
	_,err := connection.Do("SET",key,value)
	return err
}

func HSet(key, field, value interface{}) error{
	connection := pool.Get()
	defer connection.Close()
	_,err := connection.Do("HSET",key,field,value)
	return err
}

func HGet(key, field interface{}) (interface{}, error){
	connection := pool.Get()
	defer connection.Close()
	return connection.Do("HGET",key,field)
}

func HGetAll(key interface{}) (interface{}, error){
	connection := pool.Get()
	defer connection.Close()
	return connection.Do("HGETALL",key)
}

func ExpireKey(key interface{}, ttl int) error{
	connection := pool.Get()
	defer connection.Close()
	_,err := connection.Do("EXPIRE",key,ttl)
	return err
}

func ExistKey(key interface{}) bool{
	connection := pool.Get()
	defer connection.Close()
	exists,err := redis.Bool(connection.Do("EXISTS",key))
	return err == nil && exists
}

func ExistField(key, field interface{}) bool{
	connection := pool.Get()
	defer connection.Close()
	exists,err := redis.Bool(connection.Do("HEXISTS",key, field))
	return err == nil && exists
}

func Delete(key interface{}) error{
	connection := pool.Get()
	defer connection.Close()
	_,err := connection.Do("DEL",key)
	return err
}