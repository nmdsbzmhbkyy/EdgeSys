package api
// ==========================================================================
// 生成日期：2024-01-11 23:57:43 +0800 CST
// 生成路径: temp/ac/api/base_person.go
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

type BasePersonApi struct {
	BasePersonApp services.BasePersonModel
}

// GetBasePersonList Person列表数据
func (p *BasePersonApi) GetBasePersonList(rc *restfulx.ReqCtx) {
    data := entity.BasePerson{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.ThirdId = restfulx.QueryParam(rc, "thirdId")
    data.Name = restfulx.QueryParam(rc, "name")
    data.Phone = restfulx.QueryParam(rc, "phone")
    data.Status = restfulx.QueryParam(rc, "status")
    data.Gender = restfulx.QueryParam(rc, "gender")
    data.CredentialType = restfulx.QueryParam(rc, "credentialType")
    data.CredentialNo = restfulx.QueryParam(rc, "credentialNo")
    data.Enable = restfulx.QueryParam(rc, "enable")
    data.Source = restfulx.QueryParam(rc, "source")
    data.ProjectUuid = restfulx.QueryParam(rc, "projectUuid")

	list, total := p.BasePersonApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
    	Total: total,
    	PageNum: int64(pageNum),
    	PageSize: int64(pageSize),
    	Data: list,
    }
}

// GetBasePerson 获取Person
func (p *BasePersonApi) GetBasePerson(rc *restfulx.ReqCtx) {
    id := restfulx.PathParamInt(rc, "id")
    rc.ResData = p.BasePersonApp.FindOne(int64(id))
}

// InsertBasePerson 添加Person
func (p *BasePersonApi) InsertBasePerson(rc *restfulx.ReqCtx) {
	var data entity.BasePerson
	restfulx.BindQuery(rc, &data)
	data.CreateBy = rc.LoginAccount.UserName

	p.BasePersonApp.Insert(data)
}

// UpdateBasePerson 修改Person
func (p *BasePersonApi) UpdateBasePerson(rc *restfulx.ReqCtx) {
	var data entity.BasePerson
	restfulx.BindQuery(rc, &data)
	data.UpdateBy = rc.LoginAccount.UserName

	p.BasePersonApp.Update(data)
}

// DeleteBasePerson 删除Person
func (p *BasePersonApi) DeleteBasePerson(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc,"id")

	ids := utils.IdsStrToIdsIntGroup(id)
    p.BasePersonApp.Delete(ids)
}

// ExportBasePerson 导出Person
func (p *BasePersonApi) ExportBasePerson(rc *restfulx.ReqCtx) {
    data := entity.BasePerson{}

    filename := restfulx.QueryParam(rc, "filename")
    data.Uuid = restfulx.QueryParam(rc, "uuid")
    data.ThirdId = restfulx.QueryParam(rc, "thirdId")
    data.Name = restfulx.QueryParam(rc, "name")
    data.Phone = restfulx.QueryParam(rc, "phone")
    data.Status = restfulx.QueryParam(rc, "status")
    data.Gender = restfulx.QueryParam(rc, "gender")
    data.CredentialType = restfulx.QueryParam(rc, "credentialType")
    data.CredentialNo = restfulx.QueryParam(rc, "credentialNo")
    data.Enable = restfulx.QueryParam(rc, "enable")
    data.Source = restfulx.QueryParam(rc, "source")
    data.ProjectUuid = restfulx.QueryParam(rc, "projectUuid")

    list := p.BasePersonApp.FindList(data)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}


// ImportBasePerson 导入Person
func (p *BasePersonApi) ImportBasePerson(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入Person出现异常"), "导入Person出现异常")
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

        // 创建Person对象
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
		data := entity.BasePerson{

		}

		// 保存Person对象到数据库中
	   //p.BasePersonApp.Insert(data)
	    pkg.Log.Infoln(data)
	}
}
