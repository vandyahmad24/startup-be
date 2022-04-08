package auth

import (
	"errors"
	cfg "startup/config"
	"startup/config/config"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(id int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	config config.ConfigList
}

// var SECRET_KEY = []byte("VANDY_AHMAD")

func NewService() *jwtService {
	config := cfg.GetConfig()
	secret := config.Config
	return &jwtService{
		config: secret,
	}
}

func (s *jwtService) GenerateToken(id int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret := s.config.Token.Secret.Value
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil

}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	toke, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(s.config.Token.Secret.Value), nil
	})
	if err != nil {
		return toke, err
	}
	return toke, nil
}
