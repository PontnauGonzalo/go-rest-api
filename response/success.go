package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func OK(msg string, data interface{}) Response {
	return success(msg, data, http.StatusOK)
}
func Created(msg string, data interface{}) Response {
	return success(msg, data, http.StatusCreated)
}
func Accepted(msg string, data interface{}) Response {
	return success(msg, data, http.StatusAccepted)
}

func success(msg string, data interface{}, code int) Response {
	return &SuccessResponse{
		Message: msg,
		Status:  code,
		Data:    data,
	}
}

func (s *SuccessResponse) Error() string {
	return ""
}
func (s *SuccessResponse) StatusCode() int {
	return s.Status
}
func (s *SuccessResponse) GetBody() ([]byte, error) {
	return json.Marshal(s)
}
func (s *SuccessResponse) GetData() interface{} {
	return s.Data
}