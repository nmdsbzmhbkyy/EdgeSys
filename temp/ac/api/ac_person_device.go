package api
// ==========================================================================
// 生成日期：2024-01-11 23:59:35 +0800 CST
// 生成路径: temp/ac/api/ac_person_device.go
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

type AcPersonDeviceApi struct {
	AcPersonDeviceApp services.AcPersonDeviceModel
}

// GetAcPersonDeviceList Device列表数据
func (p *AcPersonDeviceApi) GetAcPersonDeviceList(rc *restfulx.ReqCtx) {
    data := entity.AcPersonDevice{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.PersonUuid = restfulx.QueryParam(rc, "personUuid")

	list, total := p.AcPersonDeviceApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetAcPersonDevice 获取Device
func (p *AcPersonDeviceApi) GetAcPersonDevice(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.AcPersonDeviceApp.FindOne(int64(id))
}

// InsertAcPersonDevice 添加Device
func (p *AcPersonDeviceApi) InsertAcPersonDevice(rc *restfulx.ReqCtx) {
	var data entity.AcPersonDevice
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.AcPersonDeviceApp.Insert(data)
}

// UpdateAcPersonDevice 修改Device
func (p *AcPersonDeviceApi) UpdateAcPersonDevice(rc *restfulx.ReqCtx) {
	var data entity.AcPersonDevice
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.AcPersonDeviceApp.Update(data)
}

// DeleteAcPersonDevice 删除Device
func (p *AcPersonDeviceApi) DeleteAcPersonDevice(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.AcPersonDeviceApp.Delete(ids)
}

// ExportAcPersonDevice 导出Device
func (p *AcPersonDeviceApi) ExportAcPersonDevice(rc *restfulx.ReqCtx) {
    data := entity.AcPersonDevice{}

    filename := restfulx.QueryParam(rc, "filename")
    data.PersonUuid = restfulx.QueryParam(rc, "personUuid")

    list := p.AcPersonDeviceApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportAcPersonDevice 导入Device
func (p *AcPersonDeviceApi) ImportAcPersonDevice(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Device出现异常"), "导入Device出现异常")
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

        // 创建Device对象
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
		data := entity.AcPersonDevice{

		}

		// 保存Device对象到数据库中
	   //p.AcPersonDeviceApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
