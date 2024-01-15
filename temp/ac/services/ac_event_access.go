// ==========================================================================
// 生成日期：2024-01-12 00:12:57 +0800 CST
// 生成路径: temp/ac/services/ac_event_access.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/ac/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	AcEventAccessModel interface {
		Insert(data entity.AcEventAccess) *entity.AcEventAccess
		FindOne(id int64) *entity.AcEventAccess
		FindListPage(page, pageSize int, data entity.AcEventAccess) (*[]entity.AcEventAccess, int64)
		FindList(data entity.AcEventAccess) *[]entity.AcEventAccess
		Update(data entity.AcEventAccess) *entity.AcEventAccess
		Delete(ids []int64)
	}

	eventAccessModelImpl struct {
		table string
	}
)


var AcEventAccessModelDao AcEventAccessModel = &eventAccessModelImpl{
	table: `ac_event_access`,
}

func (m *eventAccessModelImpl) Insert(data entity.AcEventAccess) *entity.AcEventAccess {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加人行记录失败")
	return &data
}

func (m *eventAccessModelImpl) FindOne(id int64) *entity.AcEventAccess {
	resData := new(entity.AcEventAccess)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询人行记录失败")
	return resData
}

func (m *eventAccessModelImpl) FindListPage(page, pageSize int, data entity.AcEventAccess) (*[]entity.AcEventAccess, int64) {
	list := make([]entity.AcEventAccess, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.PersonUuid != "" {
        db = db.Where("person_uuid = ?", data.PersonUuid)
    }
    if data.EventType != "" {
        db = db.Where("event_type = ?", data.EventType)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询人行记录分页列表失败")
	return &list, total
}

func (m *eventAccessModelImpl) FindList(data entity.AcEventAccess) *[]entity.AcEventAccess {
	list := make([]entity.AcEventAccess, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.PersonUuid != "" {
        db = db.Where("person_uuid = ?", data.PersonUuid)
    }
    if data.EventType != "" {
        db = db.Where("event_type = ?", data.EventType)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询人行记录列表失败")
	return &list
}

func (m *eventAccessModelImpl) Update(data entity.AcEventAccess) *entity.AcEventAccess {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改人行记录失败")
	return &data
}

func (m *eventAccessModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.AcEventAccess{}, "id in (?)", ids).Error, "删除人行记录失败")
}
