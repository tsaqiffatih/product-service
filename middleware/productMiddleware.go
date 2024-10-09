package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/tsaqiffatih/product-service/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "fail",
				"error": map[string]interface{}{
					"code":    http.StatusUnauthorized,
					"message": "No token provided",
				},
			})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify token
		claims, err := utils.VerifyJWT(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "fail",
				"error": map[string]interface{}{
					"code":    http.StatusUnauthorized,
					"message": "Invalid token",
				},
			})
			return
		}

		// Token valid, continue with the request
		r.Header.Set("User-Id", claims.Id)
		r.Header.Set("User-Email", claims.Email)
		r.Header.Set("User-Role", claims.Role)
		next.ServeHTTP(w, r)
	})
}
