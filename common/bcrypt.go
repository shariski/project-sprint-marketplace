package common

import (
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/exception"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GenerateBcryptPassword(password string, config configuration.Config) string {
	salt, err := strconv.Atoi(config.Get("BCRYPT_SALT"))
	exception.PanicLogging(err)
	bcrypted, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	exception.PanicLogging(err)
	return string(bcrypted)
}
