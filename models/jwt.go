package models

import (
	"gin/config"

	jwt "github.com/dgrijalva/jwt-go"
)

//StudentClaims ...
type StudentClaims struct {
	Student
	jwt.StandardClaims
}

//New ...
func (s *StudentClaims) New() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, s)
	ss, err := token.SignedString(config.GetConfig().GetString("jwt.key"))
	if err != nil {
		return "", err
	}
	return ss, nil
}
