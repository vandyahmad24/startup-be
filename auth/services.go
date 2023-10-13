package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"startup/config"
)

type Service interface {
	GenerateToken(id int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

// var SECRET_KEY = []byte("VANDY_AHMAD")

func NewService() *jwtService {

	return &jwtService{}
}

func (s *jwtService) GenerateToken(id int) (string, error) {
	cfg := config.NewConfig()
	claim := jwt.MapClaims{}
	claim["user_id"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret := cfg.JWT_SECRET
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil

}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	cfg := config.NewConfig()
	toke, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(cfg.JWT_SECRET), nil
	})
	if err != nil {
		return toke, err
	}
	return toke, nil
}
