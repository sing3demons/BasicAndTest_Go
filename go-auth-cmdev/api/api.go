package api

import (
	"app/config"

	"github.com/gin-gonic/gin"
)

// Setup - config

func Setup(r *gin.Engine) {

	config.SetupDB()

	SetupAuthenAPI(r)
	SetupProductAPI(r)
	// SetupTransactionAPI(r)
}
