package api
// ==========================================================================
// 生成日期：2024-01-12 00:12:57 +0800 CST
// 生成路径: temp/ac/api/ac_event_access.go
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

type AcEventAccessApi struct {
	AcEventAccessApp services.AcEventAccessModel
}

// GetAcEventAccessList Access列表数据
func (p *AcEventAccessApi) GetAcEventAccessList(rc *restfulx.ReqCtx) {
    data := entity.AcEventAccess{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.PersonUuid = restfulx.QueryParam(rc, "personUuid")
    data.EventType = restfulx.QueryParam(rc, "eventType")
	data.EventTime, _ = time.Parse(time.RFC3339, restfulx.QueryParam(rc, "eventTime"))
	data.EventTime = data.EventTime.In(time.Local)

	list, total := p.AcEventAccessApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetAcEventAccess 获取Access
func (p *AcEventAccessApi) GetAcEventAccess(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.AcEventAccessApp.FindOne(int64(id))
}

// InsertAcEventAccess 添加Access
func (p *AcEventAccessApi) InsertAcEventAccess(rc *restfulx.ReqCtx) {
	var data entity.AcEventAccess
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.AcEventAccessApp.Insert(data)
}

// UpdateAcEventAccess 修改Access
func (p *AcEventAccessApi) UpdateAcEventAccess(rc *restfulx.ReqCtx) {
	var data entity.AcEventAccess
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.AcEventAccessApp.Update(data)
}

// DeleteAcEventAccess 删除Access
func (p *AcEventAccessApi) DeleteAcEventAccess(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.AcEventAccessApp.Delete(ids)
}

// ExportAcEventAccess 导出Access
func (p *AcEventAccessApi) ExportAcEventAccess(rc *restfulx.ReqCtx) {
    data := entity.AcEventAccess{}

    filename := restfulx.QueryParam(rc, "filename")
    data.PersonUuid = restfulx.QueryParam(rc, "personUuid")
    data.EventType = restfulx.QueryParam(rc, "eventType")
    data.EventTime, _ = time.Parse(time.RFC3339, restfulx.QueryParam(rc, "eventTime"))
    data.EventTime = data.EventTime.In(time.Local)

    list := p.AcEventAccessApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportAcEventAccess 导入Access
func (p *AcEventAccessApi) ImportAcEventAccess(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Access出现异常"), "导入Access出现异常")
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

        // 创建Access对象
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
		data := entity.AcEventAccess{

		}

		// 保存Access对象到数据库中
	   //p.AcEventAccessApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
