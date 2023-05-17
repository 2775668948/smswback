package controllers

import (
	"gen/services"
	"github.com/gin-gonic/gin"
)

type locationController struct {
	*Controller
	*services.LocationService
}

var LocationController = locationController{
	Controller:      BaseController,
	LocationService: services.NewLocationService(),
}

// GetLoca 获取位置信息
func (r locationController) GetLoca(ctx *gin.Context) {
	locas, err := r.LocationService.GetNewLocation()
	if err != nil {
		r.Failed(ctx, Failed, err.Error())
	} else {
		r.Success(ctx, "ok", locas)
	}
	return
}
