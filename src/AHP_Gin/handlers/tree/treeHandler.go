package tree

import (
	"demoProject/src/AHP_Gin/model"
	"github.com/gin-gonic/gin"
)

func GetLeaves(context *gin.Context) {
	node := &model.TreeNode{}
	context.ShouldBind(&node)
	node.PreOrder()
}
