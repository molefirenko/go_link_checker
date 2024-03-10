package main

import (
	"github.com/gin-gonic/gin"
	"github.com/molefirenko/go_link_checker/controllers"
)

func main() {
	rest := gin.Default()
	routesList(rest)
	rest.Run()
}

func routesList(rest *gin.Engine) {
	rest.POST("/process", controllers.ProcessLinks)
}
