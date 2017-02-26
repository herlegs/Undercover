package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func DoRouting(r *mux.Router){
	doAdminRouting(r)
	doUserRouting(r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./webapp")))
}

func doAdminRouting(r *mux.Router){
	r.HandleFunc("/admin/create", CreateRoomHandler)
	r.HandleFunc("/admin/{room}/start}", StartGameHandler)
	r.HandleFunc("/admin/{room}/end}", EndGameHandler)
}

func doUserRouting(r *mux.Router){
	r.HandleFunc("/{user}/{room}", GameHandler)
}



