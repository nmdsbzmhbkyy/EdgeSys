package api
// ==========================================================================
// 生成日期：2024-01-11 23:06:12 +0800 CST
// 生成路径: temp/base/api/base_unit.go
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

type BaseUnitApi struct {
	BaseUnitApp services.BaseUnitModel
}

// GetBaseUnitList Unit列表数据
func (p *BaseUnitApi) GetBaseUnitList(rc *restfulx.ReqCtx) {
    data := entity.BaseUnit{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.UnitCode = restfulx.QueryParam(rc, "unitCode")
    data.FrameCode = restfulx.QueryParam(rc, "frameCode")
    data.UnitName = restfulx.QueryParam(rc, "unitName")
    data.AddressCode = restfulx.QueryParam(rc, "addressCode")
    data.PicUrl1 = restfulx.QueryParam(rc, "picUrl1")
    data.PicUrl2 = restfulx.QueryParam(rc, "picUrl2")
    data.PicUrl3 = restfulx.QueryParam(rc, "picUrl3")
    data.ProjectUuid = restfulx.QueryParam(rc, "projectUuid")

	list, total := p.BaseUnitApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseUnit 获取Unit
func (p *BaseUnitApi) GetBaseUnit(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseUnitApp.FindOne(int64(id))
}

// InsertBaseUnit 添加Unit
func (p *BaseUnitApi) InsertBaseUnit(rc *restfulx.ReqCtx) {
	var data entity.BaseUnit
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseUnitApp.Insert(data)
}

// UpdateBaseUnit 修改Unit
func (p *BaseUnitApi) UpdateBaseUnit(rc *restfulx.ReqCtx) {
	var data entity.BaseUnit
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseUnitApp.Update(data)
}

// DeleteBaseUnit 删除Unit
func (p *BaseUnitApi) DeleteBaseUnit(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseUnitApp.Delete(ids)
}

// ExportBaseUnit 导出Unit
func (p *BaseUnitApi) ExportBaseUnit(rc *restfulx.ReqCtx) {
    data := entity.BaseUnit{}

    filename := restfulx.QueryParam(rc, "filename")
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.UnitCode = restfulx.QueryParam(rc, "unitCode")
    data.FrameCode = restfulx.QueryParam(rc, "frameCode")
    data.UnitName = restfulx.QueryParam(rc, "unitName")
    data.AddressCode = restfulx.QueryParam(rc, "addressCode")
    data.PicUrl1 = restfulx.QueryParam(rc, "picUrl1")
    data.PicUrl2 = restfulx.QueryParam(rc, "picUrl2")
    data.PicUrl3 = restfulx.QueryParam(rc, "picUrl3")
    data.ProjectUuid = restfulx.QueryParam(rc, "projectUuid")

    list := p.BaseUnitApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseUnit 导入Unit
func (p *BaseUnitApi) ImportBaseUnit(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Unit出现异常"), "导入Unit出现异常")
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

        // 创建Unit对象
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
		data := entity.BaseUnit{

		}

		// 保存Unit对象到数据库中
	   //p.BaseUnitApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
