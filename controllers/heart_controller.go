package controllers

import (
	"gen/services"
	"github.com/gin-gonic/gin"
)

type heartController struct {
	*Controller
	*services.HeartService
}

var HeartController = heartController{
	Controller:   BaseController,
	HeartService: services.NewHeartService(),
}

// GetHeart 获取最新的心率数值
func (r heartController) GetHeart(ctx *gin.Context) {
	hears, err := r.HeartService.GetNewHeartValue()
	if err != nil {
		r.Failed(ctx, Failed, err.Error())
	} else {
		r.Success(ctx, "ok", hears)
	}
	return
}
