package main

import (
	"demoProject/src/AHP_Gin/global"
	"demoProject/src/AHP_Gin/model"
	"demoProject/src/AHP_Gin/routers"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func init() {

	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	model.ApiCode.Message = map[uint]string{
		model.ApiCode.SUCCESS:         "计算成功！",
		model.ApiCode.CheckFAILED:     "一致性校验未通过，请重新输入判断矩阵！",
		model.ApiCode.ConditionExceed: "评估指标数目超过上限，请减少指标数量再重新输入！",
	}

}

func main() {
	f, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.Routers()
	r.Run(":" + global.ServerSetting.HttpPort)
}
