package api
// ==========================================================================
// 生成日期：2024-01-11 23:53:55 +0800 CST
// 生成路径: temp/ac/api/base_resident.go
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

type BaseResidentApi struct {
	BaseResidentApp services.BaseResidentModel
}

// GetBaseResidentList Resident列表数据
func (p *BaseResidentApi) GetBaseResidentList(rc *restfulx.ReqCtx) {
    data := entity.BaseResident{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.PersonUuid = restfulx.QueryParam(rc, "personUuid")
    data.HouseUuid = restfulx.QueryParam(rc, "houseUuid")

	list, total := p.BaseResidentApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseResident 获取Resident
func (p *BaseResidentApi) GetBaseResident(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseResidentApp.FindOne(int64(id))
}

// InsertBaseResident 添加Resident
func (p *BaseResidentApi) InsertBaseResident(rc *restfulx.ReqCtx) {
	var data entity.BaseResident
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseResidentApp.Insert(data)
}

// UpdateBaseResident 修改Resident
func (p *BaseResidentApi) UpdateBaseResident(rc *restfulx.ReqCtx) {
	var data entity.BaseResident
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseResidentApp.Update(data)
}

// DeleteBaseResident 删除Resident
func (p *BaseResidentApi) DeleteBaseResident(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseResidentApp.Delete(ids)
}

// ExportBaseResident 导出Resident
func (p *BaseResidentApi) ExportBaseResident(rc *restfulx.ReqCtx) {
    data := entity.BaseResident{}

    filename := restfulx.QueryParam(rc, "filename")
    data.PersonUuid = restfulx.QueryParam(rc, "personUuid")
    data.HouseUuid = restfulx.QueryParam(rc, "houseUuid")

    list := p.BaseResidentApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseResident 导入Resident
func (p *BaseResidentApi) ImportBaseResident(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Resident出现异常"), "导入Resident出现异常")
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

        // 创建Resident对象
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
		data := entity.BaseResident{

		}

		// 保存Resident对象到数据库中
	   //p.BaseResidentApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
