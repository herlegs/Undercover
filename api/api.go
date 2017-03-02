package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/herlegs/Undercover/api/dto"
)

func DoRouting(r *mux.Router){
	r.HandleFunc("/recent_session",
		GenerateRequestHandler(
			GetSessionHandler,
			&dto.GetSessionRequest{})).Methods("get")

	r.HandleFunc("/admin/create",
		GenerateRequestHandler(
			CreateRoomHandler,
			&dto.CreateRoomRequest{})).Methods("get")
	r.HandleFunc("/admin/{roomID}/startgame",
		GenerateRequestHandler(
			StartGameHandler,
			&dto.StartGameRequest{})).Methods("post")
	r.HandleFunc("/admin/{roomID}/endgame",
		GenerateRequestHandler(
			EndGameHandler,
			&dto.EndGameRequest{})).Methods("post")
	r.HandleFunc("/admin/{roomID}/close",
		GenerateRequestHandler(
			CloseRoomHandler,
			&dto.CloseRoomRequest{})).Methods("delete")
	r.HandleFunc("/admin/{roomID}/validate",
		GenerateRequestHandler(
			ValidateAdminHandler,
			&dto.UserIdentityRequest{})).Methods("get")

	r.HandleFunc("/player/{roomID}/validate",
		GenerateRequestHandler(
			ValidatePlayerHandler,
			&dto.UserIdentityRequest{})).Methods("get")
	r.HandleFunc("player/{userID}/{roomID}", GameHandler).Methods("get")

	r.HandleFunc("/game/{roomID}/getplayers",
		GenerateRequestHandler(
			GetRoomPlayerHandler,
			&dto.UserIdentityRequest{})).Methods("get")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./webapp")))
}

func doAdminRouting(r *mux.Router){

}

func doUserRouting(r *mux.Router){

}



