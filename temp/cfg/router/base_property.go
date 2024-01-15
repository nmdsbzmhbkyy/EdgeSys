// ==========================================================================
// 生成日期：2024-01-12 00:41:50 +0800 CST
// 生成路径: temp/cfg/router/base_property.go
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

func InitBasePropertyRouter(container *restful.Container) {
	s := &api.BasePropertyApi{
		BasePropertyApp: services.BasePropertyModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/cfg/property").Produces(restful.MIME_JSON)
    	tags := []string{"property"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Property分页列表").Handle(s.GetBasePropertyList)
    }).
    	Doc("获取Property分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Property信息").Handle(s.GetBaseProperty)
    }).
    	Doc("获取Property信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseProperty{}). // on the response
    	Returns(200, "OK", entity.BaseProperty{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Property信息").Handle(s.InsertBaseProperty)
    }).
    	Doc("添加Property信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseProperty{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Property信息").Handle(s.UpdateBaseProperty)
    }).
    	Doc("修改Property信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseProperty{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Property信息").Handle(s.DeleteBaseProperty)
    }).
    	Doc("删除Property信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Property信息").Handle(s.ExportBaseProperty)
    }).
    	Doc("导出Property信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Property信息").Handle(s.ImportBaseProperty)
        }).
        Doc("导入Property信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseProperty{}))

    container.Add(ws)
}

