package main

import (
	"fmt"
	"net/http"
	"selling/database"
	"selling/pkg/mysql"
	"selling/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	mysql.DatabaseInit()
	database.RunMigration()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	routes.RouteInit(r.Group("/api/v1"))

	fmt.Println("server started")
	http.ListenAndServe("localhost:5000", r)
}