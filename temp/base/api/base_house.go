package api
// ==========================================================================
// 生成日期：2024-01-11 23:09:55 +0800 CST
// 生成路径: temp/base/api/base_house.go
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

type BaseHouseApi struct {
	BaseHouseApp services.BaseHouseModel
}

// GetBaseHouseList House列表数据
func (p *BaseHouseApi) GetBaseHouseList(rc *restfulx.ReqCtx) {
    data := entity.BaseHouse{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.UsageType = restfulx.QueryParam(rc, "usageType")

	list, total := p.BaseHouseApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseHouse 获取House
func (p *BaseHouseApi) GetBaseHouse(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseHouseApp.FindOne(int64(id))
}

// InsertBaseHouse 添加House
func (p *BaseHouseApi) InsertBaseHouse(rc *restfulx.ReqCtx) {
	var data entity.BaseHouse
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseHouseApp.Insert(data)
}

// UpdateBaseHouse 修改House
func (p *BaseHouseApi) UpdateBaseHouse(rc *restfulx.ReqCtx) {
	var data entity.BaseHouse
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseHouseApp.Update(data)
}

// DeleteBaseHouse 删除House
func (p *BaseHouseApi) DeleteBaseHouse(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseHouseApp.Delete(ids)
}

// ExportBaseHouse 导出House
func (p *BaseHouseApi) ExportBaseHouse(rc *restfulx.ReqCtx) {
    data := entity.BaseHouse{}

    filename := restfulx.QueryParam(rc, "filename")
    data.UsageType = restfulx.QueryParam(rc, "usageType")

    list := p.BaseHouseApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseHouse 导入House
func (p *BaseHouseApi) ImportBaseHouse(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入House出现异常"), "导入House出现异常")
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

        // 创建House对象
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
		data := entity.BaseHouse{

		}

		// 保存House对象到数据库中
	   //p.BaseHouseApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
