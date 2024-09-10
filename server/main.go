/**
 * Main Function
 */

package main

import (
	"VideoHubGo/utils/LogUtils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	router "VideoHubGo/router"
)

func main() {
	//读取配置文件 - Read The Configuration File
	path, err := os.Getwd()
	if err != nil {
		LogUtils.Logger(err.Error())
	}
	config := viper.New()
	config.AddConfigPath(path + "/configs")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	//尝试进行配置读取 - Try Reading Configuration
	if err := config.ReadInConfig(); err != nil {
		LogUtils.Logger(err.Error())
	}
	port := config.GetString("http.port")

	//Gin服务启动 - Running Gin Service
	r := gin.Default()
	r = router.Router(r)
	err = r.Run(":" + port)
	if err != nil {
		LogUtils.Logger("服务启动失败 - Service Start Fail ：" + err.Error())
	}
}
