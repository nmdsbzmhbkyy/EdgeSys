// ==========================================================================
// 生成日期：2024-01-11 23:01:03 +0800 CST
// 生成路径: temp/base/entity/base_house_type.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BaseHouseType struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // 序列
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    DesignName   string   `gorm:"design_name;type:varchar(100);comment:户型名称" json:"designName" `    // 户型名称
    RoomTotal   int   `gorm:"room_total;type:tinyint;comment:房间数量" json:"roomTotal" `    // 房间数量
    HallTotal   int   `gorm:"hall_total;type:tinyint;comment:客厅数量" json:"hallTotal" `    // 客厅数量
    BathroomTotal   int   `gorm:"bathroom_total;type:tinyint;comment:卫生间数量" json:"bathroomTotal" `    // 卫生间数量
    KitchenTotal   int   `gorm:"kitchen_total;type:tinyint;comment:厨房数量" json:"kitchenTotal" `    // 厨房数量
    DesginDesc   string   `gorm:"desgin_desc;type:varchar(64);comment:户型描述" json:"desginDesc" `    // 户型描述
    Area   float64   `gorm:"area;type:decimal(8,2);comment:面积" json:"area" `    // 面积
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目ID" json:"projectUuid" `    // 项目ID
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseHouseType) TableName() string {
	return "base_house_type"
}
