package handler

import (
	"net/http"
)

func GetFunction(){

}

func DoLogic(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}
