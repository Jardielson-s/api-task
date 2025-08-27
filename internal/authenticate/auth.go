package authenticate

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func ProtectedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenInfo := TokenInfo{
			Username: claims["username"].(string),
			Email:    claims["email"].(string),
			ID:       int(claims["id"].(float64)),
			// Permission: []
		}
		fmt.Println(tokenInfo)

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println(tokenInfo)
		ctx := context.WithValue(r.Context(), "tokenInfo", tokenInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
