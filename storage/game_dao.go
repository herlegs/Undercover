/***
Table structures:
 */
package storage

import (
	"time"
	"github.com/herlegs/Undercover/redis"
	"encoding/json"
)

/***

 */

type GameState int

const (
	NotExist GameState = iota
	//admin created room
	Created
	//waiting for players to join
	Waiting
	//started
	Started
	//ended
	Ended
)

const(
	Status              = "status"
	Admin               = "admin"
	Counter             = "playerCounter"
	RoomUserTableSuffix = "_room_user"
	UserSessionSuffix       = "_session"
	SessionTTL             = int(time.Hour / time.Second)
)

func addToSession(player *Player){
	sessionTable := player.UserID + UserSessionSuffix
	defer setExpire(sessionTable)
	bytes,_ := json.Marshal(player)
	redis.Set(sessionTable, bytes)
}

func setExpire(key string){
	redis.ExpireKey(key, SessionTTL)
}

func GetUserFromSession(userID string) *Player{
	sessionTable := userID + UserSessionSuffix
	if !redis.ExistKey(sessionTable){
		return nil
	}
	obj,err := redis.Get(sessionTable)
	if err != nil {
		return nil
	}
	player := &Player{}
	err = json.Unmarshal(obj.([]byte), player)
	if err != nil {
		return nil
	}
	return player
}


