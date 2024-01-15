package api
// ==========================================================================
// 生成日期：2024-01-12 00:41:39 +0800 CST
// 生成路径: temp/cfg/api/base_project.go
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

type BaseProjectApi struct {
	BaseProjectApp services.BaseProjectModel
}

// GetBaseProjectList Project列表数据
func (p *BaseProjectApi) GetBaseProjectList(rc *restfulx.ReqCtx) {
    data := entity.BaseProject{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)

	list, total := p.BaseProjectApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseProject 获取Project
func (p *BaseProjectApi) GetBaseProject(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseProjectApp.FindOne(int64(id))
}

// InsertBaseProject 添加Project
func (p *BaseProjectApi) InsertBaseProject(rc *restfulx.ReqCtx) {
	var data entity.BaseProject
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseProjectApp.Insert(data)
}

// UpdateBaseProject 修改Project
func (p *BaseProjectApi) UpdateBaseProject(rc *restfulx.ReqCtx) {
	var data entity.BaseProject
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseProjectApp.Update(data)
}

// DeleteBaseProject 删除Project
func (p *BaseProjectApi) DeleteBaseProject(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseProjectApp.Delete(ids)
}

// ExportBaseProject 导出Project
func (p *BaseProjectApi) ExportBaseProject(rc *restfulx.ReqCtx) {
    data := entity.BaseProject{}

    filename := restfulx.QueryParam(rc, "filename")

    list := p.BaseProjectApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseProject 导入Project
func (p *BaseProjectApi) ImportBaseProject(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Project出现异常"), "导入Project出现异常")
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

        // 创建Project对象
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
		data := entity.BaseProject{

		}

		// 保存Project对象到数据库中
	   //p.BaseProjectApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
