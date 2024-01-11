package services

import (
	"EdgeSys/apps/system/entity"
	"mod.miligc.com/edge-common/CommonKit/biz"
	"mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	SysNoticeModel interface {
		Insert(data entity.SysNotice) *entity.SysNotice
		FindOne(postId int64) *entity.SysNotice
		FindListPage(page, pageSize int, data entity.SysNotice) (*[]entity.SysNotice, int64)
		Update(data entity.SysNotice) *entity.SysNotice
		Delete(postId []int64)
	}

	sysNoticeModelImpl struct {
		table string
	}
)

var SysNoticeModelDao SysNoticeModel = &sysNoticeModelImpl{
	table: `sys_notices`,
}

func (m *sysNoticeModelImpl) Insert(data entity.SysNotice) *entity.SysNotice {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加通知失败")
	return &data
}

func (m *sysNoticeModelImpl) FindOne(postId int64) *entity.SysNotice {
	resData := new(entity.SysNotice)
	err := pkg.Db.Table(m.table).Where("post_id = ?", postId).First(resData).Error
	biz.ErrIsNil(err, "查询通知失败")
	return resData
}

func (m *sysNoticeModelImpl) FindListPage(page, pageSize int, data entity.SysNotice) (*[]entity.SysNotice, int64) {
	list := make([]entity.SysNotice, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Title != "" {
		db = db.Where("title like ?", "%"+data.Title+"%")
	}
	if data.NoticeType != "" {
		db = db.Where("notice_type = ?", data.NoticeType)
	}
	if len(data.OrganizationIds) > 0 {
		db = db.Where("organization_id in (?)", data.OrganizationIds)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询通知分页列表失败")
	return &list, total
}

func (m *sysNoticeModelImpl) Update(data entity.SysNotice) *entity.SysNotice {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改通知失败")
	return &data
}

func (m *sysNoticeModelImpl) Delete(postIds []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.SysNotice{}, "notice_id in (?)", postIds).Error, "删除通知失败")
}
