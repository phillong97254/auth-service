package routes

import (
	"net/http"

	v1 "github.com/phillong97254/auth-service/internal/handler/rest/v1"
)

func V1Routes(mux *http.ServeMux, h v1.Handler) {
	v1PublicRoutes(mux, h)
}
