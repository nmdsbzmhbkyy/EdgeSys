// ==========================================================================
// 生成日期：2024-01-11 23:02:38 +0800 CST
// 生成路径: temp/base/router/base_building.go
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

func InitBaseBuildingRouter(container *restful.Container) {
	s := &api.BaseBuildingApi{
		BaseBuildingApp: services.BaseBuildingModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/base/building").Produces(restful.MIME_JSON)
    	tags := []string{"building"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Building分页列表").Handle(s.GetBaseBuildingList)
    }).
    	Doc("获取Building分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Building信息").Handle(s.GetBaseBuilding)
    }).
    	Doc("获取Building信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseBuilding{}). // on the response
    	Returns(200, "OK", entity.BaseBuilding{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Building信息").Handle(s.InsertBaseBuilding)
    }).
    	Doc("添加Building信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseBuilding{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Building信息").Handle(s.UpdateBaseBuilding)
    }).
    	Doc("修改Building信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseBuilding{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Building信息").Handle(s.DeleteBaseBuilding)
    }).
    	Doc("删除Building信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Building信息").Handle(s.ExportBaseBuilding)
    }).
    	Doc("导出Building信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Building信息").Handle(s.ImportBaseBuilding)
        }).
        Doc("导入Building信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseBuilding{}))

    container.Add(ws)
}

