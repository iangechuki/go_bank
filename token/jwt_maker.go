package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct{
	secretKey string
}

func NewJWTMaker(secretKey string)(Maker,error){
	if len(secretKey) < minSecretKeySize {
		return nil,fmt.Errorf("invalid key size: must be at least %d characters",minSecretKeySize)
	}
	return &JWTMaker{secretKey},nil
}

func (maker *JWTMaker)CreateToken(username string,duration time.Duration)(string,error){

	payload,err := NewPayload(username,duration)
	
	if err != nil {
		return "",err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}
func (maker *JWTMaker)VerifyToken(tokenStr string)(*Payload,error){
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC);!ok {
			// return nil,fmt.Errorf("unexpected signing method: %v",token.Header["alg"])
			return nil,ErrInvalidToken
		}
		return []byte(maker.secretKey),nil
	}
	
	token,err := jwt.ParseWithClaims(tokenStr,&Payload{},keyFunc)
	if err != nil {
		if errors.Is(err,jwt.ErrTokenExpired) {
			return nil,ErrExpiredToken
		}
		return nil,ErrInvalidToken
	}
	payload,ok := token.Claims.(*Payload)
	if !ok {
        // This should ideally not happen if ParseWithClaims succeeded without error
		return nil, fmt.Errorf("%w: could not assert token claims to Payload type", ErrInvalidToken)
	}
	return payload, nil
}