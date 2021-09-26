package main

import (
	"go-module/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.Setup(router)

	router.Run(":8080")
}
