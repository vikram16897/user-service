package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/user-service/controller"
	"github.com/user-service/database"
	"github.com/user-service/repository"
	"github.com/user-service/routes"
	"github.com/user-service/service"
)

func main() {
	dsn := os.Getenv("DNS_DOCKER")
	db, err := database.NewDB(dsn)
	if err != nil {
		log.Println("unable to connect database please check connection.")
	}
	repo := repository.NewUserRepository(db)
	service := service.NewUserRepository(repo)
	userController := controller.NewUserController(service)

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()

	routes.RegisterHealthRoute(r)
	routes.RegisterUserRoute(r, userController)

	r.Run(":" + os.Getenv("PORT"))
}
