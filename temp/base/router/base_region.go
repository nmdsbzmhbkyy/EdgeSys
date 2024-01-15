// ==========================================================================
// 生成日期：2024-01-11 23:15:00 +0800 CST
// 生成路径: temp/base/router/base_region.go
// 生成人：aurine
// ==========================================================================
package router

import (
    "mod.miligc.com/edge-common/CommonKit/model"
    "mod.miligc.com/edge-common/CommonKit/restfulx"
	"EdgeSys/temp/base/api"
	"EdgeSys/temp/base/entity"
	"EdgeSys/temp/base/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
    "github.com/emicklei/go-restful/v3"
)

func InitBaseRegionRouter(container *restful.Container) {
	s := &api.BaseRegionApi{
		BaseRegionApp: services.BaseRegionModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/base/region").Produces(restful.MIME_JSON)
    	tags := []string{"region"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Region分页列表").Handle(s.GetBaseRegionList)
    }).
    	Doc("获取Region分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Region信息").Handle(s.GetBaseRegion)
    }).
    	Doc("获取Region信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseRegion{}). // on the response
    	Returns(200, "OK", entity.BaseRegion{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Region信息").Handle(s.InsertBaseRegion)
    }).
    	Doc("添加Region信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseRegion{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Region信息").Handle(s.UpdateBaseRegion)
    }).
    	Doc("修改Region信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseRegion{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Region信息").Handle(s.DeleteBaseRegion)
    }).
    	Doc("删除Region信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Region信息").Handle(s.ExportBaseRegion)
    }).
    	Doc("导出Region信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Region信息").Handle(s.ImportBaseRegion)
        }).
        Doc("导入Region信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseRegion{}))

    container.Add(ws)
}

