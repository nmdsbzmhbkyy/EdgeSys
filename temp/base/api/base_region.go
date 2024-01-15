package api
// ==========================================================================
// 生成日期：2024-01-11 23:15:00 +0800 CST
// 生成路径: temp/base/api/base_region.go
// 生成人：aurine
// ==========================================================================
import (
    "EdgeSys/temp/base/entity"
    "EdgeSys/temp/base/services"
    "EdgeSys/pkg/global"
    "github.com/xuri/excelize/v2"
    "errors"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/CommonKit/model"
    "mod.miligc.com/edge-common/CommonKit/restfulx"
    "mod.miligc.com/edge-common/CommonKit/utils"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type BaseRegionApi struct {
	BaseRegionApp services.BaseRegionModel
}

// GetBaseRegionList Region列表数据
func (p *BaseRegionApi) GetBaseRegionList(rc *restfulx.ReqCtx) {
    data := entity.BaseRegion{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.Pid = int64(restfulx.QueryInt(rc, "pid", 0))
    data.Name = restfulx.QueryParam(rc, "name")

	list, total := p.BaseRegionApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseRegion 获取Region
func (p *BaseRegionApi) GetBaseRegion(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseRegionApp.FindOne(int64(id))
}

// InsertBaseRegion 添加Region
func (p *BaseRegionApi) InsertBaseRegion(rc *restfulx.ReqCtx) {
	var data entity.BaseRegion
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseRegionApp.Insert(data)
}

// UpdateBaseRegion 修改Region
func (p *BaseRegionApi) UpdateBaseRegion(rc *restfulx.ReqCtx) {
	var data entity.BaseRegion
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseRegionApp.Update(data)
}

// DeleteBaseRegion 删除Region
func (p *BaseRegionApi) DeleteBaseRegion(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseRegionApp.Delete(ids)
}

// ExportBaseRegion 导出Region
func (p *BaseRegionApi) ExportBaseRegion(rc *restfulx.ReqCtx) {
    data := entity.BaseRegion{}

    filename := restfulx.QueryParam(rc, "filename")
    data.Pid = int64(restfulx.QueryInt(rc, "pid", 0))
    data.Name = restfulx.QueryParam(rc, "name")

    list := p.BaseRegionApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseRegion 导入Region
func (p *BaseRegionApi) ImportBaseRegion(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Region出现异常"), "导入Region出现异常")
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

        // 创建Region对象
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
		data := entity.BaseRegion{

		}

		// 保存Region对象到数据库中
	   //p.BaseRegionApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
