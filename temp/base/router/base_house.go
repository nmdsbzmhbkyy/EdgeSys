// ==========================================================================
// 生成日期：2024-01-11 23:09:55 +0800 CST
// 生成路径: temp/base/router/base_house.go
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

func InitBaseHouseRouter(container *restful.Container) {
	s := &api.BaseHouseApi{
		BaseHouseApp: services.BaseHouseModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/base/house").Produces(restful.MIME_JSON)
    	tags := []string{"house"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取House分页列表").Handle(s.GetBaseHouseList)
    }).
    	Doc("获取House分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取House信息").Handle(s.GetBaseHouse)
    }).
    	Doc("获取House信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseHouse{}). // on the response
    	Returns(200, "OK", entity.BaseHouse{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加House信息").Handle(s.InsertBaseHouse)
    }).
    	Doc("添加House信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseHouse{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改House信息").Handle(s.UpdateBaseHouse)
    }).
    	Doc("修改House信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseHouse{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除House信息").Handle(s.DeleteBaseHouse)
    }).
    	Doc("删除House信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出House信息").Handle(s.ExportBaseHouse)
    }).
    	Doc("导出House信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入House信息").Handle(s.ImportBaseHouse)
        }).
        Doc("导入House信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseHouse{}))

    container.Add(ws)
}

