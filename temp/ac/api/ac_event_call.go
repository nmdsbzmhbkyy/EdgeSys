package api
// ==========================================================================
// 生成日期：2024-01-12 00:13:10 +0800 CST
// 生成路径: temp/ac/api/ac_event_call.go
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
    "time"
)

type AcEventCallApi struct {
	AcEventCallApp services.AcEventCallModel
}

// GetAcEventCallList Call列表数据
func (p *AcEventCallApi) GetAcEventCallList(rc *restfulx.ReqCtx) {
    data := entity.AcEventCall{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.CallType = restfulx.QueryParam(rc, "callType")
    data.CallResult = restfulx.QueryParam(rc, "callResult")
	data.StartTime, _ = time.Parse(time.RFC3339, restfulx.QueryParam(rc, "startTime"))
	data.StartTime = data.StartTime.In(time.Local)
	data.EndTime, _ = time.Parse(time.RFC3339, restfulx.QueryParam(rc, "endTime"))
	data.EndTime = data.EndTime.In(time.Local)
    data.CallDuration = int64(restfulx.QueryInt(rc, "callDuration", 0))

	list, total := p.AcEventCallApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetAcEventCall 获取Call
func (p *AcEventCallApi) GetAcEventCall(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.AcEventCallApp.FindOne(int64(id))
}

// InsertAcEventCall 添加Call
func (p *AcEventCallApi) InsertAcEventCall(rc *restfulx.ReqCtx) {
	var data entity.AcEventCall
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.AcEventCallApp.Insert(data)
}

// UpdateAcEventCall 修改Call
func (p *AcEventCallApi) UpdateAcEventCall(rc *restfulx.ReqCtx) {
	var data entity.AcEventCall
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.AcEventCallApp.Update(data)
}

// DeleteAcEventCall 删除Call
func (p *AcEventCallApi) DeleteAcEventCall(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.AcEventCallApp.Delete(ids)
}

// ExportAcEventCall 导出Call
func (p *AcEventCallApi) ExportAcEventCall(rc *restfulx.ReqCtx) {
    data := entity.AcEventCall{}

    filename := restfulx.QueryParam(rc, "filename")
    data.CallType = restfulx.QueryParam(rc, "callType")
    data.CallResult = restfulx.QueryParam(rc, "callResult")
    data.StartTime, _ = time.Parse(time.RFC3339, restfulx.QueryParam(rc, "startTime"))
    data.StartTime = data.StartTime.In(time.Local)
    data.EndTime, _ = time.Parse(time.RFC3339, restfulx.QueryParam(rc, "endTime"))
    data.EndTime = data.EndTime.In(time.Local)
    data.CallDuration = int64(restfulx.QueryInt(rc, "callDuration", 0))

    list := p.AcEventCallApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportAcEventCall 导入Call
func (p *AcEventCallApi) ImportAcEventCall(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Call出现异常"), "导入Call出现异常")
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

        // 创建Call对象
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
		data := entity.AcEventCall{

		}

		// 保存Call对象到数据库中
	   //p.AcEventCallApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
