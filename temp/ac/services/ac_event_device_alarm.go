// ==========================================================================
// 生成日期：2024-01-12 00:12:41 +0800 CST
// 生成路径: temp/ac/services/ac_event_device_alarm.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/ac/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	AcEventDeviceAlarmModel interface {
		Insert(data entity.AcEventDeviceAlarm) *entity.AcEventDeviceAlarm
		FindOne(id int64) *entity.AcEventDeviceAlarm
		FindListPage(page, pageSize int, data entity.AcEventDeviceAlarm) (*[]entity.AcEventDeviceAlarm, int64)
		FindList(data entity.AcEventDeviceAlarm) *[]entity.AcEventDeviceAlarm
		Update(data entity.AcEventDeviceAlarm) *entity.AcEventDeviceAlarm
		Delete(ids []int64)
	}

	eventDeviceAlarmModelImpl struct {
		table string
	}
)


var AcEventDeviceAlarmModelDao AcEventDeviceAlarmModel = &eventDeviceAlarmModelImpl{
	table: `ac_event_device_alarm`,
}

func (m *eventDeviceAlarmModelImpl) Insert(data entity.AcEventDeviceAlarm) *entity.AcEventDeviceAlarm {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加设备事件失败")
	return &data
}

func (m *eventDeviceAlarmModelImpl) FindOne(id int64) *entity.AcEventDeviceAlarm {
	resData := new(entity.AcEventDeviceAlarm)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询设备事件失败")
	return resData
}

func (m *eventDeviceAlarmModelImpl) FindListPage(page, pageSize int, data entity.AcEventDeviceAlarm) (*[]entity.AcEventDeviceAlarm, int64) {
	list := make([]entity.AcEventDeviceAlarm, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.DeviceName != "" {
        db = db.Where("device_name like ?", "%"+data.DeviceName+"%")
    }
    if data.EventType != "" {
        db = db.Where("event_type = ?", data.EventType)
    }
    if data.HandleStatus != "" {
        db = db.Where("handle_status = ?", data.HandleStatus)
    }
    if data.Notes != "" {
        db = db.Where("notes = ?", data.Notes)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询设备事件分页列表失败")
	return &list, total
}

func (m *eventDeviceAlarmModelImpl) FindList(data entity.AcEventDeviceAlarm) *[]entity.AcEventDeviceAlarm {
	list := make([]entity.AcEventDeviceAlarm, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.DeviceName != "" {
        db = db.Where("device_name like ?", "%"+data.DeviceName+"%")
    }
    if data.EventType != "" {
        db = db.Where("event_type = ?", data.EventType)
    }
    if data.HandleStatus != "" {
        db = db.Where("handle_status = ?", data.HandleStatus)
    }
    if data.Notes != "" {
        db = db.Where("notes = ?", data.Notes)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询设备事件列表失败")
	return &list
}

func (m *eventDeviceAlarmModelImpl) Update(data entity.AcEventDeviceAlarm) *entity.AcEventDeviceAlarm {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改设备事件失败")
	return &data
}

func (m *eventDeviceAlarmModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.AcEventDeviceAlarm{}, "id in (?)", ids).Error, "删除设备事件失败")
}
