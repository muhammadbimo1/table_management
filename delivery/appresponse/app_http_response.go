package appresponse

type AppHttpResponse interface {
	SendData(message *ResponseMessage)
	SendError(errMessage ErrorMessage)
}
