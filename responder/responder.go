package responder

import (
	"encoding/json"
	"net/http"

	"github.com/williamChang80/message-service/dto/response"
	pkgError "github.com/williamChang80/message-service/error"
)

func Error(w http.ResponseWriter, err error, errorStatus int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorStatus)

	json.NewEncoder(w).Encode(&response.BaseResponse{
		Status:  errorStatus,
		Message: err.Error(),
	})
}

func NewErrorReponse(w http.ResponseWriter, err error) {
	switch pkgError.Cause(err).(type) {
	case *pkgError.NotFoundError:
		Error(w, err, http.StatusNotFound)
	case *pkgError.BadRequestError:
		Error(w, err, http.StatusBadRequest)
	case *pkgError.UnprocessableEntityError:
		Error(w, err, http.StatusUnprocessableEntity)
	default:
		Error(w, err, http.StatusInternalServerError)
	}
}

func Success(w http.ResponseWriter, successResponse interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(successResponse)
}
