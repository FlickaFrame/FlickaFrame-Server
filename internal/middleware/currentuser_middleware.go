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
			authHandler := handler.Authorize(m.secret, handler.WithUnauthorizedCallback(
				func(w http.ResponseWriter, r *http.Request, err error) {
					next(w, r) // jwt失效，不做处理
				}))
			authHandler(next).ServeHTTP(w, r)
			return
		} else {
			//no jwt Authorization
			next(w, r)
		}
	}
}
