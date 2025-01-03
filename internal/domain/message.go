package domain

import (
	"fmt"
	"net/http"
)

func MsgResponse(res http.ResponseWriter, status int, message string) {
	res.WriteHeader(status)
	fmt.Fprintf(res, `{"status": %d,"message":"%s"`, status, message)
}

func DataResponse(res http.ResponseWriter, status int, data interface{}) {
	res.WriteHeader(status)
	fmt.Fprintf(res, `{"status": %d,"data":%s`, status, data)
}

func InvalidMethodResponse(res http.ResponseWriter) {
	status := http.StatusNotFound
	res.WriteHeader(status)
	fmt.Fprintf(res, `{"status": %d,"message":"method doesn't exist"}`, status)
}
