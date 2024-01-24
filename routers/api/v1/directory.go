package v1

import (
	"net/http"

	"npos-demo/pkg/app"
	"npos-demo/pkg/msg"

	"github.com/gin-gonic/gin"
)

// @Summary Create Azure Storage Account File Share directories
// @Produce Nil
// @Param   Nil
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router  /api/v1/dirs [post]
func CreateDirectories(c *gin.Context) {
	appG := app.Gin{C: c}

	if err := app.CreateDirectories(); err != nil {
		appG.Response(http.StatusInternalServerError, msg.ERROR, err)
		return
	}

	appG.Response(http.StatusOK, msg.SUCCESS, nil)
}
