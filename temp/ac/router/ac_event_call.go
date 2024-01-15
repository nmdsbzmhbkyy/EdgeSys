// ==========================================================================
// 生成日期：2024-01-12 00:13:10 +0800 CST
// 生成路径: temp/ac/router/ac_event_call.go
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

func InitAcEventCallRouter(container *restful.Container) {
	s := &api.AcEventCallApi{
		AcEventCallApp: services.AcEventCallModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/ac/eventCall").Produces(restful.MIME_JSON)
    	tags := []string{"eventCall"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Call分页列表").Handle(s.GetAcEventCallList)
    }).
    	Doc("获取Call分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Call信息").Handle(s.GetAcEventCall)
    }).
    	Doc("获取Call信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.AcEventCall{}). // on the response
    	Returns(200, "OK", entity.AcEventCall{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Call信息").Handle(s.InsertAcEventCall)
    }).
    	Doc("添加Call信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcEventCall{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Call信息").Handle(s.UpdateAcEventCall)
    }).
    	Doc("修改Call信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcEventCall{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Call信息").Handle(s.DeleteAcEventCall)
    }).
    	Doc("删除Call信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Call信息").Handle(s.ExportAcEventCall)
    }).
    	Doc("导出Call信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Call信息").Handle(s.ImportAcEventCall)
        }).
        Doc("导入Call信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.AcEventCall{}))

    container.Add(ws)
}

