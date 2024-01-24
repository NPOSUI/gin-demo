package v1

import (
	"npos-demo/pkg/app"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Azure Storage Account File Share directories
// @Produce Nil
// @Param   Nil
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router  /api/v1/dirs [post]
func HelloWorld(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, 200, "Hello World!")
}
