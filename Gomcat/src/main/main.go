package main

import (
	"log"
	"name/lizhe/gomcat/engine"
	"name/lizhe/gomcat/config/domain"
	"name/lizhe/gomcat/tool/file/reader/impl"
)

func main() {

	var handlerConfigs domain.HandlerConfigs

	var reader = new(impl.XmlReaderImpl)
//	handlerConfigs = reader.Read("E:\\gosource\\git\\Gomcat\\src\\main\\handlers.xml")
	handlerConfigs = reader.Read("C:\\Users\\Lizhe\\git\\Gomcat\\Gomcat\\src\\main\\handlers.xml")

	log.Println(handlerConfigs.HandlerConfig[0].Id)
	log.Println(handlerConfigs.HandlerConfig[0].Path)
	log.Println(handlerConfigs.HandlerConfig[0].Handler)

	engine.Init(handlerConfigs)

}
