package router

/*
路由层
1.定义了一个 RouterGroup 结构体，内嵌 BaseRouter，并对外暴露唯一实例 RouterGroupApp。
2.通过组合可以在未来扩展更多路由组（如 UserRouter、AdminRouter 等）
*/
type RouterGroup struct {
	BaseRouter
}

var RouterGroupApp = new(RouterGroup)
