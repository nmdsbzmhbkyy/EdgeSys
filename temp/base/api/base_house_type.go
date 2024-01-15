package api
// ==========================================================================
// 生成日期：2024-01-11 23:01:03 +0800 CST
// 生成路径: temp/base/api/base_house_type.go
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

type BaseHouseTypeApi struct {
	BaseHouseTypeApp services.BaseHouseTypeModel
}

// GetBaseHouseTypeList Type列表数据
func (p *BaseHouseTypeApi) GetBaseHouseTypeList(rc *restfulx.ReqCtx) {
    data := entity.BaseHouseType{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.DesignName = restfulx.QueryParam(rc, "designName")
    data.DesginDesc = restfulx.QueryParam(rc, "desginDesc")

	list, total := p.BaseHouseTypeApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseHouseType 获取Type
func (p *BaseHouseTypeApi) GetBaseHouseType(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseHouseTypeApp.FindOne(int64(id))
}

// InsertBaseHouseType 添加Type
func (p *BaseHouseTypeApi) InsertBaseHouseType(rc *restfulx.ReqCtx) {
	var data entity.BaseHouseType
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseHouseTypeApp.Insert(data)
}

// UpdateBaseHouseType 修改Type
func (p *BaseHouseTypeApi) UpdateBaseHouseType(rc *restfulx.ReqCtx) {
	var data entity.BaseHouseType
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseHouseTypeApp.Update(data)
}

// DeleteBaseHouseType 删除Type
func (p *BaseHouseTypeApi) DeleteBaseHouseType(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseHouseTypeApp.Delete(ids)
}

// ExportBaseHouseType 导出Type
func (p *BaseHouseTypeApi) ExportBaseHouseType(rc *restfulx.ReqCtx) {
    data := entity.BaseHouseType{}

    filename := restfulx.QueryParam(rc, "filename")
    data.DesignName = restfulx.QueryParam(rc, "designName")
    data.DesginDesc = restfulx.QueryParam(rc, "desginDesc")

    list := p.BaseHouseTypeApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseHouseType 导入Type
func (p *BaseHouseTypeApi) ImportBaseHouseType(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Type出现异常"), "导入Type出现异常")
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

        // 创建Type对象
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
		data := entity.BaseHouseType{

		}

		// 保存Type对象到数据库中
	   //p.BaseHouseTypeApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
