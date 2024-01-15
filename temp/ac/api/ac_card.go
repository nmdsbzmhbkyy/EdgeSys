package api
// ==========================================================================
// 生成日期：2024-01-12 00:05:03 +0800 CST
// 生成路径: temp/ac/api/ac_card.go
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

type AcCardApi struct {
	AcCardApp services.AcCardModel
}

// GetAcCardList Card列表数据
func (p *AcCardApi) GetAcCardList(rc *restfulx.ReqCtx) {
    data := entity.AcCard{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.ThirdId = restfulx.QueryParam(rc, "thirdId")
    data.CardNumber = restfulx.QueryParam(rc, "cardNumber")
    data.Type = restfulx.QueryParam(rc, "type")
    data.Status = restfulx.QueryParam(rc, "status")

	list, total := p.AcCardApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetAcCard 获取Card
func (p *AcCardApi) GetAcCard(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.AcCardApp.FindOne(int64(id))
}

// InsertAcCard 添加Card
func (p *AcCardApi) InsertAcCard(rc *restfulx.ReqCtx) {
	var data entity.AcCard
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.AcCardApp.Insert(data)
}

// UpdateAcCard 修改Card
func (p *AcCardApi) UpdateAcCard(rc *restfulx.ReqCtx) {
	var data entity.AcCard
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.AcCardApp.Update(data)
}

// DeleteAcCard 删除Card
func (p *AcCardApi) DeleteAcCard(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.AcCardApp.Delete(ids)
}

// ExportAcCard 导出Card
func (p *AcCardApi) ExportAcCard(rc *restfulx.ReqCtx) {
    data := entity.AcCard{}

    filename := restfulx.QueryParam(rc, "filename")
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.ThirdId = restfulx.QueryParam(rc, "thirdId")
    data.CardNumber = restfulx.QueryParam(rc, "cardNumber")
    data.Type = restfulx.QueryParam(rc, "type")
    data.Status = restfulx.QueryParam(rc, "status")

    list := p.AcCardApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportAcCard 导入Card
func (p *AcCardApi) ImportAcCard(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Card出现异常"), "导入Card出现异常")
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

        // 创建Card对象
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
		data := entity.AcCard{

		}

		// 保存Card对象到数据库中
	   //p.AcCardApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
