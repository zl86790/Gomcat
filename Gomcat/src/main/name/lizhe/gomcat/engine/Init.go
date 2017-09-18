package engine

import (
	"log"
	"main/name/lizhe/gomcat/config/domain"
	"net/http"
	"main/name/lizhe/gomcat/handler"
)

func Init(handlerConfigs domain.HandlerConfigs) {
	for id, hc := range handlerConfigs.HandlerConfig {
		log.Println(id, " ", hc)
		http.HandleFunc(hc.Path, handler.DoLogic)
	}
	http.ListenAndServe(":8106", nil)
}

