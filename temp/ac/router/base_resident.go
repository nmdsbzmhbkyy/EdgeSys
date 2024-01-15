// ==========================================================================
// 生成日期：2024-01-11 23:53:55 +0800 CST
// 生成路径: temp/ac/router/base_resident.go
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

func InitBaseResidentRouter(container *restful.Container) {
	s := &api.BaseResidentApi{
		BaseResidentApp: services.BaseResidentModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/ac/resident").Produces(restful.MIME_JSON)
    	tags := []string{"resident"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Resident分页列表").Handle(s.GetBaseResidentList)
    }).
    	Doc("获取Resident分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Resident信息").Handle(s.GetBaseResident)
    }).
    	Doc("获取Resident信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseResident{}). // on the response
    	Returns(200, "OK", entity.BaseResident{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Resident信息").Handle(s.InsertBaseResident)
    }).
    	Doc("添加Resident信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseResident{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Resident信息").Handle(s.UpdateBaseResident)
    }).
    	Doc("修改Resident信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseResident{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Resident信息").Handle(s.DeleteBaseResident)
    }).
    	Doc("删除Resident信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Resident信息").Handle(s.ExportBaseResident)
    }).
    	Doc("导出Resident信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Resident信息").Handle(s.ImportBaseResident)
        }).
        Doc("导入Resident信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseResident{}))

    container.Add(ws)
}

