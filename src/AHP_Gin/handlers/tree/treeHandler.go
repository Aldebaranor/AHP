package tree

import (
	"demoProject/src/AHP_Gin/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLeaves(context *gin.Context) {
	node := &model.TreeNode{}
	context.ShouldBind(&node)
	node.PreOrder()
	context.JSON(http.StatusOK, "success")
}
