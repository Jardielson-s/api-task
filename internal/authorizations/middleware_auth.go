package authorizations

import (
	"net/http"

	"github.com/Jardielson-s/api-task/internal/authenticate"
)

func findValueInSlice(slice []interface{}, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func hasAnyRole(userRoles []interface{}, requiredRoles []string) bool {
	roleMap := make(map[string]struct{})
	for _, role := range userRoles {
		if roleStr, ok := role.(string); ok {
			roleMap[roleStr] = struct{}{}
		}
	}
	for _, requiredRole := range requiredRoles {
		if _, exists := roleMap[requiredRole]; exists {
			return true
		}
	}
	return false
}

func MiddlewareAuth(requiredRole *[]string, permission string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenInfo, _ := r.Context().Value("tokenInfo").(authenticate.TokenInfo)
		if hasAnyRole(tokenInfo.Roles, *requiredRole) {
			if findValueInSlice(tokenInfo.Permissions, permission) {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusForbidden)
	})
}

func ApplyMiddlewares(roles []string, permission string, handlerFunc http.HandlerFunc) http.Handler {
	handler := handlerFunc
	return authenticate.
		ProtectedHandler(
			MiddlewareAuth(
				&roles, permission,
				http.HandlerFunc(handler),
			),
		)
}
