package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yodfhafx/go-crud/routes"
)

func main() {
	r := gin.Default()
	routes.Serve(r)
	r.Run()
}
