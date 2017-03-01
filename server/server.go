package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"github.com/herlegs/Undercover/api"
	"net"
	"strings"
	"fmt"
)

const(
	PORT = ":8000"
)

func Start(){
	router := mux.NewRouter()

	api.DoRouting(router)

	svr := initServer(router)

	svr.ListenAndServe()
}

func initServer(r *mux.Router) *http.Server{
	ip := GetOutboundIP()
	fmt.Println("server listening on:", ip + PORT)
	return &http.Server{
		Addr: ip + PORT,
		Handler: r,
		ReadTimeout: time.Second * 3,
		WriteTimeout: time.Second * 3,
	}
}

// Get preferred ip of this machine
func GetOutboundIP() string {
	localAddrs,_ := net.InterfaceAddrs()
	address := "0.0.0.0"
	for _,addr := range localAddrs{
		idx := strings.LastIndex(addr.String(), "/")
		addrStr := addr.String()[0:idx]
		if strings.HasPrefix(addrStr, "169"){
			address = addrStr
		}
	}
	return address
}

