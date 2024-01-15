// ==========================================================================
// 生成日期：2024-01-11 23:56:35 +0800 CST
// 生成路径: temp/ac/router/base_staff_jurisdiction.go
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

func InitBaseStaffJurisdictionRouter(container *restful.Container) {
	s := &api.BaseStaffJurisdictionApi{
		BaseStaffJurisdictionApp: services.BaseStaffJurisdictionModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/ac/staffJurisdiction").Produces(restful.MIME_JSON)
    	tags := []string{"staffJurisdiction"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Jurisdiction分页列表").Handle(s.GetBaseStaffJurisdictionList)
    }).
    	Doc("获取Jurisdiction分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Jurisdiction信息").Handle(s.GetBaseStaffJurisdiction)
    }).
    	Doc("获取Jurisdiction信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseStaffJurisdiction{}). // on the response
    	Returns(200, "OK", entity.BaseStaffJurisdiction{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Jurisdiction信息").Handle(s.InsertBaseStaffJurisdiction)
    }).
    	Doc("添加Jurisdiction信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseStaffJurisdiction{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Jurisdiction信息").Handle(s.UpdateBaseStaffJurisdiction)
    }).
    	Doc("修改Jurisdiction信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseStaffJurisdiction{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Jurisdiction信息").Handle(s.DeleteBaseStaffJurisdiction)
    }).
    	Doc("删除Jurisdiction信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Jurisdiction信息").Handle(s.ExportBaseStaffJurisdiction)
    }).
    	Doc("导出Jurisdiction信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Jurisdiction信息").Handle(s.ImportBaseStaffJurisdiction)
        }).
        Doc("导入Jurisdiction信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseStaffJurisdiction{}))

    container.Add(ws)
}

