// ==========================================================================
// 生成日期：2024-01-12 00:02:25 +0800 CST
// 生成路径: temp/ac/entity/ac_face.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type AcFace struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    ThirdId   string   `gorm:"third_id;type:varchar(64);comment:ThirdId" json:"thirdId" `    // ThirdId
    PersonUuid   string   `gorm:"person_uuid;type:varchar(36);comment:人员id" json:"personUuid" `    // 人员id
    ImageUrl   string   `gorm:"image_url;type:varchar(500);comment:图片地址" json:"imageUrl" `    // 图片地址
    ImageCode   string   `gorm:"image_code;type:varchar(64);comment:人脸编号" json:"imageCode" `    // 人脸编号
    Status   string   `gorm:"status;type:varchar(5);comment:人脸开关" json:"status" binding:"required"`    // 人脸开关
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:ProjectUuid" json:"projectUuid" `    // ProjectUuid
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (AcFace) TableName() string {
	return "ac_face"
}
