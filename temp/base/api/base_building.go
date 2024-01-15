package api
// ==========================================================================
// 生成日期：2024-01-11 23:02:38 +0800 CST
// 生成路径: temp/base/api/base_building.go
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

type BaseBuildingApi struct {
	BaseBuildingApp services.BaseBuildingModel
}

// GetBaseBuildingList Building列表数据
func (p *BaseBuildingApi) GetBaseBuildingList(rc *restfulx.ReqCtx) {
    data := entity.BaseBuilding{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.BuildingName = restfulx.QueryParam(rc, "buildingName")

	list, total := p.BaseBuildingApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseBuilding 获取Building
func (p *BaseBuildingApi) GetBaseBuilding(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseBuildingApp.FindOne(int64(id))
}

// InsertBaseBuilding 添加Building
func (p *BaseBuildingApi) InsertBaseBuilding(rc *restfulx.ReqCtx) {
	var data entity.BaseBuilding
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseBuildingApp.Insert(data)
}

// UpdateBaseBuilding 修改Building
func (p *BaseBuildingApi) UpdateBaseBuilding(rc *restfulx.ReqCtx) {
	var data entity.BaseBuilding
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseBuildingApp.Update(data)
}

// DeleteBaseBuilding 删除Building
func (p *BaseBuildingApi) DeleteBaseBuilding(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseBuildingApp.Delete(ids)
}

// ExportBaseBuilding 导出Building
func (p *BaseBuildingApi) ExportBaseBuilding(rc *restfulx.ReqCtx) {
    data := entity.BaseBuilding{}

    filename := restfulx.QueryParam(rc, "filename")
    data.BuildingName = restfulx.QueryParam(rc, "buildingName")

    list := p.BaseBuildingApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseBuilding 导入Building
func (p *BaseBuildingApi) ImportBaseBuilding(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Building出现异常"), "导入Building出现异常")
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

        // 创建Building对象
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
		data := entity.BaseBuilding{

		}

		// 保存Building对象到数据库中
	   //p.BaseBuildingApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
