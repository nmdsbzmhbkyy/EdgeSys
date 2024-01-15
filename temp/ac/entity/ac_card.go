// ==========================================================================
// 生成日期：2024-01-12 00:05:03 +0800 CST
// 生成路径: temp/ac/entity/ac_card.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type AcCard struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    ThirdId   string   `gorm:"third_id;type:varchar(64);comment:ThirdId" json:"thirdId" `    // ThirdId
    PersonUuid   string   `gorm:"person_uuid;type:varchar(36);comment:人员id" json:"personUuid" `    // 人员id
    CardOrginalId   string   `gorm:"card_orginal_id;type:varchar(50);comment:卡号源id" json:"cardOrginalId" `    // 卡号源id
    CardNumber   string   `gorm:"card_number;type:varchar(50);comment:卡号" json:"cardNumber" binding:"required"`    // 卡号
    Type   string   `gorm:"type;type:varchar(5);comment:卡类型" json:"type" `    // 卡类型
    Status   string   `gorm:"status;type:varchar(5);comment:卡状态 正常，停用，黑卡" json:"status" binding:"required"`    // 卡状态 正常，停用，黑卡
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:ProjectUuid" json:"projectUuid" `    // ProjectUuid
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (AcCard) TableName() string {
	return "ac_card"
}
