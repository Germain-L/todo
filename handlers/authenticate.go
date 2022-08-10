package handlers

import (
	"context"
	"encoding/json"
	"gotodo/auth"
	"gotodo/models"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

// TODO: Fix panic when using UserKey

type Key string

const emailKey Key = Key("email")

// middleware validating bearer JWT token
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// finds token in headers
		c, err := r.Cookie("token")

		// returns error if no token provided
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("No tokens")
			return
		}

		// gets value from token
		tokenString := c.Value

		// initialise a blank claim
		tokenClaim := &models.Claims{}

		// use jwt package to parse token and return error
		tkn, err := jwt.ParseWithClaims(tokenString, tokenClaim, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtKey, nil
		})

		// checks for errors
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode("Invalid jwt signature")
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}

		// adds email to context, so that the next handler can grab and use username
		newctx := context.WithValue(r.Context(), "email", tokenClaim.Email)

		// goes to the next handler
		next.ServeHTTP(w, r.WithContext(newctx))
	})
}
