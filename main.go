package main

import (
	"TestBE/database"
	"TestBE/pkg/mysql"
	"TestBE/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	//Adding Gin Gonic to Release Mode (if you want to use debug mode use gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	mysql.DatabaseConn()
	database.MigrationDB()

	routes.Routes(r.Group("/api/v1/"))

	fmt.Println("Running on port 5065")

	r.Run(":5065")
}
