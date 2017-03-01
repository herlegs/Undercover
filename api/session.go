package api

import (
	"net/http"
	"github.com/herlegs/Undercover/api/dto"
	"github.com/herlegs/Undercover/storage"
)

func GetSessionHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.GetSessionRequest)
	user := storage.GetUserFromSession(request.UserID)
	response := dto.GetSessionResponse{
		UserID: user.UserID,
		SessionFound: false,
		UserInfo: nil,
		RoomExist: false,
	}
	if user == nil {
		WriteResponse(w, http.StatusNotFound, response)
	}else{
		response.UserInfo = user
		response.RoomExist = storage.IsRoomExist(user.RoomID)
		WriteResponse(w, OK, response)
	}
}