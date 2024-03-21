package handler

import (
	"errors"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
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
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
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

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var clientToken interface{}
		session := sessions.Default(c)
		clientToken = session.Get("state")

		if clientToken == nil {
			log.Error("client token is empty")
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
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
			log.Error("Error validating token: ", err)
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		log.Info("auth successfull , email:", claims.Email)
		c.Next()
	}
}
