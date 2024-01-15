// ==========================================================================
// 生成日期：2024-01-11 23:09:55 +0800 CST
// 生成路径: temp/base/services/base_house.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/base/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseHouseModel interface {
		Insert(data entity.BaseHouse) *entity.BaseHouse
		FindOne(id int64) *entity.BaseHouse
		FindListPage(page, pageSize int, data entity.BaseHouse) (*[]entity.BaseHouse, int64)
		FindList(data entity.BaseHouse) *[]entity.BaseHouse
		Update(data entity.BaseHouse) *entity.BaseHouse
		Delete(ids []int64)
	}

	houseModelImpl struct {
		table string
	}
)


var BaseHouseModelDao BaseHouseModel = &houseModelImpl{
	table: `base_house`,
}

func (m *houseModelImpl) Insert(data entity.BaseHouse) *entity.BaseHouse {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加房屋管理失败")
	return &data
}

func (m *houseModelImpl) FindOne(id int64) *entity.BaseHouse {
	resData := new(entity.BaseHouse)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询房屋管理失败")
	return resData
}

func (m *houseModelImpl) FindListPage(page, pageSize int, data entity.BaseHouse) (*[]entity.BaseHouse, int64) {
	list := make([]entity.BaseHouse, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.UsageType != "" {
        db = db.Where("usage_type = ?", data.UsageType)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询房屋管理分页列表失败")
	return &list, total
}

func (m *houseModelImpl) FindList(data entity.BaseHouse) *[]entity.BaseHouse {
	list := make([]entity.BaseHouse, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.UsageType != "" {
        db = db.Where("usage_type = ?", data.UsageType)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询房屋管理列表失败")
	return &list
}

func (m *houseModelImpl) Update(data entity.BaseHouse) *entity.BaseHouse {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改房屋管理失败")
	return &data
}

func (m *houseModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseHouse{}, "id in (?)", ids).Error, "删除房屋管理失败")
}
