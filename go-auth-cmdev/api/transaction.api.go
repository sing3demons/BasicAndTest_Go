package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func SetupTransactionAPI(r *gin.Engine)  {
	transactionAPI := r.Group("/api/v2")
	{
		transactionAPI.GET("/transaction", getTransaction)
		transactionAPI.POST("/transaction", createTransaction)
	}
}


func getTransaction(c *gin.Context) {
	c.String(http.StatusOK, "List Transaction")
}

func createTransaction(c *gin.Context) {
	c.String(http.StatusOK, "Create Transaction")
}