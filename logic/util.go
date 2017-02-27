package logic

import (
	"math/rand"
	"time"

	"github.com/herlegs/Undercover/redis"
)

const(
	RoomDigit = 4
	Digits = "1234567890"
)

func generateRoomName() string{
	for{
		rand.Seed(time.Now().UTC().UnixNano())
		digits := make([]byte, RoomDigit)
		for i := 0; i < RoomDigit; i++ {
			digits[i] = Digits[rand.Intn(len(Digits))]
		}
		name := string(digits)
		if !redis.ExistKey(name) {
			return name
		}
	}
}

func shuffle(src []string) []string{
	dest := make([]string, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}