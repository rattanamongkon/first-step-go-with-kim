package utils

import (
	"fmt"

	"github.com/spf13/viper"

	"golang.org/x/crypto/bcrypt"
)

var tokenSecret = viper.GetString(`SALT_PASSWORD`)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+tokenSecret), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword+tokenSecret))
}
