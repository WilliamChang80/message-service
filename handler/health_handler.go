package handler

import (
	"net/http"

	"github.com/williamChang80/message-service/dto/response"
	"github.com/williamChang80/message-service/responder"
)

func GetHealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responder.Success(w, response.SuccessResponse())
	}
}
