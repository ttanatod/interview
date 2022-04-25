package routers

import (
	"example.com/m/v2/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetCollectionRoutes(router *gin.RouterGroup, db *gorm.DB) {
	ctrls := controllers.DBController{Database: db}

	router.GET("member", ctrls.GetAllMember)
	router.POST("member", ctrls.RegisterMember)

	router.GET("field", ctrls.GetAllField)
	router.POST("field", ctrls.RegisterField)

	router.POST("rent", ctrls.RentField)
}
