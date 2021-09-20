package middleware

import "net/http"

type AuthMiddle_2Middleware struct {
}

func NewAuthMiddle_2Middleware() *AuthMiddle_2Middleware {
	return &AuthMiddle_2Middleware{}
}

func (m *AuthMiddle_2Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
