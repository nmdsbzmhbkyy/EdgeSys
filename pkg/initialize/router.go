package initialize

import (
	"EdgeSys/apps/job/jobs"
	"EdgeSys/pkg/global"
	"EdgeSys/pkg/transport"

	businessRouter "EdgeSys/apps/business/router"
	devRouter "EdgeSys/apps/develop/router"
	jobRouter "EdgeSys/apps/job/router"
	logRouter "EdgeSys/apps/log/router"
	sysRouter "EdgeSys/apps/system/router"
	"EdgeSys/pkg/middleware"
)

func InitRouter() *transport.HttpServer {
	// server配置
	serverConfig := global.Conf.Server
	server := transport.NewHttpServer(serverConfig.GetPort())

	container := server.Container
	// 防止XSS
	container.Filter(middleware.EscapeHTML)
	// 是否允许跨域
	if serverConfig.Cors {
		container.Filter(middleware.Cors(container).Filter)
		container.Filter(container.OPTIONSFilter)
	}
	// 流量限制
	if serverConfig.Rate.Enable {
		container.Filter(middleware.Rate)
	}
	// 设置路由组
	{
		sysRouter.InitSystemRouter(container)
		sysRouter.InitOrganizationRouter(container)
		sysRouter.InitConfigRouter(container)
		sysRouter.InitApiRouter(container)
		sysRouter.InitDictRouter(container)
		sysRouter.InitMenuRouter(container)
		sysRouter.InitRoleRouter(container)
		sysRouter.InitPostRouter(container)
		sysRouter.InitUserRouter(container)
		sysRouter.InitNoticeRouter(container)
		//本地图片上传接口
		sysRouter.InitUploadRouter(container)
	}
	//日志系统
	{
		logRouter.InitOperLogRouter(container)
		logRouter.InitLoginLogRouter(container)
	}
	// 代码生成
	{
		devRouter.InitGenTableRouter(container)
		devRouter.InitGenRouter(container)
	}
	{
		jobRouter.InitJobRouter(container)
		jobRouter.InitJobLogRouter(container)
	}
	{
		businessRouter.InitEventRouter(container)
	}
	// api接口
	middleware.SwaggerConfig(container)
	// 开启调度任务
	go func() {
		jobs.InitJob()
	}()

	return server
}
