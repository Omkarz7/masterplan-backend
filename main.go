package main

import (
	"masterplan-backend/database"
	"masterplan-backend/logs"
	"masterplan-backend/middleware"
	"masterplan-backend/models"
	"masterplan-backend/reader"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	err := reader.LoadConfiguration() // load configuration settings
	if err != nil {
		panic(err.Error())
	}
	go logs.LogrusHTTP() //turn on the logger to logs the requests

	database.ConnectToDatabases() //open DB connections
	//Close them when program ends because the connections are supposed to be long lived.
	//Go handles the connections in pool automatically
	defer database.DBconn.Close()

	middleware.InitMiddleware(router) //add middleware, grouping and request handlers

	serve := &http.Server{
		Addr:         models.Config.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	serve.ListenAndServe()
}
