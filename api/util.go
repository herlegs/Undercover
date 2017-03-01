package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"net/url"
	"encoding/json"
	"github.com/herlegs/Undercover/api/dto"
	"strings"
)

type ApiRequestHandler func(http.ResponseWriter, *http.Request, dto.Request)

func GenerateRequestHandler(apiHandler ApiRequestHandler, dto dto.Request) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		ParseRequest(r, dto)
		apiHandler(w, r, dto)
	}
}

func ParseRequest(r *http.Request, obj dto.Request){
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") && r.ContentLength > 0 {
		jsonDecoder := json.NewDecoder(r.Body)
		jsonDecoder.Decode(obj)
	}
	values := extractValues(r)
	obj.FromValues(values)
}

func extractValues(r *http.Request) url.Values{
	r.ParseForm()
	vars := mux.Vars(r)
	if vars != nil {
		for key, value := range vars {
			r.Form.Set(key, value)
		}
	}
	userID := GetUserIDFromRequest(r)
	r.Form.Set("userID", userID)
	return r.Form
}

func GetUserIDFromRequest(r *http.Request) string{
	fullAddr := r.RemoteAddr
	lastIdx := strings.LastIndex(fullAddr, ":")
	ipAddr := fullAddr[:lastIdx]
	return ipAddr
}

func WriteResponse(w http.ResponseWriter, status int, obj interface{}){
	bytes,_ := json.Marshal(obj)
	w.WriteHeader(status)
	w.Write(bytes)
}
