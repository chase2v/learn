package main

import (
	"io"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriterString(w, "Create User handler")
}
