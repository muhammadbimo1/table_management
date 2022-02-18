package delivery

import (
	"errors"
	"table_management/delivery/appresponse"
	"table_management/dto"
	"table_management/usecase"

	"github.com/gin-gonic/gin"
)

type CustomerTableApi struct {
	usecase     usecase.CustomerTableUseCase
	publicRoute *gin.RouterGroup
}

func NewCustomerTableApi(publicRoute *gin.RouterGroup, useCase usecase.CustomerTableUseCase) *CustomerTableApi {
	api := CustomerTableApi{
		usecase:     useCase,
		publicRoute: publicRoute,
	}
	api.initRouter()
	return &api
}

func (a *CustomerTableApi) initRouter() {
	tableRoute := a.publicRoute.Group("/table")
	tableRoute.GET("", a.getTableList)
	tableRoute.POST("/checkin", a.tableCheckIn)
	tableRoute.PUT("/checkout", a.tableCheckOut)

}

func (a *CustomerTableApi) getTableList(c *gin.Context) {
	tables, err := a.usecase.GetTodayListCustomerTable()
	response := appresponse.NewJsonResponse(c)
	if err != nil {
		response.SendError(*appresponse.NewInternalServerError(err, "Something went wrong."))
		return
	}
	response.SendData(&appresponse.ResponseMessage{Status: "SUCCESS", Description: "list table", Data: tables})
}

func (a *CustomerTableApi) tableCheckIn(c *gin.Context) {
	var checkInRequest dto.CheckInRequest
	response := appresponse.NewJsonResponse(c)
	if err := c.ShouldBindJSON(&checkInRequest); err != nil {
		response.SendError(*appresponse.NewBadRequestError(err, "failed to check in"))
		return
	}
	_, err := a.usecase.TableCheckIn(checkInRequest)
	if err != nil {
		response.SendError(*appresponse.NewBadRequestError(err, err.Error()))
		return
	}
	response.SendData(appresponse.NewResponseMessage("success", "checkin", nil))
}

func (a *CustomerTableApi) tableCheckOut(c *gin.Context) {
	billNo := c.Query("billNo")
	response := appresponse.NewJsonResponse(c)
	if billNo == "" {
		response.SendError(*appresponse.NewBadRequestError(errors.New("field required"), "failed check out"))
		return
	}

	err := a.usecase.TableCheckOut(billNo)
	if err != nil {
		response.SendError(*appresponse.NewInternalServerError(err, "something wrong"))
		return
	}
	response.SendData(&appresponse.ResponseMessage{"success", "checkout", nil})
}
