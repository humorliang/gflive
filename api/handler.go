package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

// httprouter handle
// type Handle func(http.ResponseWriter, *http.Request, Params)
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w,"create user")
}

func NotFound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(404)
	io.WriteString(w,"this is 404 ")
}