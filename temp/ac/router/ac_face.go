// ==========================================================================
// 生成日期：2024-01-12 00:02:25 +0800 CST
// 生成路径: temp/ac/router/ac_face.go
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

func InitAcFaceRouter(container *restful.Container) {
	s := &api.AcFaceApi{
		AcFaceApp: services.AcFaceModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/ac/face").Produces(restful.MIME_JSON)
    	tags := []string{"face"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Face分页列表").Handle(s.GetAcFaceList)
    }).
    	Doc("获取Face分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Face信息").Handle(s.GetAcFace)
    }).
    	Doc("获取Face信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.AcFace{}). // on the response
    	Returns(200, "OK", entity.AcFace{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Face信息").Handle(s.InsertAcFace)
    }).
    	Doc("添加Face信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcFace{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Face信息").Handle(s.UpdateAcFace)
    }).
    	Doc("修改Face信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.AcFace{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Face信息").Handle(s.DeleteAcFace)
    }).
    	Doc("删除Face信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Face信息").Handle(s.ExportAcFace)
    }).
    	Doc("导出Face信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Face信息").Handle(s.ImportAcFace)
        }).
        Doc("导入Face信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.AcFace{}))

    container.Add(ws)
}

