package dto

import "github.com/herlegs/Undercover/storage"

type GetSessionRequest struct {
	UserID string
}

func (self * GetSessionRequest) FromValues(values map[string][]string){
	self.UserID = values["userID"][0]
}

type GetSessionResponse struct {
	UserID string
	SessionFound bool
	UserInfo *storage.Player
	RoomExist bool
}
