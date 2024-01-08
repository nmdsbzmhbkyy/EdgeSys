package api

import (
	entity "EdgeSys/apps/system/entity"
	services "EdgeSys/apps/system/services"
	"EdgeSys/pkg/global"
	"fmt"
	public "mod.miligc.com/edge-common/edgesys-common/pkg"

	"mod.miligc.com/edge-common/CommonKit/biz"
	"mod.miligc.com/edge-common/CommonKit/model"
	"mod.miligc.com/edge-common/CommonKit/restfulx"
	"mod.miligc.com/edge-common/CommonKit/utils"
)

type DictApi struct {
	DictType services.SysDictTypeModel
	DictData services.SysDictDataModel
}

func (p *DictApi) GetDictTypeList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	status := restfulx.QueryParam(rc, "status")
	dictName := restfulx.QueryParam(rc, "dictName")
	dictType := restfulx.QueryParam(rc, "dictType")

	list, total := p.DictType.FindListPage(pageNum, pageSize, entity.SysDictType{Status: status, DictName: dictName, DictType: dictType})
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (p *DictApi) GetDictType(rc *restfulx.ReqCtx) {
	dictId := restfulx.PathParamInt(rc, "dictId")
	p.DictType.FindOne(int64(dictId))
}

func (p *DictApi) InsertDictType(rc *restfulx.ReqCtx) {
	var dict entity.SysDictType
	restfulx.BindJsonAndValid(rc, &dict)

	dict.CreateBy = rc.LoginAccount.UserName
	p.DictType.Insert(dict)
}

func (p *DictApi) UpdateDictType(rc *restfulx.ReqCtx) {
	var dict entity.SysDictType
	restfulx.BindJsonAndValid(rc, &dict)

	dict.CreateBy = rc.LoginAccount.UserName
	p.DictType.Update(dict)
}

func (p *DictApi) DeleteDictType(rc *restfulx.ReqCtx) {
	dictId := restfulx.PathParam(rc, "dictId")
	dictIds := utils.IdsStrToIdsIntGroup(dictId)

	deList := make([]int64, 0)
	for _, id := range dictIds {
		one := p.DictType.FindOne(id)
		list := p.DictData.FindList(entity.SysDictData{DictType: one.DictType})
		if len(*list) == 0 {
			deList = append(deList, id)
		} else {
			public.Log.Info(fmt.Sprintf("dictId: %d 存在字典数据绑定无法删除", id))
		}
	}
	p.DictType.Delete(deList)
}

func (p *DictApi) ExportDictType(rc *restfulx.ReqCtx) {
	filename := restfulx.QueryParam(rc, "filename")
	status := restfulx.QueryParam(rc, "status")
	dictName := restfulx.QueryParam(rc, "dictName")
	dictType := restfulx.QueryParam(rc, "dictType")

	list := p.DictType.FindList(entity.SysDictType{Status: status, DictName: dictName, DictType: dictType})
	fileName := utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, fileName)
	rc.Download(fileName)
}

func (p *DictApi) GetDictDataList(rc *restfulx.ReqCtx) {
	dictLabel := restfulx.QueryParam(rc, "dictLabel")
	dictType := restfulx.QueryParam(rc, "dictType")
	status := restfulx.QueryParam(rc, "status")
	rc.ResData = p.DictData.FindList(entity.SysDictData{Status: status, DictType: dictType, DictLabel: dictLabel})
}

func (p *DictApi) GetDictDataListByDictType(rc *restfulx.ReqCtx) {
	dictType := restfulx.QueryParam(rc, "dictType")
	biz.IsTrue(dictType != "", "请传入字典类型")
	rc.ResData = p.DictData.FindList(entity.SysDictData{DictType: dictType})
}

func (p *DictApi) GetDictData(rc *restfulx.ReqCtx) {
	dictCode := restfulx.PathParamInt(rc, "dictCode")
	p.DictData.FindOne(int64(dictCode))
}

func (p *DictApi) InsertDictData(rc *restfulx.ReqCtx) {
	var data entity.SysDictData
	restfulx.BindJsonAndValid(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName
	p.DictData.Insert(data)
}

func (p *DictApi) UpdateDictData(rc *restfulx.ReqCtx) {
	var data entity.SysDictData
	restfulx.BindJsonAndValid(rc, &data)

	data.CreateBy = rc.LoginAccount.UserName
	p.DictData.Update(data)
}

func (p *DictApi) DeleteDictData(rc *restfulx.ReqCtx) {
	dictCode := restfulx.PathParam(rc, "dictCode")
	p.DictData.Delete(utils.IdsStrToIdsIntGroup(dictCode))
}
