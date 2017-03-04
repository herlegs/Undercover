package dto

import (
	dao "github.com/herlegs/Undercover/storage"
)

type GameConfig struct {
	TotalNum int
	MajorityNum  int
	MinorityNum  int
	MajorityWord string
	MinorityWord string
}

type RoomInfo struct {
	RoomStatus dao.GameState
	GameConfig *GameConfig
	Players []*dao.Player
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
	Authorized bool
	RoomStatus dao.GameState
}

type EndGameResponse struct {
	Authorized bool
	RoomStatus dao.GameState
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
	RoomStatus dao.GameState
}

type JoinGameRequest struct {
	RoomID string
	UserID string
	UserName string
}

func (self *JoinGameRequest) FromValues(values map[string][]string){
	self.UserID = values["userID"][0]
	self.RoomID = values["roomID"][0]
	self.UserName = values["userName"][0]
}

type JoinGameResponse struct {
	RoomStatus dao.GameState
	UserInfo *dao.Player
}

