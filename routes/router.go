package routes

import (
	"github.com/felixlambertv/nextalent/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	personController := &controller.PersonController{DB: db}

	route.GET("/GetCountry/:name", personController.GetCountryByPersonName)
	route.GET("/GetCurrentTime", controller.GetCurrentTime)

	return route
}
