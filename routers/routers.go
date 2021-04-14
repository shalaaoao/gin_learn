package routers

import (
	"gggin/config"
	"gggin/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	// 导入配置文件
	var cfg1 config.Config
	config.ParseConfig("./config/myconf.json", &cfg1)

	// 启动handlers
	baseH := handlers.NewHandle(cfg1)

	r.GET("/hello", baseH.HelloHandler)
	r.GET("/firsttopic", baseH.FirstTopic)
	r.GET("/topiclist", baseH.TopicList)
	r.GET("/firstrds", baseH.FirstRds)
	return r
}
