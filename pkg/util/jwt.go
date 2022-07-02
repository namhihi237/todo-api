package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(id int, email string) (string, error) {
	var err error
	var token string

	err = godotenv.Load()
	if err != nil {
		return "", err
	}

	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	claims := &Claims{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	var err error
	err = godotenv.Load()
	if err != nil {
		return nil, err
	}

	jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
