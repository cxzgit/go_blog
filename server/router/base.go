package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

/*
*
路由层
1.将 URL 路径（captcha、sendEmailVerificationCode、qqLoginURL）映射到 API 层对应的处理函数。
2.api.ApiGroupApp.BaseApi 是在 API 层定义的控制器实例。
*/
type BaseRouter struct {
}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	baseApi := api.ApiGroupApp.BaseApi
	{
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("sendEmailVerificationCode", baseApi.SendEmailVerificationCode)
		baseRouter.GET("qqLoginURL", baseApi.QQLoginURL)
	}
}
