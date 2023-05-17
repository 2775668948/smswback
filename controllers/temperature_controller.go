package controllers

import (
	"gen/services"
	"github.com/gin-gonic/gin"
)

type temperatureController struct {
	*Controller
	*services.TemperatureService
}

var TemperatureController = temperatureController{
	Controller:         BaseController,
	TemperatureService: services.NewTemperatureService(),
}

// GetTemp 获取最新的体温数值
func (r temperatureController) GetTemp(ctx *gin.Context) {
	temps, err := r.TemperatureService.GetNewTemperature()
	if err != nil {
		r.Failed(ctx, Failed, err.Error())
	} else {
		r.Success(ctx, "ok", temps)
	}
	return
}
