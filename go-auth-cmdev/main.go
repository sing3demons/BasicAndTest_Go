package main

import (
	"app/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/images", "./uploaded/images")
	api.Setup(r)
	r.Run()
}