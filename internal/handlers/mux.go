package handlers

import (
	mux2 "github.com/gorilla/mux"
	"main/utils"
)

func StartServer(bus *utils.Bus) *mux2.Router {
	mux := mux2.NewRouter()

	mux.HandleFunc("/api/send", SendPhoto(bus))
	mux.HandleFunc("/api/get/{id}", GetPhoto())

	return mux
}
