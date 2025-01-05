package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NotFound(msg string) Response {
	return errorResponse(msg, http.StatusNotFound)
}
func BadRequest(msg string) Response {
	return errorResponse(msg, http.StatusBadRequest)
}
func Forbidden(msg string) Response {
	return errorResponse(msg, http.StatusForbidden)
}
func Unauthorized(msg string) Response {
	return errorResponse(msg, http.StatusUnauthorized)
}
func NonAuthoritativeInfo(msg string) Response {
	return errorResponse(msg, http.StatusNonAuthoritativeInfo)
}

func InternalServerError(msg string) Response {
	return errorResponse(msg, http.StatusInternalServerError)
}

func errorResponse(msg string, code int) Response {
	return &ErrorResponse{
		Message: msg,
		Status:  code,
	}
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
func (e *ErrorResponse) StatusCode() int {
	return e.Status
}

func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}
func (e *ErrorResponse) GetData() interface{} {
	return nil
}