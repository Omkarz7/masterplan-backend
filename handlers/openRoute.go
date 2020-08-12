package handlers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"masterplan-backend/database"
	"masterplan-backend/jwt"
	"masterplan-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login handles the /login route
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var cred models.Credentials
		err := ctx.Bind(&cred)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		} else {
			passwordHash := sha256.Sum256([]byte(cred.Password))
			cred.PasswordHash = base64.StdEncoding.EncodeToString(passwordHash[:])
			err := database.VerifyCredentials(cred)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
				return
			}
			token, err := jwt.GenerateToken(cred.Username)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"Error": fmt.Sprint("Error while generating token", err)})
				return
			}
			ctx.Header("Authorization", token)
			ctx.Status(http.StatusOK)
		}

	}
}
