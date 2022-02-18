package appresponse

import (
	"log"
	"net/http"
)

type ResponseMessage struct {
	Status      string      `json:"status"`
	Description string      `json:"message"`
	Data        interface{} `json:"data"`
}

type ErrorMessage struct {
	Errcode          int    `json:"code"`
	ErrorDescription string `json:"message"`
}

func NewResponseMessage(status string, description string, data interface{}) *ResponseMessage {
	return &ResponseMessage{
		status, description, data,
	}
}

func NewUnauthorizedError(err error, message string) *ErrorMessage {
	em := &ErrorMessage{
		Errcode:          http.StatusUnauthorized,
		ErrorDescription: message,
	}
	log.Println(err.Error())
	return em
}

func NewInternalServerError(err error, message string) *ErrorMessage {
	em := &ErrorMessage{
		Errcode:          http.StatusInternalServerError,
		ErrorDescription: message,
	}
	log.Println(err.Error())
	return em
}

func NewBadRequestError(err error, message string) *ErrorMessage {
	em := &ErrorMessage{
		Errcode:          http.StatusBadRequest,
		ErrorDescription: message,
	}
	log.Println(err.Error())
	return em
}
