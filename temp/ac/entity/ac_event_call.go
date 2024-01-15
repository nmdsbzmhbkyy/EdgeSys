// ==========================================================================
// 生成日期：2024-01-12 00:13:10 +0800 CST
// 生成路径: temp/ac/entity/ac_event_call.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
  
  "time"
)


type AcEventCall struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    CallerDeviceUuid   string   `gorm:"caller_device_uuid;type:varchar(36);comment:呼叫设备id" json:"callerDeviceUuid" `    // 呼叫设备id
    CalleeDeviceUuid   string   `gorm:"callee_device_uuid;type:varchar(36);comment:接听设备id" json:"calleeDeviceUuid" `    // 接听设备id
    CallerName   string   `gorm:"caller_name;type:varchar(200);comment:主叫名称" json:"callerName" `    // 主叫名称
    CalleeName   string   `gorm:"callee_name;type:varchar(200);comment:被叫名称" json:"calleeName" `    // 被叫名称
    CallType   string   `gorm:"call_type;type:varchar(5);comment:通话类型" json:"callType" `    // 通话类型
    CallResult   string   `gorm:"call_result;type:varchar(5);comment:通话结果" json:"callResult" `    // 通话结果
    StartTime   time.Time   `gorm:"start_time;type:datetime;comment:开始时间" json:"startTime" `    // 开始时间
    EndTime   time.Time   `gorm:"end_time;type:datetime;comment:结束时间" json:"endTime" `    // 结束时间
    CallDuration   int64   `gorm:"call_duration;type:bigint;comment:通话时长" json:"callDuration" `    // 通话时长
    Notes   string   `gorm:"notes;type:text;comment:备注" json:"notes" `    // 备注
    SnapshotUrl   string   `gorm:"snapshot_url;type:varchar(1000);comment:抓拍图片" json:"snapshotUrl" `    // 抓拍图片
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:ProjectUuid" json:"projectUuid" `    // ProjectUuid
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (AcEventCall) TableName() string {
	return "ac_event_call"
}
