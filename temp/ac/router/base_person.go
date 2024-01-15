// ==========================================================================
// 生成日期：2024-01-11 23:57:43 +0800 CST
// 生成路径: temp/ac/router/base_person.go
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

func InitBasePersonRouter(container *restful.Container) {
	s := &api.BasePersonApi{
		BasePersonApp: services.BasePersonModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/ac/person").Produces(restful.MIME_JSON)
    	tags := []string{"person"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Person分页列表").Handle(s.GetBasePersonList)
    }).
    	Doc("获取Person分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Person信息").Handle(s.GetBasePerson)
    }).
    	Doc("获取Person信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BasePerson{}). // on the response
    	Returns(200, "OK", entity.BasePerson{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Person信息").Handle(s.InsertBasePerson)
    }).
    	Doc("添加Person信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BasePerson{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Person信息").Handle(s.UpdateBasePerson)
    }).
    	Doc("修改Person信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BasePerson{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Person信息").Handle(s.DeleteBasePerson)
    }).
    	Doc("删除Person信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Person信息").Handle(s.ExportBasePerson)
    }).
    	Doc("导出Person信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Person信息").Handle(s.ImportBasePerson)
        }).
        Doc("导入Person信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BasePerson{}))

    container.Add(ws)
}

