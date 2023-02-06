package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yodfhafx/go-crud/config"
	"github.com/yodfhafx/go-crud/controllers"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	articlesGroup := r.Group("/api/v1/articles")
	articleController := controllers.Articles{DB: db}
	{
		articlesGroup.GET("", articleController.FindAll)
		articlesGroup.GET("/:id", articleController.FindOne)
		articlesGroup.POST("", articleController.Create)
	}
}
