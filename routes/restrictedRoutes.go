package routes

import (
	"masterplan-backend/handlers"

	"github.com/gin-gonic/gin"
)

func RestrictedRoutes(res *gin.RouterGroup) {
	res.POST("/getMasterplan", handlers.GetMasterplan())
}
