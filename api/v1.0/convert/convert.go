package convert

import (
	"github.com/gin-gonic/gin"
	// "github.com/danhngocphh/fptda2/lib/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	convert := r.Group("/convert")
	{
		// posts.POST("/", middlewares.Authorized, create)
		convert.POST("/", getConvert)
		// posts.GET("/:id", read)
		// posts.DELETE("/:id", middlewares.Authorized, remove)
		// posts.PATCH("/:id", middlewares.Authorized, update)
	}
}
