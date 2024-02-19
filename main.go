package main

import (
	"github.com/extimsu/JsonRestApi/db"
	"github.com/extimsu/JsonRestApi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
