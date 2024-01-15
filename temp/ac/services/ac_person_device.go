// ==========================================================================
// 生成日期：2024-01-11 23:59:35 +0800 CST
// 生成路径: temp/ac/services/ac_person_device.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/ac/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	AcPersonDeviceModel interface {
		Insert(data entity.AcPersonDevice) *entity.AcPersonDevice
		FindOne(id int64) *entity.AcPersonDevice
		FindListPage(page, pageSize int, data entity.AcPersonDevice) (*[]entity.AcPersonDevice, int64)
		FindList(data entity.AcPersonDevice) *[]entity.AcPersonDevice
		Update(data entity.AcPersonDevice) *entity.AcPersonDevice
		Delete(ids []int64)
	}

	personDeviceModelImpl struct {
		table string
	}
)


var AcPersonDeviceModelDao AcPersonDeviceModel = &personDeviceModelImpl{
	table: `ac_person_device`,
}

func (m *personDeviceModelImpl) Insert(data entity.AcPersonDevice) *entity.AcPersonDevice {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加人员权限失败")
	return &data
}

func (m *personDeviceModelImpl) FindOne(id int64) *entity.AcPersonDevice {
	resData := new(entity.AcPersonDevice)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询人员权限失败")
	return resData
}

func (m *personDeviceModelImpl) FindListPage(page, pageSize int, data entity.AcPersonDevice) (*[]entity.AcPersonDevice, int64) {
	list := make([]entity.AcPersonDevice, 0)
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
	biz.ErrIsNil(err, "查询人员权限分页列表失败")
	return &list, total
}

func (m *personDeviceModelImpl) FindList(data entity.AcPersonDevice) *[]entity.AcPersonDevice {
	list := make([]entity.AcPersonDevice, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.PersonUuid != "" {
        db = db.Where("person_uuid = ?", data.PersonUuid)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询人员权限列表失败")
	return &list
}

func (m *personDeviceModelImpl) Update(data entity.AcPersonDevice) *entity.AcPersonDevice {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改人员权限失败")
	return &data
}

func (m *personDeviceModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.AcPersonDevice{}, "id in (?)", ids).Error, "删除人员权限失败")
}
