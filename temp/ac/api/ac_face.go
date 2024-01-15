package api
// ==========================================================================
// 生成日期：2024-01-12 00:02:25 +0800 CST
// 生成路径: temp/ac/api/ac_face.go
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

type AcFaceApi struct {
	AcFaceApp services.AcFaceModel
}

// GetAcFaceList Face列表数据
func (p *AcFaceApi) GetAcFaceList(rc *restfulx.ReqCtx) {
    data := entity.AcFace{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.ThirdId = restfulx.QueryParam(rc, "thirdId")
    data.PersonUuid = restfulx.QueryParam(rc, "personUuid")
    data.ImageUrl = restfulx.QueryParam(rc, "imageUrl")
    data.ImageCode = restfulx.QueryParam(rc, "imageCode")
    data.Status = restfulx.QueryParam(rc, "status")
    data.ProjectUuid = restfulx.QueryParam(rc, "projectUuid")

	list, total := p.AcFaceApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetAcFace 获取Face
func (p *AcFaceApi) GetAcFace(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.AcFaceApp.FindOne(int64(id))
}

// InsertAcFace 添加Face
func (p *AcFaceApi) InsertAcFace(rc *restfulx.ReqCtx) {
	var data entity.AcFace
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.AcFaceApp.Insert(data)
}

// UpdateAcFace 修改Face
func (p *AcFaceApi) UpdateAcFace(rc *restfulx.ReqCtx) {
	var data entity.AcFace
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.AcFaceApp.Update(data)
}

// DeleteAcFace 删除Face
func (p *AcFaceApi) DeleteAcFace(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.AcFaceApp.Delete(ids)
}

// ExportAcFace 导出Face
func (p *AcFaceApi) ExportAcFace(rc *restfulx.ReqCtx) {
    data := entity.AcFace{}

    filename := restfulx.QueryParam(rc, "filename")
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.ThirdId = restfulx.QueryParam(rc, "thirdId")
    data.PersonUuid = restfulx.QueryParam(rc, "personUuid")
    data.ImageUrl = restfulx.QueryParam(rc, "imageUrl")
    data.ImageCode = restfulx.QueryParam(rc, "imageCode")
    data.Status = restfulx.QueryParam(rc, "status")
    data.ProjectUuid = restfulx.QueryParam(rc, "projectUuid")

    list := p.AcFaceApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportAcFace 导入Face
func (p *AcFaceApi) ImportAcFace(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Face出现异常"), "导入Face出现异常")
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

        // 创建Face对象
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
		data := entity.AcFace{

		}

		// 保存Face对象到数据库中
	   //p.AcFaceApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
