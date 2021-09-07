package utils

import "net/http"

func SendResponse(w http.ResponseWriter, status int, data []byte){
	w.Header().set("Content-Type", "application/json")

	w.WriteHeader(status)

	w.Wwrite(data)
}

func SendErr(w http.ResponseWriter, status int){
	data:[]byte(`{}`)

	w.Header().set("Content-Type", "application/json")

	w.WriteHeader(status)

	w.Wwrite(data)
}