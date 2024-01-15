// ==========================================================================
// 生成日期：2024-01-11 23:59:35 +0800 CST
// 生成路径: temp/ac/entity/ac_person_device.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
  
  "time"
)


type AcPersonDevice struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    PersonUuid   string   `gorm:"person_uuid;type:varchar(36);comment:人员id" json:"personUuid" `    // 人员id
    DeviceUuid   string   `gorm:"device_uuid;type:varchar(36);comment:设备id" json:"deviceUuid" `    // 设备id
    StartDate   time.Time   `gorm:"start_date;type:datetime;comment:授权开始时间" json:"startDate" `    // 授权开始时间
    EndDate   time.Time   `gorm:"end_date;type:datetime;comment:授权结束时间" json:"endDate" `    // 授权结束时间
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:ProjectUuid" json:"projectUuid" `    // ProjectUuid
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (AcPersonDevice) TableName() string {
	return "ac_person_device"
}
