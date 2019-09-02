package main

import (
	"github.com/gorilla/mux"
	"github.com/mojiajuzi/forum/action"
)

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", action.Register).Methods("POST")
	r.HandleFunc("/login", action.Login).Methods("POST")
	return r
}
