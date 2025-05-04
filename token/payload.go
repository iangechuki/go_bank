package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)
type Payload struct {
	jwt.RegisteredClaims
	Username string `json:"username"`

}
func NewPayload(username string,duration time.Duration)(*Payload,error){
	tokenID,err := uuid.NewRandom()
	if err != nil {
		return nil,err
	}
	now := time.Now()
	payload := &Payload{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
		Username: username,
	}
	return payload,nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt.Time) {
		return ErrExpiredToken
	}
	return nil
}