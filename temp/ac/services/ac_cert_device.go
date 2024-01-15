// ==========================================================================
// 生成日期：2024-01-12 00:04:08 +0800 CST
// 生成路径: temp/ac/services/ac_cert_device.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/ac/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	AcCertDeviceModel interface {
		Insert(data entity.AcCertDevice) *entity.AcCertDevice
		FindOne(id int64) *entity.AcCertDevice
		FindListPage(page, pageSize int, data entity.AcCertDevice) (*[]entity.AcCertDevice, int64)
		FindList(data entity.AcCertDevice) *[]entity.AcCertDevice
		Update(data entity.AcCertDevice) *entity.AcCertDevice
		Delete(ids []int64)
	}

	certDeviceModelImpl struct {
		table string
	}
)


var AcCertDeviceModelDao AcCertDeviceModel = &certDeviceModelImpl{
	table: `ac_cert_device`,
}

func (m *certDeviceModelImpl) Insert(data entity.AcCertDevice) *entity.AcCertDevice {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加门禁权限失败")
	return &data
}

func (m *certDeviceModelImpl) FindOne(id int64) *entity.AcCertDevice {
	resData := new(entity.AcCertDevice)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询门禁权限失败")
	return resData
}

func (m *certDeviceModelImpl) FindListPage(page, pageSize int, data entity.AcCertDevice) (*[]entity.AcCertDevice, int64) {
	list := make([]entity.AcCertDevice, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.PersonName != "" {
        db = db.Where("person_name like ?", "%"+data.PersonName+"%")
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询门禁权限分页列表失败")
	return &list, total
}

func (m *certDeviceModelImpl) FindList(data entity.AcCertDevice) *[]entity.AcCertDevice {
	list := make([]entity.AcCertDevice, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.PersonName != "" {
        db = db.Where("person_name like ?", "%"+data.PersonName+"%")
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询门禁权限列表失败")
	return &list
}

func (m *certDeviceModelImpl) Update(data entity.AcCertDevice) *entity.AcCertDevice {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改门禁权限失败")
	return &data
}

func (m *certDeviceModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.AcCertDevice{}, "id in (?)", ids).Error, "删除门禁权限失败")
}
