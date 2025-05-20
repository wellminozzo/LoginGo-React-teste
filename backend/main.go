package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wellminozzo/LoginGo-React-teste/controllers"
	"github.com/wellminozzo/LoginGo-React-teste/db"
)

func main() {
	fmt.Println("Starting the application...")
	db.DbConnection()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	r.GET("/user/:id", controllers.GetUser)
	r.GET("/users", controllers.GetAllUsers)
	r.POST("/user", controllers.CreateUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.POST("/login", controllers.Login)

	r.GET("/protected", controllers.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Acesso autorizado!"})
	})

	r.Run(":5000")
}
