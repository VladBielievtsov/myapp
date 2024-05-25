package middlewares

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"my-app/services"
	"my-app/types"
	"my-app/utils"
	"net/http"
	"strings"
)

type AuthContext struct {
	User types.User
}

var userService = services.NewUserService()

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "Authorization header missing"})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "Invalid Authorization header format"})
			return
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return []byte("secret"), nil
		})

		if err != nil {
			utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
			return
		}

		if !token.Valid {
			utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "Error parsing token claims", "token": tokenString})
			return
		}

		userID, ok := claims["sub"].(string)
		if !ok {
			utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "Error parsing user ID from token claims"})
			return
		}

		user, errGetUser := userService.GetUserByID(userID)
		if err != nil {
			utils.JSONResponse(w, http.StatusUnauthorized, errGetUser)
			return
		}

		ctx := context.WithValue(r.Context(), AuthContext{}, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func IsUnauthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) == 2 && strings.ToLower(tokenParts[0]) == "bearer" {
				tokenString := tokenParts[1]

				token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, errors.New("invalid signing method")
					}
					return []byte("secret"), nil
				})

				if token != nil && token.Valid {
					utils.JSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "You are already logged in"})
					return
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
