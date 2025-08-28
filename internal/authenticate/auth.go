package authenticate

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Jardielson-s/api-task/configs"
	"github.com/golang-jwt/jwt/v5"
)

func ProtectedHandler(next http.Handler) http.Handler {
	configs.Envs()

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
		permissionsInterface := []interface{}{}
		if v, ok := claims["permissions"].([]interface{}); ok {
			permissionsInterface = v
		}
		rolesInterface := []interface{}{}
		if v, ok := claims["roles"].([]interface{}); ok {
			rolesInterface = v
		}
		// permissionsInterface := []interface{}{}
		// permissions := claims["permissions"]
		// if permissions != nil {
		// 	permissionsInterface = claims["permissions"].([]interface{})
		// }
		// rolesInterface := claims["roles"].([]interface{})
		// roles := claims["roles"]
		// if roles != nil {
		// 	rolesInterface = claims["roles"].([]interface{})
		// }
		tokenInfo := TokenInfo{
			Username:    claims["username"].(string),
			Email:       claims["email"].(string),
			ID:          int(claims["id"].(float64)),
			Permissions: permissionsInterface,
			Roles:       rolesInterface,
		}

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "tokenInfo", tokenInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
