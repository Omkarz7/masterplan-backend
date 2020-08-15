package handlers

import (
	"fmt"
	"masterplan-backend/database"
	"masterplan-backend/jwt"
	"masterplan-backend/models"
	"masterplan-backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMasterplan() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, err := jwt.GetValueInToken(ctx.GetHeader("Authorization"), "username") //get username from token
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"Error": fmt.Sprint("Error while retrieving username from token", err)})
			return
		}

		masterplanList, err := database.GetMasterplanFromDB(username) //get marterplan from DB assigned to that username
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"Error": fmt.Sprint("Error while retrieving masterplan", err)})
			return
		}

		sortByStartDate, err := strconv.ParseBool(ctx.Query("sortByStartDate")) //get Query pararm to check if we need to stort over Start Date
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"Error": fmt.Sprint("Error while fetching query parameter", err)})
			return
		}

		if sortByStartDate {
			masterplanList = service.StartDateMergeSort(masterplanList)
		} else {
			masterplanList = service.WBSMergeSort(masterplanList)
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
	}
}
