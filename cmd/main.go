package main

import (
	"github.com/gin-gonic/gin"

	"go-sms-verify/api"
)

func main() {
	router := gin.Default()

	// Init Cnfg
	app := api.Config{Router: router}

	// Routes
	app.Routes()

	router.Run(":8080")
}
