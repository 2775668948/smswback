package controllers

import (
	"gen/models"
	"gen/services"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type autotaskController struct {
	*Controller
	*services.AutoTaskService
}

var AutoTaskController = autotaskController{
	Controller:      BaseController,
	AutoTaskService: services.NewAutoTaskService(),
}

// GetAllAutoTask 获取所有的自动化任务
func (r autotaskController) GetAllAutoTask(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	autotasks, err := r.AutoTaskService.GetAllAutoTask(page)
	if err != nil {
		r.Failed(ctx, Failed, err.Error())
	} else {
		r.Success(ctx, "ok", autotasks)
	}
	return
}

// UpdateAutoItemShowState 更改自动化任务启用状态
func (r autotaskController) UpdateAutoItemShowState(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil || id <= 0 {
		r.Failed(ctx, ParamError, "id不能为空")
		return
	}
	autoItem, err := r.UpdateAutoItemState(id)
	if err != nil {
		r.Failed(ctx, Failed, err.Error())
	} else {
		r.Success(ctx, "ok", autoItem)
	}
	return
}

// AddAutoItem 新曾自动化任务
func (r autotaskController) AddAutoItem(ctx *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		DoTime   string `json:"DoTime"`
		DeviceId string `json:"deviceId"`
		DeviceAc uint   `json:"deviceAc"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		r.Failed(ctx, Failed, err.Error())
		return
	}
	autoItem := models.AutoTask{
		Name:      req.Name,
		DoTime:    req.DoTime,
		DeviceId:  req.DeviceId,
		DeviceAc:  req.DeviceAc,
		Show:      1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	r.AddNewAutoItem(autoItem)
	r.Success(ctx, "ok", autoItem)
}

// FindAutoItemByDoTime 根据执行时间去查找自动化任务
func (r autotaskController) FindAutoItemByDoTime(ctx *gin.Context) {
	now := time.Now()
	formatted := now.Format("15:04")
	autoItem, err := r.FindAutoItem(formatted)
	if err != nil {
		r.Failed(ctx, Failed, err.Error())
	} else {
		r.Success(ctx, "ok", autoItem)
	}
	return
}
