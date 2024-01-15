package api
// ==========================================================================
// 生成日期：2024-01-12 00:12:41 +0800 CST
// 生成路径: temp/ac/api/ac_event_device_alarm.go
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

type AcEventDeviceAlarmApi struct {
	AcEventDeviceAlarmApp services.AcEventDeviceAlarmModel
}

// GetAcEventDeviceAlarmList Alarm列表数据
func (p *AcEventDeviceAlarmApi) GetAcEventDeviceAlarmList(rc *restfulx.ReqCtx) {
    data := entity.AcEventDeviceAlarm{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.DeviceName = restfulx.QueryParam(rc, "deviceName")
    data.EventType = restfulx.QueryParam(rc, "eventType")
	data.EventTime, _ = time.Parse(time.RFC3339, restfulx.QueryParam(rc, "eventTime"))
	data.EventTime = data.EventTime.In(time.Local)
    data.HandleStatus = restfulx.QueryParam(rc, "handleStatus")
	data.HandleTime, _ = time.Parse(time.RFC3339, restfulx.QueryParam(rc, "handleTime"))
	data.HandleTime = data.HandleTime.In(time.Local)
    data.Notes = restfulx.QueryParam(rc, "notes")

	list, total := p.AcEventDeviceAlarmApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetAcEventDeviceAlarm 获取Alarm
func (p *AcEventDeviceAlarmApi) GetAcEventDeviceAlarm(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.AcEventDeviceAlarmApp.FindOne(int64(id))
}

// InsertAcEventDeviceAlarm 添加Alarm
func (p *AcEventDeviceAlarmApi) InsertAcEventDeviceAlarm(rc *restfulx.ReqCtx) {
	var data entity.AcEventDeviceAlarm
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.AcEventDeviceAlarmApp.Insert(data)
}

// UpdateAcEventDeviceAlarm 修改Alarm
func (p *AcEventDeviceAlarmApi) UpdateAcEventDeviceAlarm(rc *restfulx.ReqCtx) {
	var data entity.AcEventDeviceAlarm
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.AcEventDeviceAlarmApp.Update(data)
}

// DeleteAcEventDeviceAlarm 删除Alarm
func (p *AcEventDeviceAlarmApi) DeleteAcEventDeviceAlarm(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.AcEventDeviceAlarmApp.Delete(ids)
}

// ExportAcEventDeviceAlarm 导出Alarm
func (p *AcEventDeviceAlarmApi) ExportAcEventDeviceAlarm(rc *restfulx.ReqCtx) {
    data := entity.AcEventDeviceAlarm{}

    filename := restfulx.QueryParam(rc, "filename")
    data.DeviceName = restfulx.QueryParam(rc, "deviceName")
    data.EventType = restfulx.QueryParam(rc, "eventType")
    data.EventTime, _ = time.Parse(time.RFC3339, restfulx.QueryParam(rc, "eventTime"))
    data.EventTime = data.EventTime.In(time.Local)
    data.HandleStatus = restfulx.QueryParam(rc, "handleStatus")
    data.HandleTime, _ = time.Parse(time.RFC3339, restfulx.QueryParam(rc, "handleTime"))
    data.HandleTime = data.HandleTime.In(time.Local)
    data.Notes = restfulx.QueryParam(rc, "notes")

    list := p.AcEventDeviceAlarmApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportAcEventDeviceAlarm 导入Alarm
func (p *AcEventDeviceAlarmApi) ImportAcEventDeviceAlarm(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Alarm出现异常"), "导入Alarm出现异常")
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

        // 创建Alarm对象
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
		data := entity.AcEventDeviceAlarm{

		}

		// 保存Alarm对象到数据库中
	   //p.AcEventDeviceAlarmApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
