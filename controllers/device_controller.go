package controllers

import (
	"gen/services"
	"gen/zlog"
	"github.com/gin-gonic/gin"
	"strconv"
)

type deviceController struct {
	*Controller
	*services.DeviceService
}

var DeviceController = deviceController{
	Controller:    BaseController,
	DeviceService: services.NewDeviceService(),
}

// GetAll 获取所有的设备开关
func (r deviceController) GetAll(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	devices, err := r.DeviceService.GetAll(page)

	if err != nil {
		r.Failed(ctx, Failed, err.Error())
	} else {
		r.Success(ctx, "ok", devices)
	}
	return
}

// UpdateSwState 修改设备状态开关
func (r deviceController) UpdateSwState(ctx *gin.Context) {
	zlog.Info("UpdateSwState...")
	var DeviceId = ctx.Query("DeviceId")
	zlog.Info("%s", DeviceId)
	if DeviceId == "" {
		r.Failed(ctx, Failed, "DeviceId不能为空")
		return
	}
	newDevice, err := r.DeviceService.UpdateDeviceSw(DeviceId)

	if err != nil {
		r.Failed(ctx, Failed, err.Error())
	} else {
		r.Success(ctx, "ok", newDevice)
	}
	return
}
