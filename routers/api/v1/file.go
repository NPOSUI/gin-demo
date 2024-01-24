package v1

import (
	"npos-demo/pkg/app"
	"npos-demo/pkg/msg"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Delete heapdump files when they expire
// @Produce nil
// @Param   Nil
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router  /api/v1/del-files [post]
func DeleteFiles(c *gin.Context) {
	appG := app.Gin{C: c}

	if err := app.DeleteFiles(); err != nil {
		appG.Response(http.StatusInternalServerError, msg.ERROR, err)
		return
	}

	appG.Response(http.StatusOK, msg.SUCCESS, nil)
}
