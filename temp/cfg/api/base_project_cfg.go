package api
// ==========================================================================
// 生成日期：2024-01-12 00:42:26 +0800 CST
// 生成路径: temp/cfg/api/base_project_cfg.go
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

type BaseProjectCfgApi struct {
	BaseProjectCfgApp services.BaseProjectCfgModel
}

// GetBaseProjectCfgList Cfg列表数据
func (p *BaseProjectCfgApi) GetBaseProjectCfgList(rc *restfulx.ReqCtx) {
    data := entity.BaseProjectCfg{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)

	list, total := p.BaseProjectCfgApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseProjectCfg 获取Cfg
func (p *BaseProjectCfgApi) GetBaseProjectCfg(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseProjectCfgApp.FindOne(int64(id))
}

// InsertBaseProjectCfg 添加Cfg
func (p *BaseProjectCfgApi) InsertBaseProjectCfg(rc *restfulx.ReqCtx) {
	var data entity.BaseProjectCfg
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseProjectCfgApp.Insert(data)
}

// UpdateBaseProjectCfg 修改Cfg
func (p *BaseProjectCfgApi) UpdateBaseProjectCfg(rc *restfulx.ReqCtx) {
	var data entity.BaseProjectCfg
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseProjectCfgApp.Update(data)
}

// DeleteBaseProjectCfg 删除Cfg
func (p *BaseProjectCfgApi) DeleteBaseProjectCfg(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseProjectCfgApp.Delete(ids)
}

// ExportBaseProjectCfg 导出Cfg
func (p *BaseProjectCfgApi) ExportBaseProjectCfg(rc *restfulx.ReqCtx) {
    data := entity.BaseProjectCfg{}

    filename := restfulx.QueryParam(rc, "filename")

    list := p.BaseProjectCfgApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseProjectCfg 导入Cfg
func (p *BaseProjectCfgApi) ImportBaseProjectCfg(rc *restfulx.ReqCtx) {
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
		data := entity.BaseProjectCfg{

		}

		// 保存Cfg对象到数据库中
	   //p.BaseProjectCfgApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
