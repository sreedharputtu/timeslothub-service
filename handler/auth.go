package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type JwtWrapper struct {
	SecretKey         string
	Issuer            string
	ExpirationMinutes int64
	ExpirationHours   int64
}

type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

func (j *JwtWrapper) GenerateToken(email string) (signedToken string, err error) {
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(j.ExpirationMinutes)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}
	return
}

func (j *JwtWrapper) RefreshToken(email string) (signedtoken string, err error) {
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedtoken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}
	return
}

func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}
	return
}

func Authz() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header from the request

		var clientToken interface{}
		session := sessions.Default(c)
		clientToken = session.Get("state")

		fmt.Println("client token:", clientToken)

		if clientToken == nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		// Create a JwtWrapper with the secret key and issuer
		jwtWrapper := JwtWrapper{
			SecretKey: "verysecretkey",
			Issuer:    "AuthService",
		}
		// Validate the token
		claims, err := jwtWrapper.ValidateToken(clientToken.(string))
		if err != nil {
			c.Redirect(301, "/login")
			return
		}
		// Set the claims in the context
		c.Set("email", claims.Email)
		// Continue to the next handler
		c.Next()
	}
}
