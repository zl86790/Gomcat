package handler

import (
	"net/http"
)

type HandlerInter interface{
	DoLogic(w http.ResponseWriter, req *http.Request)
	GetFunction() func() int
}

