package global

import (
	"demoProject/src/AHP_Gin/pkg/setting"
	"log"
)

// 和yml对应
type ServerSettingS struct {
	RunMode  string
	HttpPort string
}

// 定义全局变量
var (
	ServerSetting *ServerSettingS
)

// 读取配置到全局便量
func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	log.Printf("setting:")
	log.Printf("%+v", ServerSetting)
	return nil
}
