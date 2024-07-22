package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

/* обертка для jwt */
type JwtWrapper struct {
	SecretKey         string // подпись jwt
	Issuer            string
	ExpirationMinutes int64
	ExpirationHours   int64
}

/* требование к токену */
type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

/* генерирует jwt токен */
func (j *JwtWrapper) GenerateToken(email string) (signedToken string, err error) {
	claim := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(j.ExpirationMinutes)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}
	return
}

func (j *JwtWrapper) RefreshToken(email string) (signedToken string, err error) {
	claim := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}
	return

}

func (j *JwtWrapper) ValidationToken(signed string) (claim *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(signed, &JwtClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return
	}
	claim, ok := token.Claims.(*JwtClaim)
	if !ok {

		return nil, errors.New("не удалось спарсить токен")
	}

	if claim.ExpiresAt < time.Now().Local().Unix() {

		return nil, errors.New("jwt неактуален")
	}
	return
}
