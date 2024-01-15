// ==========================================================================
// 生成日期：2024-01-12 00:04:08 +0800 CST
// 生成路径: temp/ac/router/ac_cert_device.go
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

func InitAcCertDeviceRouter(container *restful.Container) {
	s := &api.AcCertDeviceApi{
		AcCertDeviceApp: services.AcCertDeviceModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/ac/certDevice").Produces(restful.MIME_JSON)
    	tags := []string{"certDevice"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Device分页列表").Handle(s.GetAcCertDeviceList)
    }).
    	Doc("获取Device分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Device信息").Handle(s.GetAcCertDevice)
    }).
    	Doc("获取Device信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.AcCertDevice{}). // on the response
    	Returns(200, "OK", entity.AcCertDevice{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Device信息").Handle(s.InsertAcCertDevice)
    }).
    	Doc("添加Device信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcCertDevice{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Device信息").Handle(s.UpdateAcCertDevice)
    }).
    	Doc("修改Device信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcCertDevice{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Device信息").Handle(s.DeleteAcCertDevice)
    }).
    	Doc("删除Device信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Device信息").Handle(s.ExportAcCertDevice)
    }).
    	Doc("导出Device信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Device信息").Handle(s.ImportAcCertDevice)
        }).
        Doc("导入Device信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.AcCertDevice{}))

    container.Add(ws)
}

