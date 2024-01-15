// ==========================================================================
// 生成日期：2024-01-12 00:13:10 +0800 CST
// 生成路径: temp/ac/services/ac_event_call.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/ac/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	AcEventCallModel interface {
		Insert(data entity.AcEventCall) *entity.AcEventCall
		FindOne(id int64) *entity.AcEventCall
		FindListPage(page, pageSize int, data entity.AcEventCall) (*[]entity.AcEventCall, int64)
		FindList(data entity.AcEventCall) *[]entity.AcEventCall
		Update(data entity.AcEventCall) *entity.AcEventCall
		Delete(ids []int64)
	}

	eventCallModelImpl struct {
		table string
	}
)


var AcEventCallModelDao AcEventCallModel = &eventCallModelImpl{
	table: `ac_event_call`,
}

func (m *eventCallModelImpl) Insert(data entity.AcEventCall) *entity.AcEventCall {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加呼叫记录失败")
	return &data
}

func (m *eventCallModelImpl) FindOne(id int64) *entity.AcEventCall {
	resData := new(entity.AcEventCall)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询呼叫记录失败")
	return resData
}

func (m *eventCallModelImpl) FindListPage(page, pageSize int, data entity.AcEventCall) (*[]entity.AcEventCall, int64) {
	list := make([]entity.AcEventCall, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.CallType != "" {
        db = db.Where("call_type = ?", data.CallType)
    }
    if data.CallResult != "" {
        db = db.Where("call_result = ?", data.CallResult)
    }
    if data.CallDuration != 0 {
        db = db.Where("call_duration = ?", data.CallDuration)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询呼叫记录分页列表失败")
	return &list, total
}

func (m *eventCallModelImpl) FindList(data entity.AcEventCall) *[]entity.AcEventCall {
	list := make([]entity.AcEventCall, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.CallType != "" {
        db = db.Where("call_type = ?", data.CallType)
    }
    if data.CallResult != "" {
        db = db.Where("call_result = ?", data.CallResult)
    }
    if data.CallDuration != 0 {
        db = db.Where("call_duration = ?", data.CallDuration)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询呼叫记录列表失败")
	return &list
}

func (m *eventCallModelImpl) Update(data entity.AcEventCall) *entity.AcEventCall {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改呼叫记录失败")
	return &data
}

func (m *eventCallModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.AcEventCall{}, "id in (?)", ids).Error, "删除呼叫记录失败")
}
