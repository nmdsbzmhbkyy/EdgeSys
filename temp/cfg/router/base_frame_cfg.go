// ==========================================================================
// 生成日期：2024-01-12 00:42:03 +0800 CST
// 生成路径: temp/cfg/router/base_frame_cfg.go
// 生成人：aurine
// ==========================================================================
package router

import (
    "mod.miligc.com/edge-common/CommonKit/model"
    "mod.miligc.com/edge-common/CommonKit/restfulx"
	"EdgeSys/temp/cfg/api"
	"EdgeSys/temp/cfg/entity"
	"EdgeSys/temp/cfg/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
    "github.com/emicklei/go-restful/v3"
)

func InitBaseFrameCfgRouter(container *restful.Container) {
	s := &api.BaseFrameCfgApi{
		BaseFrameCfgApp: services.BaseFrameCfgModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/cfg/framecfg").Produces(restful.MIME_JSON)
    	tags := []string{"framecfg"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Cfg分页列表").Handle(s.GetBaseFrameCfgList)
    }).
    	Doc("获取Cfg分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Cfg信息").Handle(s.GetBaseFrameCfg)
    }).
    	Doc("获取Cfg信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseFrameCfg{}). // on the response
    	Returns(200, "OK", entity.BaseFrameCfg{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Cfg信息").Handle(s.InsertBaseFrameCfg)
    }).
    	Doc("添加Cfg信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseFrameCfg{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Cfg信息").Handle(s.UpdateBaseFrameCfg)
    }).
    	Doc("修改Cfg信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseFrameCfg{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Cfg信息").Handle(s.DeleteBaseFrameCfg)
    }).
    	Doc("删除Cfg信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Cfg信息").Handle(s.ExportBaseFrameCfg)
    }).
    	Doc("导出Cfg信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Cfg信息").Handle(s.ImportBaseFrameCfg)
        }).
        Doc("导入Cfg信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseFrameCfg{}))

    container.Add(ws)
}

