package handler

import (
	"net/http"
)

func DoLogic(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}
