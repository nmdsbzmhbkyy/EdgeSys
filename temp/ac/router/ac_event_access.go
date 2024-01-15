// ==========================================================================
// 生成日期：2024-01-12 00:12:57 +0800 CST
// 生成路径: temp/ac/router/ac_event_access.go
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

func InitAcEventAccessRouter(container *restful.Container) {
	s := &api.AcEventAccessApi{
		AcEventAccessApp: services.AcEventAccessModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/ac/eventAccess").Produces(restful.MIME_JSON)
    	tags := []string{"eventAccess"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Access分页列表").Handle(s.GetAcEventAccessList)
    }).
    	Doc("获取Access分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Access信息").Handle(s.GetAcEventAccess)
    }).
    	Doc("获取Access信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.AcEventAccess{}). // on the response
    	Returns(200, "OK", entity.AcEventAccess{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Access信息").Handle(s.InsertAcEventAccess)
    }).
    	Doc("添加Access信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcEventAccess{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Access信息").Handle(s.UpdateAcEventAccess)
    }).
    	Doc("修改Access信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcEventAccess{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Access信息").Handle(s.DeleteAcEventAccess)
    }).
    	Doc("删除Access信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Access信息").Handle(s.ExportAcEventAccess)
    }).
    	Doc("导出Access信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Access信息").Handle(s.ImportAcEventAccess)
        }).
        Doc("导入Access信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.AcEventAccess{}))

    container.Add(ws)
}

