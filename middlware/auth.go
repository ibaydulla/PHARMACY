package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("mysecretkey")

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "no token", 401)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid token", 401)
			return
		}

		// role
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(w, "invalid claims", 401)
			return
		}

		role, ok := claims["role"].(string)

		if !ok {
			http.Error(w, "role not found", 401)
			return
		}

		if role != "admin" {
			http.Error(w, "forbidden", 403)
			return
		}

		next.ServeHTTP(w, r)
	})
}
