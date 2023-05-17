package controllers

import (
	"gen/services"
	"github.com/gin-gonic/gin"
)

type deviceValueController struct {
	*Controller
	*services.DeviceValueService
}

var DeviceValueController = deviceValueController{
	Controller:         BaseController,
	DeviceValueService: services.NewDeviceValueService(),
}

// GetDeviceValuesById  根据id返回最新的7条数据
func (r deviceValueController) GetDeviceValuesById(ctx *gin.Context) {
	id := ctx.Query("deviceId")
	devievalues, err := r.GetAllDeviceValue(id)

	if err != nil {
		r.Failed(ctx, Failed, err.Error())
	} else {
		r.Success(ctx, "ok", devievalues)
	}
	return
}
