// ==========================================================================
// 生成日期：2024-01-11 23:01:03 +0800 CST
// 生成路径: temp/base/services/base_house_type.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/base/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseHouseTypeModel interface {
		Insert(data entity.BaseHouseType) *entity.BaseHouseType
		FindOne(id int64) *entity.BaseHouseType
		FindListPage(page, pageSize int, data entity.BaseHouseType) (*[]entity.BaseHouseType, int64)
		FindList(data entity.BaseHouseType) *[]entity.BaseHouseType
		Update(data entity.BaseHouseType) *entity.BaseHouseType
		Delete(ids []int64)
	}

	houseTypeModelImpl struct {
		table string
	}
)


var BaseHouseTypeModelDao BaseHouseTypeModel = &houseTypeModelImpl{
	table: `base_house_type`,
}

func (m *houseTypeModelImpl) Insert(data entity.BaseHouseType) *entity.BaseHouseType {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加户型设置失败")
	return &data
}

func (m *houseTypeModelImpl) FindOne(id int64) *entity.BaseHouseType {
	resData := new(entity.BaseHouseType)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询户型设置失败")
	return resData
}

func (m *houseTypeModelImpl) FindListPage(page, pageSize int, data entity.BaseHouseType) (*[]entity.BaseHouseType, int64) {
	list := make([]entity.BaseHouseType, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.DesignName != "" {
        db = db.Where("design_name like ?", "%"+data.DesignName+"%")
    }
    if data.DesginDesc != "" {
        db = db.Where("desgin_desc = ?", data.DesginDesc)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询户型设置分页列表失败")
	return &list, total
}

func (m *houseTypeModelImpl) FindList(data entity.BaseHouseType) *[]entity.BaseHouseType {
	list := make([]entity.BaseHouseType, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.DesignName != "" {
        db = db.Where("design_name like ?", "%"+data.DesignName+"%")
    }
    if data.DesginDesc != "" {
        db = db.Where("desgin_desc = ?", data.DesginDesc)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询户型设置列表失败")
	return &list
}

func (m *houseTypeModelImpl) Update(data entity.BaseHouseType) *entity.BaseHouseType {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改户型设置失败")
	return &data
}

func (m *houseTypeModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseHouseType{}, "id in (?)", ids).Error, "删除户型设置失败")
}
