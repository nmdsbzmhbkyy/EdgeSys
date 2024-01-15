package api
// ==========================================================================
// 生成日期：2024-01-12 00:26:09 +0800 CST
// 生成路径: temp/device/api/ac_cert_device.go
// 生成人：aurine
// ==========================================================================
import (
    "EdgeSys/temp/device/entity"
    "EdgeSys/temp/device/services"
    "EdgeSys/pkg/global"
    "github.com/xuri/excelize/v2"
    "errors"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/CommonKit/model"
    "mod.miligc.com/edge-common/CommonKit/restfulx"
    "mod.miligc.com/edge-common/CommonKit/utils"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type AcCertDeviceApi struct {
	AcCertDeviceApp services.AcCertDeviceModel
}

// GetAcCertDeviceList Device列表数据
func (p *AcCertDeviceApi) GetAcCertDeviceList(rc *restfulx.ReqCtx) {
    data := entity.AcCertDevice{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.CertType = restfulx.QueryParam(rc, "certType")
    data.DlStatus = restfulx.QueryParam(rc, "dlStatus")
    data.PersonName = restfulx.QueryParam(rc, "personName")

	list, total := p.AcCertDeviceApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetAcCertDevice 获取Device
func (p *AcCertDeviceApi) GetAcCertDevice(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.AcCertDeviceApp.FindOne(int64(id))
}

// InsertAcCertDevice 添加Device
func (p *AcCertDeviceApi) InsertAcCertDevice(rc *restfulx.ReqCtx) {
	var data entity.AcCertDevice
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.AcCertDeviceApp.Insert(data)
}

// UpdateAcCertDevice 修改Device
func (p *AcCertDeviceApi) UpdateAcCertDevice(rc *restfulx.ReqCtx) {
	var data entity.AcCertDevice
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.AcCertDeviceApp.Update(data)
}

// DeleteAcCertDevice 删除Device
func (p *AcCertDeviceApi) DeleteAcCertDevice(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.AcCertDeviceApp.Delete(ids)
}

// ExportAcCertDevice 导出Device
func (p *AcCertDeviceApi) ExportAcCertDevice(rc *restfulx.ReqCtx) {
    data := entity.AcCertDevice{}

    filename := restfulx.QueryParam(rc, "filename")
    data.CertType = restfulx.QueryParam(rc, "certType")
    data.DlStatus = restfulx.QueryParam(rc, "dlStatus")
    data.PersonName = restfulx.QueryParam(rc, "personName")

    list := p.AcCertDeviceApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportAcCertDevice 导入Device
func (p *AcCertDeviceApi) ImportAcCertDevice(rc *restfulx.ReqCtx) {
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
		data := entity.AcCertDevice{

		}

		// 保存Device对象到数据库中
	   //p.AcCertDeviceApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
