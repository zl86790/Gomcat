package impl

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"main/name/lizhe/gomcat/config/domain"
)

type XmlReaderImpl struct {
}

func (XmlReaderImpl) Read(fileName string) domain.HandlerConfigs {
	content, err := ioutil.ReadFile(fileName)

	var handlerConfigs domain.HandlerConfigs

	err = xml.Unmarshal(content, &handlerConfigs)
	if err != nil {
		log.Fatal("err:",err)
	}

	return handlerConfigs
}

