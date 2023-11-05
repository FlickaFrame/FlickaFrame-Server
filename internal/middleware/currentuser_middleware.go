package middleware

import (
	"github.com/zeromicro/go-zero/rest/handler"
	"net/http"
)

type CurrentUserMiddleware struct {
	secret string
}

func NewCurrentUserMiddleware(secret string) *CurrentUserMiddleware {
	return &CurrentUserMiddleware{
		secret: secret,
	}
}

func (m *CurrentUserMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) > 0 {
			//has jwt Authorization
			authHandler := handler.Authorize(m.secret)
			authHandler(next).ServeHTTP(w, r)
		}
		next(w, r)
	}
}
