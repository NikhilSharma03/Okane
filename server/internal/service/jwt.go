package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// The JWTService interface defines methods needs to implement
type JWTService interface {
	GenerateJWT(id, email string) (string, error)
}

// The jWTService struct implements interface methods
type jWTService struct {
	lg *log.Logger
}

// The NewJWTService returns jwtService struct which implement JWTService interface
func NewJWTService(lg *log.Logger) JWTService {
	return &jWTService{lg: lg}
}

// GenerateJWT takes id and email and returns JWT token
func (js *jWTService) GenerateJWT(id, email string) (string, error) {
	// Get SecretKey  from .env
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		js.lg.Printf("Jwt Secret key not found")
		return "", fmt.Errorf("jwt secret key not found")
	}
	mySigningKey := []byte(secretKey)

	// Generate Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		js.lg.Printf("failed to sign token. %v", err)
		return "", fmt.Errorf("failed to sign token")
	}

	return tokenString, nil
}
