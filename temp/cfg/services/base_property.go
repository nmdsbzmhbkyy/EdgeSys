// ==========================================================================
// 生成日期：2024-01-12 00:41:50 +0800 CST
// 生成路径: temp/cfg/services/base_property.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/cfg/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BasePropertyModel interface {
		Insert(data entity.BaseProperty) *entity.BaseProperty
		FindOne(id int64) *entity.BaseProperty
		FindListPage(page, pageSize int, data entity.BaseProperty) (*[]entity.BaseProperty, int64)
		FindList(data entity.BaseProperty) *[]entity.BaseProperty
		Update(data entity.BaseProperty) *entity.BaseProperty
		Delete(ids []int64)
	}

	propertyModelImpl struct {
		table string
	}
)


var BasePropertyModelDao BasePropertyModel = &propertyModelImpl{
	table: `base_property`,
}

func (m *propertyModelImpl) Insert(data entity.BaseProperty) *entity.BaseProperty {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加基础信息2失败")
	return &data
}

func (m *propertyModelImpl) FindOne(id int64) *entity.BaseProperty {
	resData := new(entity.BaseProperty)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询基础信息2失败")
	return resData
}

func (m *propertyModelImpl) FindListPage(page, pageSize int, data entity.BaseProperty) (*[]entity.BaseProperty, int64) {
	list := make([]entity.BaseProperty, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.PropertyCompany != "" {
        db = db.Where("property_company = ?", data.PropertyCompany)
    }
    if data.PropertyPrincipal != "" {
        db = db.Where("property_principal = ?", data.PropertyPrincipal)
    }
    if data.PropertyCode != "" {
        db = db.Where("property_code = ?", data.PropertyCode)
    }
    if data.PropertyPhone != "" {
        db = db.Where("property_phone = ?", data.PropertyPhone)
    }
    if data.PropertyAddress != "" {
        db = db.Where("property_address = ?", data.PropertyAddress)
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询基础信息2分页列表失败")
	return &list, total
}

func (m *propertyModelImpl) FindList(data entity.BaseProperty) *[]entity.BaseProperty {
	list := make([]entity.BaseProperty, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.PropertyCompany != "" {
        db = db.Where("property_company = ?", data.PropertyCompany)
    }
    if data.PropertyPrincipal != "" {
        db = db.Where("property_principal = ?", data.PropertyPrincipal)
    }
    if data.PropertyCode != "" {
        db = db.Where("property_code = ?", data.PropertyCode)
    }
    if data.PropertyPhone != "" {
        db = db.Where("property_phone = ?", data.PropertyPhone)
    }
    if data.PropertyAddress != "" {
        db = db.Where("property_address = ?", data.PropertyAddress)
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询基础信息2列表失败")
	return &list
}

func (m *propertyModelImpl) Update(data entity.BaseProperty) *entity.BaseProperty {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改基础信息2失败")
	return &data
}

func (m *propertyModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseProperty{}, "id in (?)", ids).Error, "删除基础信息2失败")
}
