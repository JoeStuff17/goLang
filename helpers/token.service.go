package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func GenerateToken(userId string, role string, name string) string {
	secretKey := []byte(os.Getenv("JWT_KEY"))
	claims := AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
		},
		ID:   userId,
		Role: role,
		Name: name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(secretKey)
	return ss
}

func GenerateRefreshToken(userId string) string {
	secretKey := []byte(os.Getenv("JWT_KEY"))
	//claims := AuthClaims{
	//	RegisteredClaims: jwt.RegisteredClaims{
	//		ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
	//		Issuer:    "TeknoPlay",
	//		Subject:   "Refresh",
	//	},
	//	ID: userId,
	//}
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 360).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(secretKey)
	return ss
}
