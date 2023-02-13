package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"description,omitempty"`
	Data interface{} `json:"data,omitempty"`     
}

func (res Response) Send(w http.ResponseWriter) {
	w.WriteHeader(res.Code)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte("Error When Encode Response"))
	}
}

func New(code int, message interface{}, data ...interface{}) *Response {
	if code > 300 {
		return &Response{
			Code: code,
			Status:  getStatus(code),
			Message: message,
		}
	}
	return &Response{
		Code: code,
		Status:  getStatus(code),
		Message: message,
		Data: data[0],
	}
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unautorized"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	case 304:
		desc = "Not Modified"
	default:
		desc = ""
	}

	return desc
}
