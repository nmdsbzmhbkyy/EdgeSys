// ==========================================================================
// 生成日期：2024-01-11 23:18:43 +0800 CST
// 生成路径: temp/base/router/base_staff_department.go
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

func InitBaseStaffDepartmentRouter(container *restful.Container) {
	s := &api.BaseStaffDepartmentApi{
		BaseStaffDepartmentApp: services.BaseStaffDepartmentModelDao,
	}

	ws := new(restful.WebService)
    	ws.Path("/base/staffDepartment").Produces(restful.MIME_JSON)
    	tags := []string{"staffDepartment"}


    ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
    		restfulx.NewReqCtx(request, response).WithLog("获取Department分页列表").Handle(s.GetBaseStaffDepartmentList)
    }).
    	Doc("获取Department分页列表").
    	Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
        Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(model.ResultPage{}).
    	Returns(200, "OK", model.ResultPage{}))

    ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("获取Department信息").Handle(s.GetBaseStaffDepartment)
    }).
    	Doc("获取Department信息").
    	Param(ws.PathParameter("id", "Id").DataType("int64")).
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Writes(entity.BaseStaffDepartment{}). // on the response
    	Returns(200, "OK", entity.BaseStaffDepartment{}).
    	Returns(404, "Not Found", nil))

    ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("添加Department信息").Handle(s.InsertBaseStaffDepartment)
    }).
    	Doc("添加Department信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseStaffDepartment{}))

    ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("修改Department信息").Handle(s.UpdateBaseStaffDepartment)
    }).
    	Doc("修改Department信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Reads(entity.BaseStaffDepartment{}))

    ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("删除Department信息").Handle(s.DeleteBaseStaffDepartment)
    }).
    	Doc("删除Department信息").
    	Metadata(restfulspec.KeyOpenAPITags, tags).
    	Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))


    ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
    	restfulx.NewReqCtx(request, response).WithLog("导出Department信息").Handle(s.ExportBaseStaffDepartment)
    }).
    	Doc("导出Department信息").
    	Param(ws.QueryParameter("filename", "filename").DataType("string")).
    	Metadata(restfulspec.KeyOpenAPITags, tags))

    ws.Route(ws.POST("/import").To(func(request *restful.Request, response *restful.Response) {
        restfulx.NewReqCtx(request, response).WithLog("导入Department信息").Handle(s.ImportBaseStaffDepartment)
        }).
        Doc("导入Department信息").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        Reads(entity.BaseStaffDepartment{}))

    container.Add(ws)
}

