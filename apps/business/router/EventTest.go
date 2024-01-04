package router

import (
	"EdgeSys/apps/business"
	"EdgeSys/apps/business/entity"
	"encoding/json"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"mod.miligc.com/edge-common/CommonKit/restfulx"
)

func InitEventRouter(container *restful.Container) {

	ws := new(restful.WebService)
	ws.Path("/event/api").Produces(restful.MIME_JSON)
	tags := []string{"api"}

	ws.Route(ws.POST("/test").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("测试事件上报").Handle(Event)
	}).
		Doc("测试事件上报").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.Event{}).
		Returns(200, "OK", []entity.Event{}))

	container.Add(ws)
}

func Event(rc *restfulx.ReqCtx) {
	var api entity.Event
	restfulx.BindJsonAndValid(rc, &api)
	data, _ := json.Marshal(api)
	_ = business.PersonComponentInterface.EventHandle(string(data))

}
