package main

import (
	"TestBE/database"
	"TestBE/pkg/mysql"
	"TestBE/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	mysql.DatabaseConn()
	database.MigrationDB()

	routes.Routes(r.Group("/api/v1/"))

	fmt.Println("Running on port 5065")

	r.Run(":5065")
}
