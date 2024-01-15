// ==========================================================================
// 生成日期：2024-01-12 00:12:57 +0800 CST
// 生成路径: temp/ac/entity/ac_event_access.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
  "time"
)


type AcEventAccess struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    DeviceUuid   string   `gorm:"device_uuid;type:varchar(36);comment:设备id" json:"deviceUuid" `    // 设备id
    DeviceName   string   `gorm:"device_name;type:varchar(128);comment:设备名称" json:"deviceName" `    // 设备名称
    DeviceFrameDesc   string   `gorm:"device_frame_desc;type:varchar(1000);comment:设备位置" json:"deviceFrameDesc" `    // 设备位置
    CertType   string   `gorm:"cert_type;type:varchar(5);comment:凭证类型" json:"certType" `    // 凭证类型
    CertCode   string   `gorm:"cert_code;type:varchar(1000);comment:凭证数据" json:"certCode" `    // 凭证数据
    PersonUuid   string   `gorm:"person_uuid;type:varchar(36);comment:通行人员" json:"personUuid" `    // 通行人员
    PersonName   string   `gorm:"person_name;type:varchar(128);comment:手机号" json:"personName" `    // 手机号
    PersonFrameDesc   string   `gorm:"person_frame_desc;type:varchar(128);comment:人员位置描述" json:"personFrameDesc" `    // 人员位置描述
    CredentialNo   string   `gorm:"credential_no;type:varchar(255);comment:证件号码" json:"credentialNo" `    // 证件号码
    EventType   string   `gorm:"event_type;type:varchar(5);comment:事件类型" json:"eventType" `    // 事件类型
    EventTime   time.Time   `gorm:"event_time;type:datetime;comment:事件时间" json:"eventTime" `    // 事件时间
    AccessType   string   `gorm:"access_type;type:varchar(5);comment:通行方式（开门方式）" json:"accessType" `    // 通行方式（开门方式）
    SnapshotUrl   string   `gorm:"snapshot_url;type:varchar(1000);comment:抓拍图片" json:"snapshotUrl" `    // 抓拍图片
    Notes   string   `gorm:"notes;type:text;comment:备注" json:"notes" `    // 备注
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:ProjectUuid" json:"projectUuid" `    // ProjectUuid
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (AcEventAccess) TableName() string {
	return "ac_event_access"
}
