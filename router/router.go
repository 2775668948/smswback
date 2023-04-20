package router

import (
	"gen/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", controllers.BaseController.Index)
	r := server.Group("/api")
	{
		v1 := r.Group("/v1")
		{
			deviceCtrl := controllers.DeviceController
			v1.Group("/devicesws").
				GET("", deviceCtrl.GetAll)
			v1.Group("/changesw").
				GET("", deviceCtrl.UpdateSwState)

			autoTaskCtrl := controllers.AutoTaskController
			v1.Group("/autotask").
				GET("", autoTaskCtrl.GetAllAutoTask)
			v1.Group("/changeau").
				GET("", autoTaskCtrl.UpdateAutoItemShowState)
			v1.Group("/newau").
				POST("", autoTaskCtrl.AddAutoItem)
			v1.Group("/checkauto").
				GET("", autoTaskCtrl.FindAutoItemByDoTime)

			deviceValueCtrl := controllers.DeviceValueController
			v1.Group("devicevalue").
				GET("", deviceValueCtrl.GetDeviceValuesById)
		}
	}
}
