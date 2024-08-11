package auth

import (
	"github.com/ThailanTec/challenger/pousada/src/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userID uuid.UUID, cfg config.Config) (string, error) {
	expirationTime := time.Now().Add(time.Duration(cfg.JWTExpirationMinutes) * time.Minute)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string, cfg config.Config) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	cfg := &config.Config{}
	println(cfg.JWTSecret)
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("Invalid signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(cfg.JWTSecret), nil
	})
}
