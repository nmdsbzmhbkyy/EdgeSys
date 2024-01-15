// ==========================================================================
// 生成日期：2024-01-11 23:59:35 +0800 CST
// 生成路径: temp/ac/router/ac_person_device.go
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

func InitAcPersonDeviceRouter(container *restful.Container) {
	s := &api.AcPersonDeviceApi{
		AcPersonDeviceApp: services.AcPersonDeviceModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/ac/personDevice").Produces(restful.MIME_JSON)
    	tags := []string{"personDevice"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Device分页列表").Handle(s.GetAcPersonDeviceList)
    }).
    	Doc("获取Device分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Device信息").Handle(s.GetAcPersonDevice)
    }).
    	Doc("获取Device信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.AcPersonDevice{}). // on the response
    	Returns(200, "OK", entity.AcPersonDevice{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Device信息").Handle(s.InsertAcPersonDevice)
    }).
    	Doc("添加Device信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcPersonDevice{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Device信息").Handle(s.UpdateAcPersonDevice)
    }).
    	Doc("修改Device信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcPersonDevice{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Device信息").Handle(s.DeleteAcPersonDevice)
    }).
    	Doc("删除Device信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Device信息").Handle(s.ExportAcPersonDevice)
    }).
    	Doc("导出Device信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Device信息").Handle(s.ImportAcPersonDevice)
        }).
        Doc("导入Device信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.AcPersonDevice{}))

    container.Add(ws)
}

