// ==========================================================================
// 生成日期：2024-01-12 00:04:08 +0800 CST
// 生成路径: temp/ac/entity/ac_cert_device.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
  
  "time"
)


type AcCertDevice struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    DeviceUuid   string   `gorm:"device_uuid;type:varchar(36);comment:设备id" json:"deviceUuid" `    // 设备id
    CertUuid   string   `gorm:"cert_uuid;type:varchar(36);comment:凭证id" json:"certUuid" `    // 凭证id
    CertType   string   `gorm:"cert_type;type:varchar(1);comment:凭证类型,卡片，人脸" json:"certType" `    // 凭证类型,卡片，人脸
    CertCode   string   `gorm:"cert_code;type:varchar(64);comment:凭证编码，卡片为卡号，人脸为凭证编号" json:"certCode" `    // 凭证编码，卡片为卡号，人脸为凭证编号
    DlOperate   string   `gorm:"dl_operate;type:varchar(1);comment:下发业务，添加，删除，冻结等" json:"dlOperate" `    // 下发业务，添加，删除，冻结等
    DlStatus   string   `gorm:"dl_status;type:varchar(1);comment:下发中，下发完成，下发失败" json:"dlStatus" `    // 下发中，下发完成，下发失败
    PersonId   int64   `gorm:"person_id;type:bigint;comment:人员id" json:"personId" `    // 人员id
    PersonType   string   `gorm:"person_type;type:varchar(5);comment:人员类型" json:"personType" `    // 人员类型
    PersonName   string   `gorm:"person_name;type:varchar(128);comment:人员名称" json:"personName" binding:"required"`    // 人员名称
    StartDate   time.Time   `gorm:"start_date;type:datetime;comment:StartDate" json:"startDate" `    // StartDate
    EndDate   time.Time   `gorm:"end_date;type:datetime;comment:EndDate" json:"endDate" `    // EndDate
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:ProjectUuid" json:"projectUuid" `    // ProjectUuid
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (AcCertDevice) TableName() string {
	return "ac_cert_device"
}
