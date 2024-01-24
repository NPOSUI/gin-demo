package routers

import (
	v1 "npos-demo/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		// demo request
		r.POST("/hw", v1.HelloWorld)
		// bulk create Azure SA File Shared directories.
		r.POST("/dirs", v1.CreateDirectories)
		// bulk delete the expired Heapdump files in Azure SA File Shared
		r.POST("/del-files", v1.DeleteFiles)
	}

	return r
}
