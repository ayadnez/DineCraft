package main

import (
	"fmt"
	"os"

	"PROJECT_1/database"
	"PROJECT_1/middleware"
	"PROJECT_1/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.openCollection(database.Client, "food")

func main() {
	fmt.Println("hello world")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())
	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderItemRoutes(router)
	routes.IvoiceRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)

	router.Run(":" + port)

}
