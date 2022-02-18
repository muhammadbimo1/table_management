package appresponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	c *gin.Context
}

func (j *JsonResponse) SendData(message *ResponseMessage) {
	j.c.JSON(http.StatusOK, message)
}

func (j *JsonResponse) SendError(errMessage ErrorMessage) {
	j.c.AbortWithStatusJSON(errMessage.Errcode, errMessage.ErrorDescription)
}

func NewJsonResponse(c *gin.Context) AppHttpResponse {
	return &JsonResponse{c: c}
}
