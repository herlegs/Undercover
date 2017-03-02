package api

import (
	"net/http"
	"github.com/herlegs/Undercover/api/dto"
	"github.com/herlegs/Undercover/logic"
)


func GetRoomPlayerHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.UserIdentityRequest)
	players := logic.GetRoomPlayers(request)
	WriteResponse(w, OK, players)
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	//add player to room
	//return playerInfo and room status
	//client:
	//if status is created periodically call GameControlHandler
}



func GameControlHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	//(delete admin configuring)
	//get room status

	//if it's created
	//return roomstatus + playerInfo

	//if it's waiting for player
	//call getword(change player status in game)
	//return updated playerInfo and gameinfo ( roomstatus and num of people)

	//if it's started
	//return playerinfo and roomstatus

	//if it's ended
	//return playerinfo and room status

	//client:
	//if get room status is created:
	//periodically call GameControlHandler

	//if waiting
	//show word and call getallplayers until player in room reach num of people
	//periodically call getgamestatus until game ended (if so call get all players)

	//if game started
	//display room full

	//if game ended
	//periodically call
}

func GameHandler(w http.ResponseWriter, r *http.Request){

}
