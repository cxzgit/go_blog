package initialize

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/global"
	"server/router"
)

/**
初始化层
1.根据配置初始化 Gin 引擎和模式。
2.加载并启用 session 中间件。
3.挂载静态文件目录。
4.调用路由层的 InitBaseRouter 完成各个 endpoint 的注册。
*/

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 设置 gin 模式（debug、release 等）
	gin.SetMode(global.Config.System.Env)
	//创建默认的 gin 引擎，包含 Logger、Recovery 中间件
	Router := gin.Default()
	// 使用gin会话路由
	// 配置基于 Cookie 的 session 存储
	var store = cookie.NewStore([]byte(global.Config.System.SessionsSecret))
	Router.Use(sessions.Sessions("session", store))
	// 将指定目录下的文件提供给客户端
	// "uploads" 是URL路径前缀，http.Dir("uploads")是实际文件系统中存储文件的目录
	Router.StaticFS(global.Config.Upload.Path, http.Dir(global.Config.Upload.Path))
	// 创建路由组（取得路由组实例（router层入口））
	routerGroup := router.RouterGroupApp

	//按全局配置的前缀创建分组
	publicGroup := Router.Group(global.Config.System.RouterPrefix)
	{
		//将base路由注册到publicGroup下
		routerGroup.InitBaseRouter(publicGroup)
	}
	return Router
}
