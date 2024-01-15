// ==========================================================================
// 生成日期：2024-01-11 23:56:35 +0800 CST
// 生成路径: temp/ac/entity/base_staff_jurisdiction.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
  
  "time"
)


type BaseStaffJurisdiction struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    PersonUuid   string   `gorm:"person_uuid;type:varchar(36);comment:人员id" json:"personUuid" `    // 人员id
    FrameUuid   string   `gorm:"frame_uuid;type:varchar(36);comment:框架id" json:"frameUuid" `    // 框架id
    Status   string   `gorm:"status;type:varchar(1);comment:状态" json:"status" `    // 状态
    StartDate   time.Time   `gorm:"start_date;type:datetime;comment:工作开始时间" json:"startDate" binding:"required"`    // 工作开始时间
    EndDate   time.Time   `gorm:"end_date;type:datetime;comment:工作结束时间" json:"endDate" binding:"required"`    // 工作结束时间
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:ProjectUuid" json:"projectUuid" `    // ProjectUuid
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseStaffJurisdiction) TableName() string {
	return "base_staff_jurisdiction"
}
