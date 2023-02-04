package helpers

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Meta Meta `json:"meta"`
    Data 
}

type Meta struct{
    Status      string      `json:"status"`
    IsError     bool        `json:"isError"`
    Message interface{} `json:"description,omitempty"`
}

type Data struct{
    Data interface{} `json:"data,omitempty"`
}

func (res Response) Send(w http.ResponseWriter) {
    err := json.NewEncoder(w).Encode(res)
    if err != nil {
        w.Write([]byte("Error When Encode Response"))
    }
}

func New(code int, message interface{},data ...interface{}) *Response {
    return &Response{
        Meta{
            Status: getStatus(code),
            Message: message,
        },
        Data{
            Data: data,
        },
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
