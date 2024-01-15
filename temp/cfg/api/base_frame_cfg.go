package api
// ==========================================================================
// 生成日期：2024-01-12 00:42:03 +0800 CST
// 生成路径: temp/cfg/api/base_frame_cfg.go
// 生成人：aurine
// ==========================================================================
import (
    "EdgeSys/temp/cfg/entity"
    "EdgeSys/temp/cfg/services"
    "EdgeSys/pkg/global"
    "github.com/xuri/excelize/v2"
    "errors"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/CommonKit/model"
    "mod.miligc.com/edge-common/CommonKit/restfulx"
    "mod.miligc.com/edge-common/CommonKit/utils"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type BaseFrameCfgApi struct {
	BaseFrameCfgApp services.BaseFrameCfgModel
}

// GetBaseFrameCfgList Cfg列表数据
func (p *BaseFrameCfgApi) GetBaseFrameCfgList(rc *restfulx.ReqCtx) {
    data := entity.BaseFrameCfg{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.LevelCode = restfulx.QueryInt(rc, "levelCode", 0)
    data.CodeRule = restfulx.QueryParam(rc, "codeRule")
    data.LevelDesc = restfulx.QueryParam(rc, "levelDesc")
    data.LevelType = restfulx.QueryParam(rc, "levelType")
    data.IsDisable = restfulx.QueryParam(rc, "isDisable")
    data.ProjectUuid = restfulx.QueryParam(rc, "projectUuid")

	list, total := p.BaseFrameCfgApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseFrameCfg 获取Cfg
func (p *BaseFrameCfgApi) GetBaseFrameCfg(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseFrameCfgApp.FindOne(int64(id))
}

// InsertBaseFrameCfg 添加Cfg
func (p *BaseFrameCfgApi) InsertBaseFrameCfg(rc *restfulx.ReqCtx) {
	var data entity.BaseFrameCfg
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseFrameCfgApp.Insert(data)
}

// UpdateBaseFrameCfg 修改Cfg
func (p *BaseFrameCfgApi) UpdateBaseFrameCfg(rc *restfulx.ReqCtx) {
	var data entity.BaseFrameCfg
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseFrameCfgApp.Update(data)
}

// DeleteBaseFrameCfg 删除Cfg
func (p *BaseFrameCfgApi) DeleteBaseFrameCfg(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseFrameCfgApp.Delete(ids)
}

// ExportBaseFrameCfg 导出Cfg
func (p *BaseFrameCfgApi) ExportBaseFrameCfg(rc *restfulx.ReqCtx) {
    data := entity.BaseFrameCfg{}

    filename := restfulx.QueryParam(rc, "filename")
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.LevelCode = restfulx.QueryInt(rc, "levelCode", 0)
    data.CodeRule = restfulx.QueryParam(rc, "codeRule")
    data.LevelDesc = restfulx.QueryParam(rc, "levelDesc")
    data.LevelType = restfulx.QueryParam(rc, "levelType")
    data.IsDisable = restfulx.QueryParam(rc, "isDisable")
    data.ProjectUuid = restfulx.QueryParam(rc, "projectUuid")

    list := p.BaseFrameCfgApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseFrameCfg 导入Cfg
func (p *BaseFrameCfgApi) ImportBaseFrameCfg(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Cfg出现异常"), "导入Cfg出现异常")
	}

	// 解析Excel文件
	xlFile, err := excelize.OpenReader(file)
	if err != nil {
		biz.ErrIsNil(errors.New("解析Excel文件出现异常"), "解析Excel文件出现异常")
	}

	// 遍历Excel文件中的所有行
	rows, _ := xlFile.GetRows("Sheet1")
	for rowIndex, row := range rows {
		// 跳过标题行
		if rowIndex == 0 {
			continue
		}
		pkg.Log.Infoln(row)

        // 创建Cfg对象
        /*
            todo 以导入岗位数据为例 需要完善下方的赋值
            post := entity.SysPost{
            	PostName: row[0],
            	PostCode: row[1],
            	Sort:     sort,
            	Status:   row[3],
            	Remark:   row[4],
            }
        */
		data := entity.BaseFrameCfg{

		}

		// 保存Cfg对象到数据库中
	   //p.BaseFrameCfgApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
