package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ksrnnb/chat-app-server/infrastructure/middleware"
	"github.com/ksrnnb/chat-app-server/infrastructure/route"
)

const (
	port = ":9000"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("FRONT_END_URL")}
	config.AllowCredentials = true

	router.Use(cors.New(config))
	router.Use(middleware.NewSessionMiddleware())

	newRouter := route.SetRoute(router)

	newRouter.Run(port)
}
