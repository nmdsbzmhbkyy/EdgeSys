package api
// ==========================================================================
// 生成日期：2024-01-11 23:25:46 +0800 CST
// 生成路径: temp/base/api/base_staff_position.go
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

type BaseStaffPositionApi struct {
	BaseStaffPositionApp services.BaseStaffPositionModel
}

// GetBaseStaffPositionList Position列表数据
func (p *BaseStaffPositionApi) GetBaseStaffPositionList(rc *restfulx.ReqCtx) {
    data := entity.BaseStaffPosition{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.PositionName = restfulx.QueryParam(rc, "positionName")
    data.PositionType = restfulx.QueryParam(rc, "positionType")
    data.Desc = restfulx.QueryParam(rc, "desc")

	list, total := p.BaseStaffPositionApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseStaffPosition 获取Position
func (p *BaseStaffPositionApi) GetBaseStaffPosition(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseStaffPositionApp.FindOne(int64(id))
}

// InsertBaseStaffPosition 添加Position
func (p *BaseStaffPositionApi) InsertBaseStaffPosition(rc *restfulx.ReqCtx) {
	var data entity.BaseStaffPosition
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseStaffPositionApp.Insert(data)
}

// UpdateBaseStaffPosition 修改Position
func (p *BaseStaffPositionApi) UpdateBaseStaffPosition(rc *restfulx.ReqCtx) {
	var data entity.BaseStaffPosition
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseStaffPositionApp.Update(data)
}

// DeleteBaseStaffPosition 删除Position
func (p *BaseStaffPositionApi) DeleteBaseStaffPosition(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseStaffPositionApp.Delete(ids)
}

// ExportBaseStaffPosition 导出Position
func (p *BaseStaffPositionApi) ExportBaseStaffPosition(rc *restfulx.ReqCtx) {
    data := entity.BaseStaffPosition{}

    filename := restfulx.QueryParam(rc, "filename")
    data.PositionName = restfulx.QueryParam(rc, "positionName")
    data.PositionType = restfulx.QueryParam(rc, "positionType")
    data.Desc = restfulx.QueryParam(rc, "desc")

    list := p.BaseStaffPositionApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseStaffPosition 导入Position
func (p *BaseStaffPositionApi) ImportBaseStaffPosition(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Position出现异常"), "导入Position出现异常")
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

        // 创建Position对象
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
		data := entity.BaseStaffPosition{

		}

		// 保存Position对象到数据库中
	   //p.BaseStaffPositionApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
