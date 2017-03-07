package api

import (
	"net/http"
	"github.com/herlegs/Undercover/logic"
	"github.com/herlegs/Undercover/api/dto"
	dao "github.com/herlegs/Undercover/storage"
)

const(
	ServerError = http.StatusInternalServerError
	OK = http.StatusOK
)

func CreateRoomHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request,_ := reqDto.(*dto.CreateRoomRequest)
	resp,err := logic.CreateNewRoom(request)
	if err != nil {
		WriteResponse(w, ServerError, err.Error())
		return
	}
	WriteResponse(w, OK, resp)
}

func StartGameHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.StartGameRequest)
	resp := logic.StartGame(request)
	if !resp.Authorized {
		WriteResponse(w, http.StatusForbidden, resp)
		return
	}
	WriteResponse(w, OK, resp)
}

func EndGameHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.UserIdentityRequest)
	resp := logic.EndGame(request)
	if !resp.Authorized {
		WriteResponse(w, http.StatusForbidden, resp)
		return
	}
	WriteResponse(w, OK, resp)
}

func CloseRoomHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.CloseRoomRequest)
	request.AdminID = GetUserIDFromRequest(r)
	logic.CloseRoom(request)
}

func ValidateAdminHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.UserIdentityRequest)
	isAdmin := logic.IsRoomAdmin(request)
	resp := &dto.ValidateUserResponse{
		RoomID: request.RoomID,
		RoomStatus: dao.GetRoomStatus(request.RoomID),
	}
	if isAdmin {
		WriteResponse(w, OK, resp)
	}else{
		WriteResponse(w, http.StatusForbidden, "")
	}
}

func ValidatePlayerHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.UserIdentityRequest)
	isAdmin := logic.IsRoomAdmin(request)
	resp := &dto.ValidateUserResponse{
		RoomID: request.RoomID,
		RoomStatus: dao.GetRoomStatus(request.RoomID),
	}
	if isAdmin {
		WriteResponse(w, http.StatusForbidden, "")
	}else if resp.RoomStatus == dao.NotExist{
		WriteResponse(w, http.StatusNotFound, "")
	}else{
		WriteResponse(w, OK, resp)
	}
}