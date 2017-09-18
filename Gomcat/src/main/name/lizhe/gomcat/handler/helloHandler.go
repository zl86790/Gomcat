package handler

import (
	"net/http"
	"fmt"
	"html/template"
)

func DoLogic(w http.ResponseWriter, req *http.Request) {
//	w.Write([]byte("Hello"))
	t, _ := template.ParseFiles("C:\\Users\\Lizhe\\git\\Gomcat\\Gomcat\\src\\main\\name\\lizhe\\gomcat\\handler\\hello.html")
    fmt.Println(t.Name())
    t.Execute(w, "Hello world")
}
