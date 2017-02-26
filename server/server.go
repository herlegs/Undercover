package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"github.com/herlegs/Undercover/api"
)

const(
	LOCALHOST = "0.0.0.0"
)

func Start(){
	router := mux.NewRouter()

	api.DoRouting(router)

	server := initServer(router)

	server.ListenAndServe()
}

func initServer(r *mux.Router) *http.Server{
	return &http.Server{
		Addr: LOCALHOST,
		Handler: r,
		ReadTimeout: time.Second * 3,
		WriteTimeout: time.Second * 3,
	}
}

