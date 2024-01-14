package router

import (
	"goweb/docs"
	"goweb/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/", service.GetIndex)

	r.GET("/index", service.GetIndex)

	r.GET("/table", service.GetTable)

	r.GET("/rule", service.GetRule)

	r.GET("/analyze", service.GetAnalyze)

	r.GET("/coffee", service.GetCoffee)

	r.GET("/details", service.GetAddressDetails)

	r.GET("/adduser", service.CreateUser)

	r.GET("/BatchQuery", service.BatchQuery)

	r.GET("/deleteUser", service.DeleteUser)

	r.POST("/updateUser", service.BindAddress)

	r.POST("/findUserByEmailAndPwd", service.FindUserByEmailAndPwd)

	r.GET("/survey", service.Survey)

	r.GET("/update", service.Update)

	return r
}
