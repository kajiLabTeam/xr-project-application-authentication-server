package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kajiLabTeam/xr-project-application-authentication-server/handlers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	handlers.AuthUserHandler(r)
	handlers.CreateUserHandler(r)

	handlers.AuthApplicationHandler(r)
	handlers.RegisterApplicationHandler(r)

	r.Run(":8003")
}
