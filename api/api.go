package api

import (
	apiv1 "github.com/danhngocphh/fptda2/api/v1.0"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api)

	}
}
