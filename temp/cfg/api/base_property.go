package api
// ==========================================================================
// 生成日期：2024-01-12 00:41:50 +0800 CST
// 生成路径: temp/cfg/api/base_property.go
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

type BasePropertyApi struct {
	BasePropertyApp services.BasePropertyModel
}

// GetBasePropertyList Property列表数据
func (p *BasePropertyApi) GetBasePropertyList(rc *restfulx.ReqCtx) {
    data := entity.BaseProperty{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.PropertyCompany = restfulx.QueryParam(rc, "propertyCompany")
    data.PropertyPrincipal = restfulx.QueryParam(rc, "propertyPrincipal")
    data.PropertyCode = restfulx.QueryParam(rc, "propertyCode")
    data.PropertyPhone = restfulx.QueryParam(rc, "propertyPhone")
    data.PropertyAddress = restfulx.QueryParam(rc, "propertyAddress")
    data.ProjectUuid = restfulx.QueryParam(rc, "projectUuid")

	list, total := p.BasePropertyApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseProperty 获取Property
func (p *BasePropertyApi) GetBaseProperty(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BasePropertyApp.FindOne(int64(id))
}

// InsertBaseProperty 添加Property
func (p *BasePropertyApi) InsertBaseProperty(rc *restfulx.ReqCtx) {
	var data entity.BaseProperty
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BasePropertyApp.Insert(data)
}

// UpdateBaseProperty 修改Property
func (p *BasePropertyApi) UpdateBaseProperty(rc *restfulx.ReqCtx) {
	var data entity.BaseProperty
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BasePropertyApp.Update(data)
}

// DeleteBaseProperty 删除Property
func (p *BasePropertyApi) DeleteBaseProperty(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BasePropertyApp.Delete(ids)
}

// ExportBaseProperty 导出Property
func (p *BasePropertyApi) ExportBaseProperty(rc *restfulx.ReqCtx) {
    data := entity.BaseProperty{}

    filename := restfulx.QueryParam(rc, "filename")
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.PropertyCompany = restfulx.QueryParam(rc, "propertyCompany")
    data.PropertyPrincipal = restfulx.QueryParam(rc, "propertyPrincipal")
    data.PropertyCode = restfulx.QueryParam(rc, "propertyCode")
    data.PropertyPhone = restfulx.QueryParam(rc, "propertyPhone")
    data.PropertyAddress = restfulx.QueryParam(rc, "propertyAddress")
    data.ProjectUuid = restfulx.QueryParam(rc, "projectUuid")

    list := p.BasePropertyApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseProperty 导入Property
func (p *BasePropertyApi) ImportBaseProperty(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Property出现异常"), "导入Property出现异常")
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

        // 创建Property对象
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
		data := entity.BaseProperty{

		}

		// 保存Property对象到数据库中
	   //p.BasePropertyApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
