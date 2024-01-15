// ==========================================================================
// 生成日期：2024-01-11 23:25:46 +0800 CST
// 生成路径: temp/base/services/base_staff_position.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/base/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseStaffPositionModel interface {
		Insert(data entity.BaseStaffPosition) *entity.BaseStaffPosition
		FindOne(id int64) *entity.BaseStaffPosition
		FindListPage(page, pageSize int, data entity.BaseStaffPosition) (*[]entity.BaseStaffPosition, int64)
		FindList(data entity.BaseStaffPosition) *[]entity.BaseStaffPosition
		Update(data entity.BaseStaffPosition) *entity.BaseStaffPosition
		Delete(ids []int64)
	}

	staffPositionModelImpl struct {
		table string
	}
)


var BaseStaffPositionModelDao BaseStaffPositionModel = &staffPositionModelImpl{
	table: `base_staff_position`,
}

func (m *staffPositionModelImpl) Insert(data entity.BaseStaffPosition) *entity.BaseStaffPosition {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加岗位管理失败")
	return &data
}

func (m *staffPositionModelImpl) FindOne(id int64) *entity.BaseStaffPosition {
	resData := new(entity.BaseStaffPosition)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询岗位管理失败")
	return resData
}

func (m *staffPositionModelImpl) FindListPage(page, pageSize int, data entity.BaseStaffPosition) (*[]entity.BaseStaffPosition, int64) {
	list := make([]entity.BaseStaffPosition, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.PositionName != "" {
        db = db.Where("position_name like ?", "%"+data.PositionName+"%")
    }
    if data.PositionType != "" {
        db = db.Where("position_type = ?", data.PositionType)
    }
    if data.Desc != "" {
        db = db.Where("desc = ?", data.Desc)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询岗位管理分页列表失败")
	return &list, total
}

func (m *staffPositionModelImpl) FindList(data entity.BaseStaffPosition) *[]entity.BaseStaffPosition {
	list := make([]entity.BaseStaffPosition, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.PositionName != "" {
        db = db.Where("position_name like ?", "%"+data.PositionName+"%")
    }
    if data.PositionType != "" {
        db = db.Where("position_type = ?", data.PositionType)
    }
    if data.Desc != "" {
        db = db.Where("desc = ?", data.Desc)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询岗位管理列表失败")
	return &list
}

func (m *staffPositionModelImpl) Update(data entity.BaseStaffPosition) *entity.BaseStaffPosition {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改岗位管理失败")
	return &data
}

func (m *staffPositionModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseStaffPosition{}, "id in (?)", ids).Error, "删除岗位管理失败")
}
