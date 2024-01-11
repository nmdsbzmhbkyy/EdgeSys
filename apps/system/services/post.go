package services

import (
	"EdgeSys/apps/system/entity"
	"mod.miligc.com/edge-common/CommonKit/biz"
	"mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	SysPostModel interface {
		Insert(data entity.SysPost) *entity.SysPost
		FindOne(postId int64) *entity.SysPost
		FindListPage(page, pageSize int, data entity.SysPost) (*[]entity.SysPost, int64)
		FindList(data entity.SysPost) *[]entity.SysPost
		Update(data entity.SysPost) *entity.SysPost
		Delete(postId []int64)
	}

	sysPostModelImpl struct {
		table string
	}
)

var SysPostModelDao SysPostModel = &sysPostModelImpl{
	table: `sys_posts`,
}

func (m *sysPostModelImpl) Insert(data entity.SysPost) *entity.SysPost {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加岗位失败")
	return &data
}

func (m *sysPostModelImpl) FindOne(postId int64) *entity.SysPost {
	resData := new(entity.SysPost)
	err := pkg.Db.Table(m.table).Where("post_id = ?", postId).First(resData).Error
	biz.ErrIsNil(err, "查询岗位失败")
	return resData
}

func (m *sysPostModelImpl) FindListPage(page, pageSize int, data entity.SysPost) (*[]entity.SysPost, int64) {
	list := make([]entity.SysPost, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.PostId != 0 {
		db = db.Where("post_id = ?", data.PostId)
	}
	if data.PostName != "" {
		db = db.Where("post_name like ?", "%"+data.PostName+"%")
	}
	if data.PostCode != "" {
		db = db.Where("post_code like ?", "%"+data.PostCode+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Order("sort").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询岗位分页列表失败")
	return &list, total
}

func (m *sysPostModelImpl) FindList(data entity.SysPost) *[]entity.SysPost {
	list := make([]entity.SysPost, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.PostId != 0 {
		db = db.Where("post_id = ?", data.PostId)
	}
	if data.PostName != "" {
		db = db.Where("post_name = ?", data.PostName)
	}
	if data.PostCode != "" {
		db = db.Where("post_code = ?", data.PostCode)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("sort").Find(&list).Error, "查询岗位列表失败")
	return &list
}

func (m *sysPostModelImpl) Update(data entity.SysPost) *entity.SysPost {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改岗位失败")
	return &data
}

func (m *sysPostModelImpl) Delete(postIds []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.SysPost{}, "post_id in (?)", postIds).Error, "删除岗位失败")
}
