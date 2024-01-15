// ==========================================================================
// 生成日期：2024-01-12 00:41:39 +0800 CST
// 生成路径: temp/cfg/services/base_project.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/cfg/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseProjectModel interface {
		Insert(data entity.BaseProject) *entity.BaseProject
		FindOne(id int64) *entity.BaseProject
		FindListPage(page, pageSize int, data entity.BaseProject) (*[]entity.BaseProject, int64)
		FindList(data entity.BaseProject) *[]entity.BaseProject
		Update(data entity.BaseProject) *entity.BaseProject
		Delete(ids []int64)
	}

	projectModelImpl struct {
		table string
	}
)


var BaseProjectModelDao BaseProjectModel = &projectModelImpl{
	table: `base_project`,
}

func (m *projectModelImpl) Insert(data entity.BaseProject) *entity.BaseProject {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加基础信息失败")
	return &data
}

func (m *projectModelImpl) FindOne(id int64) *entity.BaseProject {
	resData := new(entity.BaseProject)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询基础信息失败")
	return resData
}

func (m *projectModelImpl) FindListPage(page, pageSize int, data entity.BaseProject) (*[]entity.BaseProject, int64) {
	list := make([]entity.BaseProject, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询基础信息分页列表失败")
	return &list, total
}

func (m *projectModelImpl) FindList(data entity.BaseProject) *[]entity.BaseProject {
	list := make([]entity.BaseProject, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询基础信息列表失败")
	return &list
}

func (m *projectModelImpl) Update(data entity.BaseProject) *entity.BaseProject {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改基础信息失败")
	return &data
}

func (m *projectModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseProject{}, "id in (?)", ids).Error, "删除基础信息失败")
}
