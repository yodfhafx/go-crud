package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yodfhafx/go-crud/controllers"
)

func Serve(r *gin.Engine) {
	articlesGroup := r.Group("/api/v1/articles")
	articleController := controllers.Articles{}
	{
		articlesGroup.GET("", articleController.FindAll)
		articlesGroup.GET("/:id", articleController.FindOne)
		articlesGroup.POST("", articleController.Create)
	}
}
