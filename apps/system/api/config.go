package api

import (
	entity "EdgeSys/apps/system/entity"
	services "EdgeSys/apps/system/services"

	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
)

type ConfigApi struct {
	ConfigApp services.SysConfigModel
}

func (p *ConfigApi) GetConfigList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	configName := rc.Request.QueryParameter("configName")
	configKey := rc.Request.QueryParameter("configKey")
	configType := rc.Request.QueryParameter("configType")
	config := entity.SysConfig{ConfigName: configName, ConfigKey: configKey, ConfigType: configType}
	list, total := p.ConfigApp.FindListPage(pageNum, pageSize, config)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (p *ConfigApi) GetConfigListByKey(rc *restfulx.ReqCtx) {
	configKey := rc.Request.QueryParameter("configKey")
	biz.IsTrue(configKey != "", "请传入配置Key")
	rc.ResData = p.ConfigApp.FindList(entity.SysConfig{ConfigKey: configKey})
}

func (p *ConfigApi) GetConfig(rc *restfulx.ReqCtx) {
	id := restfulx.PathParamInt(rc, "configId")
	p.ConfigApp.FindOne(int64(id))
}

func (p *ConfigApi) InsertConfig(rc *restfulx.ReqCtx) {
	var config entity.SysConfig
	restfulx.BindJsonAndValid(rc, &config)

	p.ConfigApp.Insert(config)
}

func (p *ConfigApi) UpdateConfig(rc *restfulx.ReqCtx) {
	var post entity.SysConfig
	restfulx.BindJsonAndValid(rc, &post)
	p.ConfigApp.Update(post)
}

func (p *ConfigApi) DeleteConfig(rc *restfulx.ReqCtx) {
	configId := rc.Request.PathParameter("configId")
	p.ConfigApp.Delete(utils.IdsStrToIdsIntGroup(configId))
}
