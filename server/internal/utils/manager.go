package utils

import "github.com/gin-gonic/gin"

func RegisterHandlers(r *gin.Engine) {
	// Swagger
	//docs.SwaggerInfo.BasePath = "/api"
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	registerBaseHandler(r)
	registerAdminHandler(r)
	registerBlogHandler(r)
}

// 通用接口: 全部不需要 登录 + 鉴权
func registerBaseHandler(r *gin.Engine) {
	base := r.Group("/api/v1")
	base.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// TODO: 登录, 注册 记录日志
	//base.POST("/login", userAuthAPI.Login)          // 登录
	//base.POST("/register", userAuthAPI.Register)    // 注册
	//base.GET("/logout", userAuthAPI.Logout)         // 退出登录
	//base.POST("/report", blogInfoAPI.Report)        // 上报信息
	//base.GET("/config", blogInfoAPI.GetConfigMap)   // 获取配置
	//base.PATCH("/config", blogInfoAPI.UpdateConfig) // 更新配置
	//base.GET("/email/verify", userAuthAPI.VerifyCode)
}

// 后台管理系统的接口: 全部需要 登录 + 鉴权
func registerAdminHandler(r *gin.Engine) {
	//auth := r.Group("/api/v1/admin")

}

// 前台的接口: 大部分不需要登录, 部分需要登录
func registerBlogHandler(r *gin.Engine) {
	//front := r.Group("/api/v1/front")
}
