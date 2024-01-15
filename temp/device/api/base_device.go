package api
// ==========================================================================
// 生成日期：2024-01-12 00:25:58 +0800 CST
// 生成路径: temp/device/api/base_device.go
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

type BaseDeviceApi struct {
	BaseDeviceApp services.BaseDeviceModel
}

// GetBaseDeviceList Device列表数据
func (p *BaseDeviceApi) GetBaseDeviceList(rc *restfulx.ReqCtx) {
    data := entity.BaseDevice{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.Name = restfulx.QueryParam(rc, "name")
    data.Sn = restfulx.QueryParam(rc, "sn")
    data.Ipv4 = restfulx.QueryParam(rc, "ipv4")
    data.Status = restfulx.QueryParam(rc, "status")

	list, total := p.BaseDeviceApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBaseDevice 获取Device
func (p *BaseDeviceApi) GetBaseDevice(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BaseDeviceApp.FindOne(int64(id))
}

// InsertBaseDevice 添加Device
func (p *BaseDeviceApi) InsertBaseDevice(rc *restfulx.ReqCtx) {
	var data entity.BaseDevice
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BaseDeviceApp.Insert(data)
}

// UpdateBaseDevice 修改Device
func (p *BaseDeviceApi) UpdateBaseDevice(rc *restfulx.ReqCtx) {
	var data entity.BaseDevice
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BaseDeviceApp.Update(data)
}

// DeleteBaseDevice 删除Device
func (p *BaseDeviceApi) DeleteBaseDevice(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BaseDeviceApp.Delete(ids)
}

// ExportBaseDevice 导出Device
func (p *BaseDeviceApi) ExportBaseDevice(rc *restfulx.ReqCtx) {
    data := entity.BaseDevice{}

    filename := restfulx.QueryParam(rc, "filename")
    data.Name = restfulx.QueryParam(rc, "name")
    data.Sn = restfulx.QueryParam(rc, "sn")
    data.Ipv4 = restfulx.QueryParam(rc, "ipv4")
    data.Status = restfulx.QueryParam(rc, "status")

    list := p.BaseDeviceApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBaseDevice 导入Device
func (p *BaseDeviceApi) ImportBaseDevice(rc *restfulx.ReqCtx) {
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
		data := entity.BaseDevice{

		}

		// 保存Device对象到数据库中
	   //p.BaseDeviceApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
