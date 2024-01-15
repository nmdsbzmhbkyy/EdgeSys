// ==========================================================================
// 生成日期：2024-01-11 23:15:00 +0800 CST
// 生成路径: temp/base/services/base_region.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/base/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseRegionModel interface {
		Insert(data entity.BaseRegion) *entity.BaseRegion
		FindOne(id int64) *entity.BaseRegion
		FindListPage(page, pageSize int, data entity.BaseRegion) (*[]entity.BaseRegion, int64)
		FindList(data entity.BaseRegion) *[]entity.BaseRegion
		Update(data entity.BaseRegion) *entity.BaseRegion
		Delete(ids []int64)
	}

	regionModelImpl struct {
		table string
	}
)


var BaseRegionModelDao BaseRegionModel = &regionModelImpl{
	table: `base_region`,
}

func (m *regionModelImpl) Insert(data entity.BaseRegion) *entity.BaseRegion {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加区域设置失败")
	return &data
}

func (m *regionModelImpl) FindOne(id int64) *entity.BaseRegion {
	resData := new(entity.BaseRegion)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询区域设置失败")
	return resData
}

func (m *regionModelImpl) FindListPage(page, pageSize int, data entity.BaseRegion) (*[]entity.BaseRegion, int64) {
	list := make([]entity.BaseRegion, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Pid != 0 {
        db = db.Where("pid = ?", data.Pid)
    }
    if data.Name != "" {
        db = db.Where("name like ?", "%"+data.Name+"%")
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询区域设置分页列表失败")
	return &list, total
}

func (m *regionModelImpl) FindList(data entity.BaseRegion) *[]entity.BaseRegion {
	list := make([]entity.BaseRegion, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Pid != 0 {
        db = db.Where("pid = ?", data.Pid)
    }
    if data.Name != "" {
        db = db.Where("name like ?", "%"+data.Name+"%")
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询区域设置列表失败")
	return &list
}

func (m *regionModelImpl) Update(data entity.BaseRegion) *entity.BaseRegion {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改区域设置失败")
	return &data
}

func (m *regionModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseRegion{}, "id in (?)", ids).Error, "删除区域设置失败")
}
