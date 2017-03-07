package redis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

const(
	UndercoverDB = 3
	redisNetwork = "tcp"
	redisAddr = "127.0.0.1:6379"
	MaxIdle = 3
	MaxActive = 20
	IdleTimeout = 300 * time.Second
)

var pool *redis.Pool

func init(){
	pool = &redis.Pool{
		MaxIdle: MaxIdle,
		MaxActive: MaxActive,
		IdleTimeout: IdleTimeout,
		TestOnBorrow: func(c redis.Conn, t time.Time) error{
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(redisNetwork, redisAddr)
			if err != nil {
				return nil, err
			}
			if _,err := c.Do("SELECT", UndercoverDB); err != nil{
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func Close(){
	pool.Close()
}
