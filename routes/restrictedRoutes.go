package routes

import (
	"masterplan-backend/handlers"

	"github.com/gin-gonic/gin"
)

func RestrictedRoutes(res *gin.RouterGroup) {
	res.GET("/getMasterplan", handlers.GetMasterplan())
}
