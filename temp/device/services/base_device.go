// ==========================================================================
// 生成日期：2024-01-12 00:25:58 +0800 CST
// 生成路径: temp/device/services/base_device.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/device/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseDeviceModel interface {
		Insert(data entity.BaseDevice) *entity.BaseDevice
		FindOne(id int64) *entity.BaseDevice
		FindListPage(page, pageSize int, data entity.BaseDevice) (*[]entity.BaseDevice, int64)
		FindList(data entity.BaseDevice) *[]entity.BaseDevice
		Update(data entity.BaseDevice) *entity.BaseDevice
		Delete(ids []int64)
	}

	deviceModelImpl struct {
		table string
	}
)


var BaseDeviceModelDao BaseDeviceModel = &deviceModelImpl{
	table: `base_device`,
}

func (m *deviceModelImpl) Insert(data entity.BaseDevice) *entity.BaseDevice {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加设备管理失败")
	return &data
}

func (m *deviceModelImpl) FindOne(id int64) *entity.BaseDevice {
	resData := new(entity.BaseDevice)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询设备管理失败")
	return resData
}

func (m *deviceModelImpl) FindListPage(page, pageSize int, data entity.BaseDevice) (*[]entity.BaseDevice, int64) {
	list := make([]entity.BaseDevice, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Name != "" {
        db = db.Where("name like ?", "%"+data.Name+"%")
    }
    if data.Sn != "" {
        db = db.Where("sn = ?", data.Sn)
    }
    if data.Ipv4 != "" {
        db = db.Where("ipv4 = ?", data.Ipv4)
    }
    if data.Status != "" {
        db = db.Where("status = ?", data.Status)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询设备管理分页列表失败")
	return &list, total
}

func (m *deviceModelImpl) FindList(data entity.BaseDevice) *[]entity.BaseDevice {
	list := make([]entity.BaseDevice, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Name != "" {
        db = db.Where("name like ?", "%"+data.Name+"%")
    }
    if data.Sn != "" {
        db = db.Where("sn = ?", data.Sn)
    }
    if data.Ipv4 != "" {
        db = db.Where("ipv4 = ?", data.Ipv4)
    }
    if data.Status != "" {
        db = db.Where("status = ?", data.Status)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询设备管理列表失败")
	return &list
}

func (m *deviceModelImpl) Update(data entity.BaseDevice) *entity.BaseDevice {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改设备管理失败")
	return &data
}

func (m *deviceModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseDevice{}, "id in (?)", ids).Error, "删除设备管理失败")
}
