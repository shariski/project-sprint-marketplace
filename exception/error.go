package exception

import (
	"github.com/gofiber/fiber/v2/log"
)

func PanicLogging(err error) {
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
