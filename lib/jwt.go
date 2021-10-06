package lib

import (
	"github.com/golang-jwt/jwt"
	"github.com/vandyahmad24/gin-app/constants"
)

type AuthService interface {
	GenerateToken(userId int) (string, error)
}
type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(constants.JWTKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
