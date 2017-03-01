package dto

type CreateRoomRequest struct {
	AdminID string
}

func (self *CreateRoomRequest) FromValues(values map[string][]string){
	self.AdminID = values["userID"][0]
}

type CreateRoomResponse struct {
	RoomID string
}

type CloseRoomRequest struct {
	AdminID string
	RoomID string
}

func (self *CloseRoomRequest) FromValues(values map[string][]string){
	self.AdminID = values["userID"][0]
	self.RoomID = values["roomID"][0]
}