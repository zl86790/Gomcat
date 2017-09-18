package inter

import(
	"main/name/lizhe/gomcat/config/domain"
)

type Reader interface {
	Read(fileName string) domain.HandlerConfigs
}
