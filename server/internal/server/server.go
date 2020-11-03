package server

import (
	"github.com/gin-gonic/gin"
	"{{ variable "module" }}/internal/versions"
)

func Start() error {

	router := gin.Default()

	router.GET("/version", func(context *gin.Context) {
		context.JSON(200, gin.H{"version": versions.Version(), "commit": versions.Commit()})
	})

	return router.Run()

}
