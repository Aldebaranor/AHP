package routers

import (
	"demoProject/src/AHP_Gin/handlers/ahp"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	routers := gin.Default()
	ahpRouter := routers.Group("/demoProject")
	{
		ahpRouter.POST("/weight", ahp.GetWeight)
		ahpRouter.POST("/points", ahp.GetPoints)
	}
	return routers
}
