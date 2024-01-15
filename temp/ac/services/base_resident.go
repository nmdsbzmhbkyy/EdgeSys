// ==========================================================================
// 生成日期：2024-01-11 23:53:55 +0800 CST
// 生成路径: temp/ac/services/base_resident.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/ac/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseResidentModel interface {
		Insert(data entity.BaseResident) *entity.BaseResident
		FindOne(id int64) *entity.BaseResident
		FindListPage(page, pageSize int, data entity.BaseResident) (*[]entity.BaseResident, int64)
		FindList(data entity.BaseResident) *[]entity.BaseResident
		Update(data entity.BaseResident) *entity.BaseResident
		Delete(ids []int64)
	}

	residentModelImpl struct {
		table string
	}
)


var BaseResidentModelDao BaseResidentModel = &residentModelImpl{
	table: `base_resident`,
}

func (m *residentModelImpl) Insert(data entity.BaseResident) *entity.BaseResident {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加住户登记失败")
	return &data
}

func (m *residentModelImpl) FindOne(id int64) *entity.BaseResident {
	resData := new(entity.BaseResident)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询住户登记失败")
	return resData
}

func (m *residentModelImpl) FindListPage(page, pageSize int, data entity.BaseResident) (*[]entity.BaseResident, int64) {
	list := make([]entity.BaseResident, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.PersonUuid != "" {
        db = db.Where("person_uuid = ?", data.PersonUuid)
    }
    if data.HouseUuid != "" {
        db = db.Where("house_uuid = ?", data.HouseUuid)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询住户登记分页列表失败")
	return &list, total
}

func (m *residentModelImpl) FindList(data entity.BaseResident) *[]entity.BaseResident {
	list := make([]entity.BaseResident, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.PersonUuid != "" {
        db = db.Where("person_uuid = ?", data.PersonUuid)
    }
    if data.HouseUuid != "" {
        db = db.Where("house_uuid = ?", data.HouseUuid)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询住户登记列表失败")
	return &list
}

func (m *residentModelImpl) Update(data entity.BaseResident) *entity.BaseResident {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改住户登记失败")
	return &data
}

func (m *residentModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseResident{}, "id in (?)", ids).Error, "删除住户登记失败")
}
