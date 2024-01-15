package api
// ==========================================================================
// 生成日期：2024-01-12 00:42:15 +0800 CST
// 生成路径: temp/cfg/api/base_frame.go
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

type BaseFrameApi struct {
	BaseFrameApp services.BaseFrameModel
}

// GetBaseFrameList Frame列表数据
func (p *BaseFrameApi) GetBaseFrameList(rc *restfulx.ReqCtx) {
    data := entity.BaseFrame{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)

	list, total := p.BaseFrameApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseFrame 获取Frame
func (p *BaseFrameApi) GetBaseFrame(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseFrameApp.FindOne(int64(id))
}

// InsertBaseFrame 添加Frame
func (p *BaseFrameApi) InsertBaseFrame(rc *restfulx.ReqCtx) {
	var data entity.BaseFrame
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseFrameApp.Insert(data)
}

// UpdateBaseFrame 修改Frame
func (p *BaseFrameApi) UpdateBaseFrame(rc *restfulx.ReqCtx) {
	var data entity.BaseFrame
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseFrameApp.Update(data)
}

// DeleteBaseFrame 删除Frame
func (p *BaseFrameApi) DeleteBaseFrame(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseFrameApp.Delete(ids)
}

// ExportBaseFrame 导出Frame
func (p *BaseFrameApi) ExportBaseFrame(rc *restfulx.ReqCtx) {
    data := entity.BaseFrame{}

    filename := restfulx.QueryParam(rc, "filename")

    list := p.BaseFrameApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseFrame 导入Frame
func (p *BaseFrameApi) ImportBaseFrame(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Frame出现异常"), "导入Frame出现异常")
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

        // 创建Frame对象
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
		data := entity.BaseFrame{

		}

		// 保存Frame对象到数据库中
	   //p.BaseFrameApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
