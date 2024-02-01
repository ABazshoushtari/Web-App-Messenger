package helpers

import (
	"errors"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserClaim struct {
	jwt.RegisteredClaims
	ID     uint64
	Expiry time.Time
}

func GenerateJWT(userID uint64) (string, error) {
	key := config.GetConfig().JWTSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{},
		ID:               userID,
		Expiry:           time.Now().Add(1 * time.Hour),
	})
	s, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return s, nil
}
func ParseJWT(jwtToken string) (*UserClaim, error) {
	userClaim := UserClaim{}
	key := config.GetConfig().JWTSecret
	parsedToken, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}
	return &userClaim, nil
}
