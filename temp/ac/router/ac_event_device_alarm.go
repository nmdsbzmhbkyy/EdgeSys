// ==========================================================================
// 生成日期：2024-01-12 00:12:41 +0800 CST
// 生成路径: temp/ac/router/ac_event_device_alarm.go
// 生成人：aurine
// ==========================================================================
package router

import (
    "mod.miligc.com/edge-common/CommonKit/model"
    "mod.miligc.com/edge-common/CommonKit/restfulx"
	"EdgeSys/temp/ac/api"
	"EdgeSys/temp/ac/entity"
	"EdgeSys/temp/ac/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
    "github.com/emicklei/go-restful/v3"
)

func InitAcEventDeviceAlarmRouter(container *restful.Container) {
	s := &api.AcEventDeviceAlarmApi{
		AcEventDeviceAlarmApp: services.AcEventDeviceAlarmModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/ac/eventDeviceAlarm").Produces(restful.MIME_JSON)
    	tags := []string{"eventDeviceAlarm"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Alarm分页列表").Handle(s.GetAcEventDeviceAlarmList)
    }).
    	Doc("获取Alarm分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Alarm信息").Handle(s.GetAcEventDeviceAlarm)
    }).
    	Doc("获取Alarm信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.AcEventDeviceAlarm{}). // on the response
    	Returns(200, "OK", entity.AcEventDeviceAlarm{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Alarm信息").Handle(s.InsertAcEventDeviceAlarm)
    }).
    	Doc("添加Alarm信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcEventDeviceAlarm{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Alarm信息").Handle(s.UpdateAcEventDeviceAlarm)
    }).
    	Doc("修改Alarm信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcEventDeviceAlarm{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Alarm信息").Handle(s.DeleteAcEventDeviceAlarm)
    }).
    	Doc("删除Alarm信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Alarm信息").Handle(s.ExportAcEventDeviceAlarm)
    }).
    	Doc("导出Alarm信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Alarm信息").Handle(s.ImportAcEventDeviceAlarm)
        }).
        Doc("导入Alarm信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.AcEventDeviceAlarm{}))

    container.Add(ws)
}

