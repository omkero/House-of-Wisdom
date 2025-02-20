package main

import (
	"log"
	"runtime"
	"wisdom/src/config"
	"wisdom/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Use all available CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error initializing enviorment variables")
	}

	// run database pool
	config.InitDB()

	ctx := gin.Default() // gin.Default() for Debugging and gin.New() for production

	// open public dir
	ctx.Static("/assets", "public/static/assets")

	ctx.Use(func(ctx *gin.Context) {
		corsMiddleware := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"}, // Allow specific origins
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Authorization", "Content-Type"},
			AllowCredentials: true,
		})
		corsMiddleware.HandlerFunc(ctx.Writer, ctx.Request)
		ctx.Next()
	})

	// load all routes

	routes.Router(ctx)

	// run the server
	ctx.Run()
}
