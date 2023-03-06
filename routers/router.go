package routers

import (
	"app/actions"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Load() *gin.Engine {
	r.GET("/status", actions.Status)
	return r
}

func init() {
	r = gin.Default()
}
