package handlers

import (
	"fmt"
	"masterplan-backend/database"
	"masterplan-backend/jwt"
	"masterplan-backend/models"
	"masterplan-backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMasterplan() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, err := jwt.GetValueInToken(ctx.GetHeader("Authorization"), "username")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"Error": fmt.Sprint("Error while retrieving username from token", err)})
			return
		}
		masterplanList, err := database.GetMasterplanFromDB(username)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"Error": fmt.Sprint("Error while retrieving masterplan", err)})
			return
		}
		err = service.GeneratExcelFile(masterplanList)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"Error": fmt.Sprint("Error while retrieving masterplan", err)})
			return
		}
		fileName := models.Config.MasterplanFilename + ".csv"
		//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attachment; filename="+fileName)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(fileName)
		// ctx.Status(http.StatusOK)
	}
}
