package router

import (
	"github.com/gorilla/mux"
	"github.com/williamChang80/message-service/handler"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods(GET).Path("/health").Handler(handler.GetHealthHandler())

	return router
}
