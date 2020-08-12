package middleware

import (
	"bytes"
	"fmt"
	"masterplan-backend/jwt"
	"masterplan-backend/logs"
	"masterplan-backend/routes"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//InitMiddleware adds logging, JWT checks and route grouping for handling requets
func InitMiddleware(router *gin.Engine) {
	router.Use(ginBodyLogMiddleware())

	open := router.Group("o")
	routes.OpenRoutes(open)

	restricted := router.Group("r")
	restricted.Use(RestrictedRequestMiddleware())
	routes.RestrictedRoutes(restricted)
}

//RestrictedRequestMiddleware verifies JWT token
func RestrictedRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authToken := c.GetHeader("Authorization")
		if strings.TrimSpace(authToken) == "" {
			c.AbortWithStatusJSON(401, gin.H{"Error": fmt.Sprint("Invalid API token")})
			return
		}
		_, err := jwt.VerifyToken(authToken)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"Error": fmt.Sprintln("Invalid API token", err)})
			return
		}

		c.Next()
	}
}

//Middleware that handles logging of requests and errors

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginBodyLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path
		if path == "/o/time" {
			return
		}
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		end := time.Now().UTC()
		latency := end.Sub(start)

		//pass the data to channel to be logged
		if c.Writer.Status() == 200 {
			logs.HTTPlog <- logrus.Fields{
				"status":     c.Writer.Status(),
				"method":     c.Request.Method,
				"path":       path,
				"ip":         c.ClientIP(),
				"duration":   latency,
				"user_agent": c.Request.UserAgent(),
				"error":      nil,
			}
		} else {
			logs.HTTPlog <- logrus.Fields{
				"status":     c.Writer.Status(),
				"method":     c.Request.Method,
				"path":       path,
				"ip":         c.ClientIP(),
				"duration":   latency,
				"user_agent": c.Request.UserAgent(),
				"error":      blw.body.String(),
			}
		}
	}
}
