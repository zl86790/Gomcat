package domain

import (
	"encoding/xml"
)

type HandlerConfigs struct {
	XMLName        xml.Name         `xml:"handlerConfigs"`
	HandlerConfig []HandlerConfig `xml:"handlerConfig"`
}

type HandlerConfig struct {
	XMLName    xml.Name `xml:"handlerConfig"`
	Id      string `xml:"id,attr"`
	Path    string `xml:"path,attr"`
	Handler string `xml:"handler,attr"`
	Template string `xml:"template,attr"`
}
