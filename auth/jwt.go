package auth

import (
	"gotodo/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GetNewToken(email string) (tokenString string, expirationTime time.Time, err error) {
	// Declare the expiration time of the token, token available for 5 minutes
	expirationTime = time.Now().Add(5 * time.Minute)

	// Register the new token claims, StandardClaims is comming from jwt pac
	claims := &models.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err = token.SignedString(JwtKey)

	return tokenString, expirationTime, err
}
