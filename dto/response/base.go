package response

import "net/http"

type BaseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SuccessResponse() BaseResponse {
	return BaseResponse{
		Status:  http.StatusOK,
		Message: "success",
	}
}
