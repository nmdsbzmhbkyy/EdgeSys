package initialize

import (
	devRouter "EdgeSys/apps/develop/router"
	"EdgeSys/apps/job/jobs"
	jobRouter "EdgeSys/apps/job/router"
	logRouter "EdgeSys/apps/log/router"
	sysRouter "EdgeSys/apps/system/router"
	baseRouter "EdgeSys/temp/base/router"
	acRouter "EdgeSys/temp/ac/router"
	deviceRouter "EdgeSys/temp/device/router"
	cfgRouter "EdgeSys/temp/cfg/router"
	"EdgeSys/pkg/global"
	"EdgeSys/pkg/middleware"
	"EdgeSys/pkg/transport"
	businessRouter "mod.miligc.com/edge-common/business-common/business/router"
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
	// 社区信息管理
	{
		baseRouter.InitBaseBuildingRouter(container)
		baseRouter.InitBaseFrameRouter(container)
		baseRouter.InitBaseHouseRouter(container)
		baseRouter.InitBaseHouseTypeRouter(container)
		baseRouter.InitBaseRegionRouter(container)
		baseRouter.InitBaseStaffDepartmentRouter(container)
		baseRouter.InitBaseStaffPositionRouter(container)
		baseRouter.InitBaseUnitRouter(container)
	}
    // 门禁对讲系统
    {
    	acRouter.InitAcCardRouter(container)
    	acRouter.InitAcCertDeviceRouter(container)
    	acRouter.InitAcEventAccessRouter(container)
    	acRouter.InitAcEventCallRouter(container)
    	acRouter.InitAcEventDeviceAlarmRouter(container)
    	acRouter.InitAcFaceRouter(container)
    	acRouter.InitAcPersonDeviceRouter(container)
    	acRouter.InitBasePersonRouter(container)
    	acRouter.InitBaseResidentRouter(container)
    	acRouter.InitBaseStaffJurisdictionRouter(container)
    }
    // 设备管理
    {
    	deviceRouter.InitAcCertDeviceRouter(container)
    	deviceRouter.InitBaseDeviceRouter(container)
    }
    // 社区基础配置
    {
    	cfgRouter.InitBaseFrameRouter(container)
    	cfgRouter.InitBaseFrameCfgRouter(container)
    	cfgRouter.InitBaseProjectRouter(container)
    	cfgRouter.InitBaseProjectCfgRouter(container)
    	cfgRouter.InitBasePropertyRouter(container)
    }
	// 开启调度任务
	go func() {
		jobs.InitJob()
	}()

	return server
}
