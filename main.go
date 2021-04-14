package main

import (
	"fmt"
	"gggin/routers"
	"github.com/gin-gonic/gin"
)

func main() {



	// 启动handler
	//handlers.NewHandle(cfg1)

	// 设置日志
	//logfile, err := os.Create("./gin_http.log")
	//if err != nil {
	//	fmt.Println("Could not create log file")
	//}
	gin.SetMode(gin.DebugMode)
	//gin.DefaultWriter = io.MultiWriter(logfile)

	// 设置路由
	r := routers.SetUpRouter()

	// 启动
	if err := r.Run(":9999"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
