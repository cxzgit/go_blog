package service

/*
*Service层
将不同业务服务（BaseService、EsService）组合在一起，暴露单例 ServiceGroupApp。
*/
type ServiceGroup struct {
	BaseService
	EsService
}

var ServiceGroupApp = new(ServiceGroup)
