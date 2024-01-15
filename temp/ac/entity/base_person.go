// ==========================================================================
// 生成日期：2024-01-11 23:57:43 +0800 CST
// 生成路径: temp/ac/entity/base_person.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BasePerson struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    ThirdId   string   `gorm:"third_id;type:varchar(64);comment:第三方id" json:"thirdId" `    // 第三方id
    Name   string   `gorm:"name;type:varchar(128);comment:姓名" json:"name" binding:"required"`    // 姓名
    Phone   string   `gorm:"phone;type:varchar(20);comment:手机号" json:"phone" `    // 手机号
    Status   string   `gorm:"status;type:varchar(1);comment:状态：添加中，已添加，添加失败" json:"status" binding:"required"`    // 状态：添加中，已添加，添加失败
    Gender   string   `gorm:"gender;type:varchar(1);comment:性别" json:"gender" `    // 性别
    CredentialType   string   `gorm:"credential_type;type:varchar(5);comment:证件类型" json:"credentialType" `    // 证件类型
    CredentialNo   string   `gorm:"credential_no;type:varchar(255);comment:证件号码" json:"credentialNo" `    // 证件号码
    Enable   string   `gorm:"enable;type:varchar(1);comment:是否启用" json:"enable" `    // 是否启用
    Source   string   `gorm:"source;type:varchar(5);comment:数据来源 业务系统,对接,公安" json:"source" `    // 数据来源 业务系统,对接,公安
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目ID" json:"projectUuid" `    // 项目ID
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BasePerson) TableName() string {
	return "base_person"
}
