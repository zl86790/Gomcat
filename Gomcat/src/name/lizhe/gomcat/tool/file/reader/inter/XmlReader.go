package inter

import(
	"name/lizhe/gomcat/config/domain"
)

type Reader interface {
	Read(fileName string) domain.HandlerConfigs
}
