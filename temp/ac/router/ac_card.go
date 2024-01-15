// ==========================================================================
// 生成日期：2024-01-12 00:05:03 +0800 CST
// 生成路径: temp/ac/router/ac_card.go
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

func InitAcCardRouter(container *restful.Container) {
	s := &api.AcCardApi{
		AcCardApp: services.AcCardModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/ac/card").Produces(restful.MIME_JSON)
    	tags := []string{"card"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Card分页列表").Handle(s.GetAcCardList)
    }).
    	Doc("获取Card分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Card信息").Handle(s.GetAcCard)
    }).
    	Doc("获取Card信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.AcCard{}). // on the response
    	Returns(200, "OK", entity.AcCard{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Card信息").Handle(s.InsertAcCard)
    }).
    	Doc("添加Card信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcCard{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Card信息").Handle(s.UpdateAcCard)
    }).
    	Doc("修改Card信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcCard{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Card信息").Handle(s.DeleteAcCard)
    }).
    	Doc("删除Card信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Card信息").Handle(s.ExportAcCard)
    }).
    	Doc("导出Card信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Card信息").Handle(s.ImportAcCard)
        }).
        Doc("导入Card信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.AcCard{}))

    container.Add(ws)
}

