// ==========================================================================
// 生成日期：2024-01-12 00:12:41 +0800 CST
// 生成路径: temp/ac/entity/ac_event_device_alarm.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
  
  "time"
)


type AcEventDeviceAlarm struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    DeviceUuid   string   `gorm:"device_uuid;type:varchar(36);comment:设备id" json:"deviceUuid" `    // 设备id
    DeviceName   string   `gorm:"device_name;type:varchar(128);comment:设备名称" json:"deviceName" `    // 设备名称
    EventType   string   `gorm:"event_type;type:varchar(1);comment:事件类型" json:"eventType" `    // 事件类型
    EventTime   time.Time   `gorm:"event_time;type:datetime;comment:事件时间" json:"eventTime" `    // 事件时间
    HandleStatus   string   `gorm:"handle_status;type:varchar(5);comment:事件处理状态" json:"handleStatus" `    // 事件处理状态
    HandlerId   int64   `gorm:"handler_id;type:bigint;comment:处理人id" json:"handlerId" `    // 处理人id
    HandleTime   time.Time   `gorm:"handle_time;type:datetime;comment:处理时间" json:"handleTime" `    // 处理时间
    Params   string   `gorm:"params;type:text;comment:告警参数，如室内机防区编号等" json:"params" `    // 告警参数，如室内机防区编号等
    Notes   string   `gorm:"notes;type:text;comment:备注" json:"notes" `    // 备注
    SnapshotUrl   string   `gorm:"snapshot_url;type:varchar(1000);comment:抓拍图片" json:"snapshotUrl" `    // 抓拍图片
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:ProjectUuid" json:"projectUuid" `    // ProjectUuid
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (AcEventDeviceAlarm) TableName() string {
	return "ac_event_device_alarm"
}
