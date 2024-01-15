// ==========================================================================
// 生成日期：2024-01-11 23:56:35 +0800 CST
// 生成路径: temp/ac/services/base_staff_jurisdiction.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/ac/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseStaffJurisdictionModel interface {
		Insert(data entity.BaseStaffJurisdiction) *entity.BaseStaffJurisdiction
		FindOne(id int64) *entity.BaseStaffJurisdiction
		FindListPage(page, pageSize int, data entity.BaseStaffJurisdiction) (*[]entity.BaseStaffJurisdiction, int64)
		FindList(data entity.BaseStaffJurisdiction) *[]entity.BaseStaffJurisdiction
		Update(data entity.BaseStaffJurisdiction) *entity.BaseStaffJurisdiction
		Delete(ids []int64)
	}

	staffJurisdictionModelImpl struct {
		table string
	}
)


var BaseStaffJurisdictionModelDao BaseStaffJurisdictionModel = &staffJurisdictionModelImpl{
	table: `base_staff_jurisdiction`,
}

func (m *staffJurisdictionModelImpl) Insert(data entity.BaseStaffJurisdiction) *entity.BaseStaffJurisdiction {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加员工登记失败")
	return &data
}

func (m *staffJurisdictionModelImpl) FindOne(id int64) *entity.BaseStaffJurisdiction {
	resData := new(entity.BaseStaffJurisdiction)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询员工登记失败")
	return resData
}

func (m *staffJurisdictionModelImpl) FindListPage(page, pageSize int, data entity.BaseStaffJurisdiction) (*[]entity.BaseStaffJurisdiction, int64) {
	list := make([]entity.BaseStaffJurisdiction, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.PersonUuid != "" {
        db = db.Where("person_uuid = ?", data.PersonUuid)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询员工登记分页列表失败")
	return &list, total
}

func (m *staffJurisdictionModelImpl) FindList(data entity.BaseStaffJurisdiction) *[]entity.BaseStaffJurisdiction {
	list := make([]entity.BaseStaffJurisdiction, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.PersonUuid != "" {
        db = db.Where("person_uuid = ?", data.PersonUuid)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询员工登记列表失败")
	return &list
}

func (m *staffJurisdictionModelImpl) Update(data entity.BaseStaffJurisdiction) *entity.BaseStaffJurisdiction {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改员工登记失败")
	return &data
}

func (m *staffJurisdictionModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseStaffJurisdiction{}, "id in (?)", ids).Error, "删除员工登记失败")
}
