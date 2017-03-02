package dto

import "github.com/herlegs/Undercover/storage"

type GameConfig struct {
	MajorityNum int
	MinorityNum int
	MajorWord string
	MinorWord string
}

type StartGameRequest struct {
	RoomID string
	AdminID string
	GameConfig
}

func (self *StartGameRequest) FromValues(values map[string][]string){
	self.RoomID = values["roomID"][0]
	self.AdminID = values["userID"][0]
}

type StartGameResponse struct {
	Success bool
	ErrorMessage string
}

type EndGameRequest struct {
	RoomID string
	AdminID string
}

func (self *EndGameRequest) FromValues(values map[string][]string){
	self.RoomID = values["roomID"][0]
	self.AdminID = values["userID"][0]
}

type EndGameResponse struct {

}

type JoinGameRequest struct {
	UserID string
	UserName string
}

func (self *JoinGameRequest) FromValues(values map[string][]string){
	self.UserID = values["userID"][0]
	self.UserName = values["userName"][0]
}

type UserIdentityRequest struct {
	UserID string
	RoomID string
}

func (self *UserIdentityRequest) FromValues(values map[string][]string){
	self.UserID = values["userID"][0]
	self.RoomID = values["roomID"][0]
}

type ValidateUserResponse struct {
	RoomID string
	RoomStatus storage.GameState
}