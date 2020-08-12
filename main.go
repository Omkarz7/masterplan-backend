package main

import (
	"fmt"
	"masterplan-backend/database"
	"masterplan-backend/logs"
	"masterplan-backend/middleware"
	"masterplan-backend/models"
	"masterplan-backend/reader"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	err := reader.LoadConfiguration() // load configuration settings
	if err != nil {
		fmt.Println(err)
	}
	go logs.LogrusHTTP() //turn on the logger to logs the requests

	database.ConnectToDatabases() //open DB connections
	//Close them when program ends because the connections are supposed to be long lived.
	//Go handles the connections in pool automatically
	defer database.DBconn.Close()
	md := cors.DefaultConfig()

	//Allow configs set to avoid issues when communicating with Vue since this is just a prototype project
	md.AllowAllOrigins = true
	md.AllowHeaders = []string{"*"}
	md.AllowMethods = []string{"*"}
	md.ExposeHeaders = []string{"Authorization"} // allows vue to read authorization token

	router.Use(cors.New(md))          //use the declared configs
	middleware.InitMiddleware(router) //add middleware, grouping and request handlers

	serve := &http.Server{
		Addr:         models.Config.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	serve.ListenAndServe()
}
