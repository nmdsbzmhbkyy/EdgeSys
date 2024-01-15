// ==========================================================================
// 生成日期：2024-01-12 00:41:39 +0800 CST
// 生成路径: temp/cfg/router/base_project.go
// 生成人：aurine
// ==========================================================================
package router

import (
    "mod.miligc.com/edge-common/CommonKit/model"
    "mod.miligc.com/edge-common/CommonKit/restfulx"
	"EdgeSys/temp/cfg/api"
	"EdgeSys/temp/cfg/entity"
	"EdgeSys/temp/cfg/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
    "github.com/emicklei/go-restful/v3"
)

func InitBaseProjectRouter(container *restful.Container) {
	s := &api.BaseProjectApi{
		BaseProjectApp: services.BaseProjectModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/cfg/project").Produces(restful.MIME_JSON)
    	tags := []string{"project"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Project分页列表").Handle(s.GetBaseProjectList)
    }).
    	Doc("获取Project分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Project信息").Handle(s.GetBaseProject)
    }).
    	Doc("获取Project信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseProject{}). // on the response
    	Returns(200, "OK", entity.BaseProject{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Project信息").Handle(s.InsertBaseProject)
    }).
    	Doc("添加Project信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseProject{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Project信息").Handle(s.UpdateBaseProject)
    }).
    	Doc("修改Project信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseProject{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Project信息").Handle(s.DeleteBaseProject)
    }).
    	Doc("删除Project信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Project信息").Handle(s.ExportBaseProject)
    }).
    	Doc("导出Project信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Project信息").Handle(s.ImportBaseProject)
        }).
        Doc("导入Project信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseProject{}))

    container.Add(ws)
}

