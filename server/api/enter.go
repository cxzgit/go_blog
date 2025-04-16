package api

import "server/service"

/*
*
API层
1.定义 ApiGroup 组合了各功能 API，比如 BaseApi。
2.暴露唯一实例 ApiGroupApp，供路由层调用。
3.baseService 是 Service 层的实例，供 API 层执行业务逻辑。
*/
type ApiGroup struct {
	BaseApi
}

var ApiGroupApp = new(ApiGroup)

var baseService = service.ServiceGroupApp.BaseService
