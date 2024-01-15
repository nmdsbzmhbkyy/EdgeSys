// ==========================================================================
// 生成日期：2024-01-11 23:25:46 +0800 CST
// 生成路径: temp/base/router/base_staff_position.go
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

func InitBaseStaffPositionRouter(container *restful.Container) {
	s := &api.BaseStaffPositionApi{
		BaseStaffPositionApp: services.BaseStaffPositionModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/base/staffPosition").Produces(restful.MIME_JSON)
    	tags := []string{"staffPosition"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Position分页列表").Handle(s.GetBaseStaffPositionList)
    }).
    	Doc("获取Position分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Position信息").Handle(s.GetBaseStaffPosition)
    }).
    	Doc("获取Position信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseStaffPosition{}). // on the response
    	Returns(200, "OK", entity.BaseStaffPosition{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Position信息").Handle(s.InsertBaseStaffPosition)
    }).
    	Doc("添加Position信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseStaffPosition{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Position信息").Handle(s.UpdateBaseStaffPosition)
    }).
    	Doc("修改Position信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseStaffPosition{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Position信息").Handle(s.DeleteBaseStaffPosition)
    }).
    	Doc("删除Position信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Position信息").Handle(s.ExportBaseStaffPosition)
    }).
    	Doc("导出Position信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Position信息").Handle(s.ImportBaseStaffPosition)
        }).
        Doc("导入Position信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseStaffPosition{}))

    container.Add(ws)
}

