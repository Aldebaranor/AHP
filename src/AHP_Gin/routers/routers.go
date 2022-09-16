package routers

import (
	"demoProject/src/AHP_Gin/handlers/ahp"
	"demoProject/src/AHP_Gin/handlers/tree"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	routers := gin.Default()
	ahpRouter := routers.Group("/ahp")
	{
		ahpRouter.POST("/weight", ahp.GetWeight)
		ahpRouter.POST("/tree", tree.GetLeaves)
		ahpRouter.POST("/points", ahp.GetPoints)
	}
	return routers
}
