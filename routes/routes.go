package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yodfhafx/go-crud/config"
	"github.com/yodfhafx/go-crud/controllers"
	"github.com/yodfhafx/go-crud/middleware"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	v1 := r.Group("/api/v1")
	authenticate := middleware.Authenticate().MiddlewareFunc()
	authorize := middleware.Authorize()

	authGroup := v1.Group("auth")
	authController := controllers.Auth{DB: db}
	{
		authGroup.POST("/sign-up", authController.Signup)
		authGroup.POST("/sign-in", middleware.Authenticate().LoginHandler)
		authGroup.GET("/profile", authenticate, authController.GetProfile)
		authGroup.PATCH("/profile", authenticate, authController.UpdateProfile)
	}

	usersController := controllers.Users{DB: db}
	usersGroup := v1.Group("users")
	usersGroup.Use(authenticate, authorize)
	{
		usersGroup.GET("", usersController.FindAll)
		usersGroup.POST("", usersController.Create)
		usersGroup.GET("/:id", usersController.FindOne)
		usersGroup.PATCH("/:id", usersController.Update)
		usersGroup.DELETE("/:id", usersController.Delete)
		usersGroup.PATCH("/:id/promote", usersController.Promote)
		usersGroup.PATCH("/:id/demote", usersController.Demote)
	}

	articleController := controllers.Articles{DB: db}
	articlesGroup := v1.Group("articles")
	articlesGroup.GET("", articleController.FindAll)
	articlesGroup.GET("/:id", articleController.FindOne)
	articlesGroup.Use(authenticate, authorize)
	{
		articlesGroup.PATCH("/:id", articleController.Update)
		articlesGroup.DELETE("/:id", articleController.Delete)
		articlesGroup.POST("", authenticate, articleController.Create)
	}

	categoryController := controllers.Categories{DB: db}
	categoriesGroup := v1.Group("categories")
	categoriesGroup.GET("", categoryController.FindAll)
	categoriesGroup.GET("/:id", categoryController.FindOne)
	categoriesGroup.Use(authenticate, authorize)
	{
		categoriesGroup.PATCH("/:id", categoryController.Update)
		categoriesGroup.DELETE("/:id", categoryController.Delete)
		categoriesGroup.POST("", categoryController.Create)
	}

	dashboardController := controllers.Dashboard{DB: db}
	dashboardGroup := v1.Group("dashboard")
	dashboardGroup.Use(authenticate, authorize)
	{
		dashboardGroup.GET("", dashboardController.GetInfo)
	}
}
