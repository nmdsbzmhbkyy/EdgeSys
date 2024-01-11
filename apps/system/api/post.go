package api

import (
	"EdgeSys/apps/system/entity"
	"EdgeSys/apps/system/services"
	"EdgeSys/pkg/global"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"mod.miligc.com/edge-common/CommonKit/biz"
	"mod.miligc.com/edge-common/CommonKit/model"
	"mod.miligc.com/edge-common/CommonKit/restfulx"
	"mod.miligc.com/edge-common/CommonKit/utils"
	"mod.miligc.com/edge-common/business-common/business/pkg"
	"strconv"
)

type PostApi struct {
	PostApp services.SysPostModel
	UserApp services.SysUserModel
	RoleApp services.SysRoleModel
}

// GetPostList 职位列表数据
func (p *PostApi) GetPostList(rc *restfulx.ReqCtx) {

	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	status := restfulx.QueryParam(rc, "status")
	postName := restfulx.QueryParam(rc, "postName")
	postCode := restfulx.QueryParam(rc, "postCode")
	post := entity.SysPost{Status: status, PostName: postName, PostCode: postCode}

	list, total := p.PostApp.FindListPage(pageNum, pageSize, post)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetPost 获取职位
func (p *PostApi) GetPost(rc *restfulx.ReqCtx) {
	postId := restfulx.PathParamInt(rc, "postId")
	p.PostApp.FindOne(int64(postId))
}

// InsertPost 添加职位
func (p *PostApi) InsertPost(rc *restfulx.ReqCtx) {
	var post entity.SysPost
	restfulx.BindJsonAndValid(rc, &post)
	post.CreateBy = rc.LoginAccount.UserName
	p.PostApp.Insert(post)
}

// UpdatePost 修改职位
func (p *PostApi) UpdatePost(rc *restfulx.ReqCtx) {
	var post entity.SysPost
	restfulx.BindJsonAndValid(rc, &post)

	post.CreateBy = rc.LoginAccount.UserName
	p.PostApp.Update(post)
}

// DeletePost 删除职位
func (p *PostApi) DeletePost(rc *restfulx.ReqCtx) {
	postId := restfulx.PathParam(rc, "postId")
	postIds := utils.IdsStrToIdsIntGroup(postId)

	deList := make([]int64, 0)
	for _, id := range postIds {
		user := entity.SysUser{}
		user.PostId = id
		list := p.UserApp.FindList(user)
		if len(*list) == 0 {
			deList = append(deList, id)
		} else {
			pkg.Log.Info(fmt.Sprintf("dictId: %d 存在岗位绑定用户无法删除", id))
		}
	}
	if len(deList) == 0 {
		biz.ErrIsNil(errors.New("所有岗位都已绑定用户，无法删除"), "所有岗位都已绑定用户，无法删除")
	}
	p.PostApp.Delete(deList)
}

// ExportPost 导出岗位
func (p *PostApi) ExportPost(rc *restfulx.ReqCtx) {
	filename := restfulx.QueryParam(rc, "filename")
	postName := restfulx.QueryParam(rc, "postName")
	postCode := restfulx.QueryParam(rc, "postCode")
	status := restfulx.QueryParam(rc, "status")
	post := entity.SysPost{Status: status, PostName: postName, PostCode: postCode}

	list := p.PostApp.FindList(post)

	filename = utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, filename)
	rc.Download(filename)
}

// ImportPost 导入岗位
func (p *PostApi) ImportPost(rc *restfulx.ReqCtx) {
	// 从请求中读取Excel文件
	file, _, err := rc.Request.Request.FormFile("file")
	if err != nil {
		biz.ErrIsNil(errors.New("导入出现异常"), "导入出现异常")
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

		sort, _ := strconv.ParseInt(row[2], 10, 64)
		// 创建岗位对象
		post := entity.SysPost{
			PostName: row[0],
			PostCode: row[1],
			Sort:     sort,
			Status:   row[3],
			Remark:   row[4],
		}
		// 保存岗位对象到数据库中
		pkg.Log.Infoln(post)
		p.PostApp.Insert(post)
	}
}
