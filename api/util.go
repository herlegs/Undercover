package api

import "net/http"

func GetUserIDFromRequest(r *http.Request) string{
	return r.RemoteAddr
}
