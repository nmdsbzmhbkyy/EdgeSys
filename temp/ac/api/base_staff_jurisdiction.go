package api
// ==========================================================================
// 生成日期：2024-01-11 23:56:35 +0800 CST
// 生成路径: temp/ac/api/base_staff_jurisdiction.go
// 生成人：aurine
// ==========================================================================
import (
    "EdgeSys/temp/ac/entity"
    "EdgeSys/temp/ac/services"
    "EdgeSys/pkg/global"
    "github.com/xuri/excelize/v2"
    "errors"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/CommonKit/model"
    "mod.miligc.com/edge-common/CommonKit/restfulx"
    "mod.miligc.com/edge-common/CommonKit/utils"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type BaseStaffJurisdictionApi struct {
	BaseStaffJurisdictionApp services.BaseStaffJurisdictionModel
}

// GetBaseStaffJurisdictionList Jurisdiction列表数据
func (p *BaseStaffJurisdictionApi) GetBaseStaffJurisdictionList(rc *restfulx.ReqCtx) {
    data := entity.BaseStaffJurisdiction{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.PersonUuid = restfulx.QueryParam(rc, "personUuid")

	list, total := p.BaseStaffJurisdictionApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseStaffJurisdiction 获取Jurisdiction
func (p *BaseStaffJurisdictionApi) GetBaseStaffJurisdiction(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseStaffJurisdictionApp.FindOne(int64(id))
}

// InsertBaseStaffJurisdiction 添加Jurisdiction
func (p *BaseStaffJurisdictionApi) InsertBaseStaffJurisdiction(rc *restfulx.ReqCtx) {
	var data entity.BaseStaffJurisdiction
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseStaffJurisdictionApp.Insert(data)
}

// UpdateBaseStaffJurisdiction 修改Jurisdiction
func (p *BaseStaffJurisdictionApi) UpdateBaseStaffJurisdiction(rc *restfulx.ReqCtx) {
	var data entity.BaseStaffJurisdiction
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseStaffJurisdictionApp.Update(data)
}

// DeleteBaseStaffJurisdiction 删除Jurisdiction
func (p *BaseStaffJurisdictionApi) DeleteBaseStaffJurisdiction(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseStaffJurisdictionApp.Delete(ids)
}

// ExportBaseStaffJurisdiction 导出Jurisdiction
func (p *BaseStaffJurisdictionApi) ExportBaseStaffJurisdiction(rc *restfulx.ReqCtx) {
    data := entity.BaseStaffJurisdiction{}

    filename := restfulx.QueryParam(rc, "filename")
    data.PersonUuid = restfulx.QueryParam(rc, "personUuid")

    list := p.BaseStaffJurisdictionApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseStaffJurisdiction 导入Jurisdiction
func (p *BaseStaffJurisdictionApi) ImportBaseStaffJurisdiction(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Jurisdiction出现异常"), "导入Jurisdiction出现异常")
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

        // 创建Jurisdiction对象
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
		data := entity.BaseStaffJurisdiction{

		}

		// 保存Jurisdiction对象到数据库中
	   //p.BaseStaffJurisdictionApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
