package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type article struct {
	Title string `json:title binding:required`
	Body  string `json:body binding:required`
}


func Serve(r *gin.Engine) {

	articles := []article{
		{Title: "Title#1", Body: "Body#1"},
		{Title: "Title#2", Body: "Body#2"},
		{Title: "Title#3", Body: "Body#3"},
		{Title: "Title#4", Body: "Body#4"},
		{Title: "Title#5", Body: "Body#5"},
	}

	articleGroup := r.Group("api/v1/article")
	articleGroup.GET("", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{"articles": articles})
	})
}