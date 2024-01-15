package api
// ==========================================================================
// 生成日期：2024-01-11 23:18:43 +0800 CST
// 生成路径: temp/base/api/base_staff_department.go
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

type BaseStaffDepartmentApi struct {
	BaseStaffDepartmentApp services.BaseStaffDepartmentModel
}

// GetBaseStaffDepartmentList Department列表数据
func (p *BaseStaffDepartmentApi) GetBaseStaffDepartmentList(rc *restfulx.ReqCtx) {
    data := entity.BaseStaffDepartment{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.DepartmentName = restfulx.QueryParam(rc, "departmentName")
    data.EmployeeCount = restfulx.QueryInt(rc, "employeeCount", 0)
    data.Desc = restfulx.QueryParam(rc, "desc")
    data.ContactPhone = restfulx.QueryParam(rc, "contactPhone")

	list, total := p.BaseStaffDepartmentApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseStaffDepartment 获取Department
func (p *BaseStaffDepartmentApi) GetBaseStaffDepartment(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseStaffDepartmentApp.FindOne(int64(id))
}

// InsertBaseStaffDepartment 添加Department
func (p *BaseStaffDepartmentApi) InsertBaseStaffDepartment(rc *restfulx.ReqCtx) {
	var data entity.BaseStaffDepartment
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseStaffDepartmentApp.Insert(data)
}

// UpdateBaseStaffDepartment 修改Department
func (p *BaseStaffDepartmentApi) UpdateBaseStaffDepartment(rc *restfulx.ReqCtx) {
	var data entity.BaseStaffDepartment
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseStaffDepartmentApp.Update(data)
}

// DeleteBaseStaffDepartment 删除Department
func (p *BaseStaffDepartmentApi) DeleteBaseStaffDepartment(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseStaffDepartmentApp.Delete(ids)
}

// ExportBaseStaffDepartment 导出Department
func (p *BaseStaffDepartmentApi) ExportBaseStaffDepartment(rc *restfulx.ReqCtx) {
    data := entity.BaseStaffDepartment{}

    filename := restfulx.QueryParam(rc, "filename")
    data.DepartmentName = restfulx.QueryParam(rc, "departmentName")
    data.EmployeeCount = restfulx.QueryInt(rc, "employeeCount", 0)
    data.Desc = restfulx.QueryParam(rc, "desc")
    data.ContactPhone = restfulx.QueryParam(rc, "contactPhone")

    list := p.BaseStaffDepartmentApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseStaffDepartment 导入Department
func (p *BaseStaffDepartmentApi) ImportBaseStaffDepartment(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Department出现异常"), "导入Department出现异常")
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

        // 创建Department对象
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
		data := entity.BaseStaffDepartment{

		}

		// 保存Department对象到数据库中
	   //p.BaseStaffDepartmentApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
