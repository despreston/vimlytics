package api

import (
	"log"
	"net/http"
)

type Error struct {
	Error   error
	Message string
	Code    int
}

type Handler func(http.ResponseWriter, *http.Request) *Error

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		log.Printf("%+v", err)
		http.Error(w, err.Message, err.Code)
	}
}
