// ==========================================================================
// 生成日期：2024-01-11 23:06:12 +0800 CST
// 生成路径: temp/base/router/base_unit.go
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

func InitBaseUnitRouter(container *restful.Container) {
	s := &api.BaseUnitApi{
		BaseUnitApp: services.BaseUnitModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/base/unit").Produces(restful.MIME_JSON)
    	tags := []string{"unit"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Unit分页列表").Handle(s.GetBaseUnitList)
    }).
    	Doc("获取Unit分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Unit信息").Handle(s.GetBaseUnit)
    }).
    	Doc("获取Unit信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseUnit{}). // on the response
    	Returns(200, "OK", entity.BaseUnit{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Unit信息").Handle(s.InsertBaseUnit)
    }).
    	Doc("添加Unit信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseUnit{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Unit信息").Handle(s.UpdateBaseUnit)
    }).
    	Doc("修改Unit信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseUnit{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Unit信息").Handle(s.DeleteBaseUnit)
    }).
    	Doc("删除Unit信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Unit信息").Handle(s.ExportBaseUnit)
    }).
    	Doc("导出Unit信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Unit信息").Handle(s.ImportBaseUnit)
        }).
        Doc("导入Unit信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseUnit{}))

    container.Add(ws)
}

