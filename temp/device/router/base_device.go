// ==========================================================================
// 生成日期：2024-01-12 00:25:58 +0800 CST
// 生成路径: temp/device/router/base_device.go
// 生成人：aurine
// ==========================================================================
package router

import (
    "mod.miligc.com/edge-common/CommonKit/model"
    "mod.miligc.com/edge-common/CommonKit/restfulx"
	"EdgeSys/temp/device/api"
	"EdgeSys/temp/device/entity"
	"EdgeSys/temp/device/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
    "github.com/emicklei/go-restful/v3"
)

func InitBaseDeviceRouter(container *restful.Container) {
	s := &api.BaseDeviceApi{
		BaseDeviceApp: services.BaseDeviceModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/device/device").Produces(restful.MIME_JSON)
    	tags := []string{"device"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Device分页列表").Handle(s.GetBaseDeviceList)
    }).
    	Doc("获取Device分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Device信息").Handle(s.GetBaseDevice)
    }).
    	Doc("获取Device信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseDevice{}). // on the response
    	Returns(200, "OK", entity.BaseDevice{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Device信息").Handle(s.InsertBaseDevice)
    }).
    	Doc("添加Device信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseDevice{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Device信息").Handle(s.UpdateBaseDevice)
    }).
    	Doc("修改Device信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseDevice{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Device信息").Handle(s.DeleteBaseDevice)
    }).
    	Doc("删除Device信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Device信息").Handle(s.ExportBaseDevice)
    }).
    	Doc("导出Device信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Device信息").Handle(s.ImportBaseDevice)
        }).
        Doc("导入Device信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseDevice{}))

    container.Add(ws)
}

