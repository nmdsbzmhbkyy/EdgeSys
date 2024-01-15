// ==========================================================================
// 生成日期：2024-01-12 00:42:15 +0800 CST
// 生成路径: temp/cfg/entity/base_frame.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BaseFrame struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // 序列
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    PUuid   string   `gorm:"p_uuid;type:varchar(36);comment:上级id" json:"pUuid" `    // 上级id
    ThirdId   string   `gorm:"third_id;type:varchar(100);comment:第三方id,无对接为空" json:"thirdId" `    // 第三方id,无对接为空
    EntityCode   string   `gorm:"entity_code;type:varchar(64);comment:实体编码 如：02" json:"entityCode" `    // 实体编码 如：02
    FrameCode   string   `gorm:"frame_code;type:varchar(64);comment:框架编码 如：0102" json:"frameCode" `    // 框架编码 如：0102
    EntityName   string   `gorm:"entity_name;type:varchar(32);comment:实体名称" json:"entityName" `    // 实体名称
    FrameName   string   `gorm:"frame_name;type:varchar(32);comment:框架名称" json:"frameName" `    // 框架名称
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目ID" json:"projectUuid" `    // 项目ID
    LevelCode   int   `gorm:"level_code;type:int;comment:层级" json:"levelCode" `    // 层级
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseFrame) TableName() string {
	return "base_frame"
}
