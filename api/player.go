package api

import (
	"net/http"
	"github.com/herlegs/Undercover/api/dto"
	"github.com/herlegs/Undercover/logic"
)

func GetRoomInfoHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.UserIdentityRequest)
	resp := logic.GetRoomInfo(request)
	WriteResponse(w, OK, resp)
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.JoinGameRequest)
	resp,err := logic.JoinRoom(request)
	if err != nil {
		WriteResponse(w, http.StatusForbidden, err.Error())
	}
	WriteResponse(w, OK, resp)
}
