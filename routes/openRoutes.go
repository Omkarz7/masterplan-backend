package routes

import (
	"masterplan-backend/handlers"

	"github.com/gin-gonic/gin"
)

func OpenRoutes(open *gin.RouterGroup) {
	open.POST("/login", handlers.Login())
}
