package main

import (
	"flag"
	"log"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	"server-api/internal/router"
	"server-api/internal/runtime"
)

func main() {
	configPath := flag.String("c", "../config/config.yml", "配置文件路径")
	flag.Parse()
	// 根据命令行参数读取配置文件, 其他变量的初始化依赖于配置文件对象
	conf := runtime.GetConfig(*configPath)

	// 初始化 gin 服务
	gin.SetMode(conf.Server.Mode)

	r := gin.New()
	err := r.SetTrustedProxies([]string{"0.0.0.0/0"})
	if err != nil {
		log.Fatalf("set trusted proxies: %s\n", err)
	}
	// 开发模式使用 gin 自带的日志和恢复中间件, 生产模式使用自定义的中间件
	if conf.Server.Mode == "debug" {
		r.Use(gin.Logger(), gin.Recovery()) // gin 自带的日志和恢复中间件, 挺好用的
	} else {
		//r.Use(middleware.Recovery(true), middleware.Logger())
	}

	router.RegisterHandlers(r)

	// 使用本地文件上传, 需要静态文件服务
	if conf.Upload.OssType == "local" {
		r.Static(conf.Upload.Path, conf.Upload.StorePath)
	}
	serverAddr := conf.Server.Port
	if serverAddr[0] == ':' || strings.HasPrefix(serverAddr, "0.0.0.0:") {
		log.Printf("Serving HTTP on (http://localhost:%s/) ... \n", strings.Split(serverAddr, ":")[1])
	} else {
		log.Printf("Serving HTTP on (http://%s/) ... \n", serverAddr)
	}
	if err := endless.ListenAndServe(serverAddr, r); err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}
