package api

import (
	"EdgeSys/apps/log/entity"
	"EdgeSys/apps/log/services"

	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
)

type LogLoginApi struct {
	LogLoginApp services.LogLoginModel
}

func (l *LogLoginApi) GetLoginLogList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	loginLocation := restfulx.QueryParam(rc, "loginLocation")
	username := restfulx.QueryParam(rc, "username")
	list, total := l.LogLoginApp.FindListPage(pageNum, pageSize, entity.LogLogin{LoginLocation: loginLocation, Username: username})
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (l *LogLoginApi) GetLoginLog(rc *restfulx.ReqCtx) {
	infoId := restfulx.PathParamInt(rc, "infoId")
	rc.ResData = l.LogLoginApp.FindOne(int64(infoId))
}

func (l *LogLoginApi) UpdateLoginLog(rc *restfulx.ReqCtx) {
	var log entity.LogLogin
	restfulx.BindQuery(rc, &log)
	l.LogLoginApp.Update(log)
}

func (l *LogLoginApi) DeleteLoginLog(rc *restfulx.ReqCtx) {
	infoIds := restfulx.PathParam(rc, "infoId")
	group := utils.IdsStrToIdsIntGroup(infoIds)
	l.LogLoginApp.Delete(group)
}

func (l *LogLoginApi) DeleteAll(rc *restfulx.ReqCtx) {
	l.LogLoginApp.DeleteAll()
}
