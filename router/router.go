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

			deviceValueCtrl := controllers.DeviceValueController
			v1.Group("devicevalue").
				GET("", deviceValueCtrl.GetDeviceValuesById)

			locationCtrl := controllers.LocationController
			v1.Group("locations").
				GET("", locationCtrl.GetLoca)
			heartCtrl := controllers.HeartController
			v1.Group("hearts").
				GET("", heartCtrl.GetHeart)
			tempCtrl := controllers.TemperatureController
			v1.Group("temps").
				GET("", tempCtrl.GetTemp)
		}
	}
}
