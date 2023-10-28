package open_api

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func FileHandler(filepath string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, filepath)
	}
}

func RegisterSwagger(server *rest.Server) {
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/api/v1/swagger",
		Handler: FileHandler("docs/swagger/main.json")},
	)
}
