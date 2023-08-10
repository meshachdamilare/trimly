package middleware

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github/meshachdamilare/trimly/settings/config"
	"github/meshachdamilare/trimly/settings/constant"
	"time"
)

type SignClaims struct {
	Email string
	Role  string
	jwt.RegisteredClaims
}

var key = config.Config.Secret_Key

func CreateToken(t time.Duration, id, email, role string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, SignClaims{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    constant.TokenIssuer,
			ID:        id,
		},
	}).SignedString([]byte(key))

	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}
	return token, nil
}

func ValidateToken(tokenString string) (*SignClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, &SignClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if iss, err := token.Claims.GetIssuer(); iss != constant.TokenIssuer || err != nil {
			return nil, fmt.Errorf("unknown issuer: %v", token.Header["iss"])
		}
		return []byte(key), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token expired: %w", err)
		}
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(*SignClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}
	return claims, nil
}
