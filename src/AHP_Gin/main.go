package main

import (
	"demoProject/src/AHP_Gin/global"
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

}

func main() {
	f, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.Routers()
	r.Run(":" + global.ServerSetting.HttpPort)
}
