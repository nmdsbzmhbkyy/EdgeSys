// ==========================================================================
// 生成日期：2024-01-11 23:13:08 +0800 CST
// 生成路径: temp/base/router/base_frame.go
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

func InitBaseFrameRouter(container *restful.Container) {
	s := &api.BaseFrameApi{
		BaseFrameApp: services.BaseFrameModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/base/frame").Produces(restful.MIME_JSON)
    	tags := []string{"frame"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Frame分页列表").Handle(s.GetBaseFrameList)
    }).
    	Doc("获取Frame分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Frame信息").Handle(s.GetBaseFrame)
    }).
    	Doc("获取Frame信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseFrame{}). // on the response
    	Returns(200, "OK", entity.BaseFrame{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Frame信息").Handle(s.InsertBaseFrame)
    }).
    	Doc("添加Frame信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseFrame{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Frame信息").Handle(s.UpdateBaseFrame)
    }).
    	Doc("修改Frame信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseFrame{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Frame信息").Handle(s.DeleteBaseFrame)
    }).
    	Doc("删除Frame信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Frame信息").Handle(s.ExportBaseFrame)
    }).
    	Doc("导出Frame信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Frame信息").Handle(s.ImportBaseFrame)
        }).
        Doc("导入Frame信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseFrame{}))

    container.Add(ws)
}

