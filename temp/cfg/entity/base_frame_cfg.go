// ==========================================================================
// 生成日期：2024-01-12 00:42:03 +0800 CST
// 生成路径: temp/cfg/entity/base_frame_cfg.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BaseFrameCfg struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // 序列
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    LevelCode   int   `gorm:"level_code;type:int;comment:层级编号" json:"levelCode" `    // 层级编号
    CodeRule   string   `gorm:"code_rule;type:varchar(10);comment:用于和第三方对接，编号位数" json:"codeRule" `    // 用于和第三方对接，编号位数
    LevelDesc   string   `gorm:"level_desc;type:varchar(64);comment:层级描述" json:"levelDesc" `    // 层级描述
    LevelType   string   `gorm:"level_type;type:varchar(1);comment:层级类型（楼栋、房屋、单元）" json:"levelType" `    // 层级类型（楼栋、房屋、单元）
    IsDisable   string   `gorm:"is_disable;type:varchar(1);comment:是否禁用 0 否 1 是" json:"isDisable" `    // 是否禁用 0 否 1 是
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目ID" json:"projectUuid" `    // 项目ID
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseFrameCfg) TableName() string {
	return "base_frame_cfg"
}
