package utils

import "net/http"

func SendResponse(w http.ResponseWriter, status int, data []byte){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func SendErr(w http.ResponseWriter, status int) {
	data := []byte(`{}`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}