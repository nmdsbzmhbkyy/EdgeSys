// ==========================================================================
// 生成日期：2024-01-11 23:25:46 +0800 CST
// 生成路径: temp/base/entity/base_staff_position.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BaseStaffPosition struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // 序列
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    PositionName   string   `gorm:"position_name;type:varchar(100);comment:岗位名称" json:"positionName" binding:"required"`    // 岗位名称
    PositionType   string   `gorm:"position_type;type:varchar(5);comment:岗位类型，来字字典" json:"positionType" binding:"required"`    // 岗位类型，来字字典
    Desc   string   `gorm:"desc;type:varchar(200);comment:描述" json:"desc" `    // 描述
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目ID" json:"projectUuid" `    // 项目ID
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseStaffPosition) TableName() string {
	return "base_staff_position"
}
