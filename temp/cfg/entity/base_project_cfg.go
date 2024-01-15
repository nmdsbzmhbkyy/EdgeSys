// ==========================================================================
// 生成日期：2024-01-12 00:42:26 +0800 CST
// 生成路径: temp/cfg/entity/base_project_cfg.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BaseProjectCfg struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:项目ID" json:"uuid" `    // 项目ID
    Type   string   `gorm:"type;type:varchar(1);comment:类型" json:"type" `    // 类型
    Name   string   `gorm:"name;type:varchar(30);comment:key" json:"name" `    // key
    Value   string   `gorm:"value;type:text;comment:value" json:"value" `    // value
    Desc   string   `gorm:"desc;type:varchar(200);comment:描述" json:"desc" `    // 描述
    DataType   string   `gorm:"data_type;type:varchar(15);comment:数据类型" json:"dataType" `    // 数据类型
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目ID" json:"projectUuid" `    // 项目ID
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseProjectCfg) TableName() string {
	return "base_project_cfg"
}
