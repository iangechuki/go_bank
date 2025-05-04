package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iangechuki/go_bank/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T){
	maker,err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token,err := maker.CreateToken(username,duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload,err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotEmpty(t, payload.ID)
	require.Equal(t, payload.Username, username)
	require.WithinDuration(t, payload.IssuedAt.Time, issuedAt, time.Second)
	require.WithinDuration(t, payload.ExpiresAt.Time, expiredAt, time.Second)
}

func TestExpiredToken(t *testing.T){
	maker,err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := -time.Minute

	token,err := maker.CreateToken(username,duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload,err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T){
	payload,err := NewPayload(util.RandomOwner(),time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone,payload)
	token,err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker,err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	verifiedPayload,err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, verifiedPayload)

}
