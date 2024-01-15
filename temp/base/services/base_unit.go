// ==========================================================================
// 生成日期：2024-01-11 23:06:12 +0800 CST
// 生成路径: temp/base/services/base_unit.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/base/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseUnitModel interface {
		Insert(data entity.BaseUnit) *entity.BaseUnit
		FindOne(id int64) *entity.BaseUnit
		FindListPage(page, pageSize int, data entity.BaseUnit) (*[]entity.BaseUnit, int64)
		FindList(data entity.BaseUnit) *[]entity.BaseUnit
		Update(data entity.BaseUnit) *entity.BaseUnit
		Delete(ids []int64)
	}

	unitModelImpl struct {
		table string
	}
)


var BaseUnitModelDao BaseUnitModel = &unitModelImpl{
	table: `base_unit`,
}

func (m *unitModelImpl) Insert(data entity.BaseUnit) *entity.BaseUnit {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加楼栋设置2失败")
	return &data
}

func (m *unitModelImpl) FindOne(id int64) *entity.BaseUnit {
	resData := new(entity.BaseUnit)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询楼栋设置2失败")
	return resData
}

func (m *unitModelImpl) FindListPage(page, pageSize int, data entity.BaseUnit) (*[]entity.BaseUnit, int64) {
	list := make([]entity.BaseUnit, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.UnitCode != "" {
        db = db.Where("unit_code = ?", data.UnitCode)
    }
    if data.FrameCode != "" {
        db = db.Where("frame_code = ?", data.FrameCode)
    }
    if data.UnitName != "" {
        db = db.Where("unit_name like ?", "%"+data.UnitName+"%")
    }
    if data.AddressCode != "" {
        db = db.Where("address_code = ?", data.AddressCode)
    }
    if data.PicUrl1 != "" {
        db = db.Where("pic_url1 = ?", data.PicUrl1)
    }
    if data.PicUrl2 != "" {
        db = db.Where("pic_url2 = ?", data.PicUrl2)
    }
    if data.PicUrl3 != "" {
        db = db.Where("pic_url3 = ?", data.PicUrl3)
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询楼栋设置2分页列表失败")
	return &list, total
}

func (m *unitModelImpl) FindList(data entity.BaseUnit) *[]entity.BaseUnit {
	list := make([]entity.BaseUnit, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.UnitCode != "" {
        db = db.Where("unit_code = ?", data.UnitCode)
    }
    if data.FrameCode != "" {
        db = db.Where("frame_code = ?", data.FrameCode)
    }
    if data.UnitName != "" {
        db = db.Where("unit_name like ?", "%"+data.UnitName+"%")
    }
    if data.AddressCode != "" {
        db = db.Where("address_code = ?", data.AddressCode)
    }
    if data.PicUrl1 != "" {
        db = db.Where("pic_url1 = ?", data.PicUrl1)
    }
    if data.PicUrl2 != "" {
        db = db.Where("pic_url2 = ?", data.PicUrl2)
    }
    if data.PicUrl3 != "" {
        db = db.Where("pic_url3 = ?", data.PicUrl3)
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询楼栋设置2列表失败")
	return &list
}

func (m *unitModelImpl) Update(data entity.BaseUnit) *entity.BaseUnit {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改楼栋设置2失败")
	return &data
}

func (m *unitModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseUnit{}, "id in (?)", ids).Error, "删除楼栋设置2失败")
}
