package jwt

import (
	"fmt"
	"time"

	"github.com/DmitriyGoryntsev/marketplace/internal/models"
	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	SecretKey              string
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
}

func NewJWTService(secretKey string, accessExpiry, refreshExpiry time.Duration) *JWTService {
	return &JWTService{
		SecretKey:              secretKey,
		AccessTokenExpiration:  accessExpiry,
		RefreshTokenExpiration: refreshExpiry,
	}
}

func (s *JWTService) GenerateAccessToken(claims models.AccessTokenClaims) (string, error) {
	claims.ExpiresAt = time.Now().Add(s.AccessTokenExpiration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.SecretKey))
}

func (s *JWTService) GenerateRefreshToken(userID int) (string, error) {
	claims := models.RefreshTokenClaims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.RefreshTokenExpiration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.SecretKey))
}

func (s *JWTService) ValidateAccessToken(tokenString string) (*models.AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&models.AccessTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(s.SecretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*models.AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (s *JWTService) ValidateRefreshToken(tokenString string) (*models.RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&models.RefreshTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(s.SecretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*models.RefreshTokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
