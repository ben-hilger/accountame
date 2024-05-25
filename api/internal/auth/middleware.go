package auth

import (
	"context"
	"net/http"
	"strings"
)

type middlewareService interface {
	ValidateAuthenticationToken(tokenString string) (string, error)
}

type Middleware struct {
	authService middlewareService
}

func NewMiddleware(authService middlewareService) Middleware {
	return Middleware{authService: authService}
}

func (m Middleware) Protect(handler func(response http.ResponseWriter, request *http.Request)) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		authorizationHeader := request.Header.Get("Authorization")
		token := extractTokenFromHeaders(authorizationHeader)
		if token == "" {
			http.Error(response, "unauthorized", http.StatusUnauthorized)
			return
		}
		userId, err := m.authService.ValidateAuthenticationToken(token)
		if err != nil {
			http.Error(response, "unauthorized", http.StatusUnauthorized)
			return
		}

		updatedContext := context.WithValue(request.Context(), "userId", userId)

		protectedRoute := http.HandlerFunc(handler)
		protectedRoute.ServeHTTP(response, request.WithContext(updatedContext))
	})
}

func extractTokenFromHeaders(authorizationHeader string) string {
	splitToken := strings.Split(authorizationHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}

	return splitToken[1]
}
