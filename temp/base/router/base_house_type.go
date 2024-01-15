// ==========================================================================
// 生成日期：2024-01-11 23:01:03 +0800 CST
// 生成路径: temp/base/router/base_house_type.go
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

func InitBaseHouseTypeRouter(container *restful.Container) {
	s := &api.BaseHouseTypeApi{
		BaseHouseTypeApp: services.BaseHouseTypeModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/base/houseType").Produces(restful.MIME_JSON)
    	tags := []string{"houseType"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Type分页列表").Handle(s.GetBaseHouseTypeList)
    }).
    	Doc("获取Type分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Type信息").Handle(s.GetBaseHouseType)
    }).
    	Doc("获取Type信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseHouseType{}). // on the response
    	Returns(200, "OK", entity.BaseHouseType{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Type信息").Handle(s.InsertBaseHouseType)
    }).
    	Doc("添加Type信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseHouseType{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Type信息").Handle(s.UpdateBaseHouseType)
    }).
    	Doc("修改Type信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseHouseType{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Type信息").Handle(s.DeleteBaseHouseType)
    }).
    	Doc("删除Type信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Type信息").Handle(s.ExportBaseHouseType)
    }).
    	Doc("导出Type信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Type信息").Handle(s.ImportBaseHouseType)
        }).
        Doc("导入Type信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseHouseType{}))

    container.Add(ws)
}

