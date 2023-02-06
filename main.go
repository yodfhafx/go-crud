package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yodfhafx/go-crud/config"
	"github.com/yodfhafx/go-crud/migrations"
	"github.com/yodfhafx/go-crud/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	migrations.Migrate()

	r := gin.Default()
	r.Static("uploads/", "./uploads")

	uploadDirs := [...]string{"articles", "users"}
	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755)
	}

	routes.Serve(r)
	r.Run(":" + os.Getenv("PORT"))
}
