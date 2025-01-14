package api

import "github.com/gin-gonic/gin"



func SetupProductAPI(r *gin.Engine) {
	productAPI := r.Group("/api/v2")
	{
		productAPI.GET("/product", getProduct)
		productAPI.POST("/product", createProduct)
	}
}

func getProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "get product"})
}

func createProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "create product"})
}

	
	
	