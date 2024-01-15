// ==========================================================================
// 生成日期：2024-01-11 23:09:55 +0800 CST
// 生成路径: temp/base/entity/base_house.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BaseHouse struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // 序列
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    Code   string   `gorm:"code;type:varchar(64);comment:房屋编码" json:"code" `    // 房屋编码
    Floor   int   `gorm:"floor;type:int;comment:楼层" json:"floor" binding:"required"`    // 楼层
    HouseTypeUuid   string   `gorm:"house_type_uuid;type:varchar(36);comment:户型ID" json:"houseTypeUuid" `    // 户型ID
    HouseName   string   `gorm:"house_name;type:varchar(64);comment:房屋名称" json:"houseName" `    // 房屋名称
    HouseLabelCode   string   `gorm:"house_label_code;type:varchar(5);comment:房屋类别编码" json:"houseLabelCode" binding:"required"`    // 房屋类别编码
    HousePurposeCode   string   `gorm:"house_purpose_code;type:varchar(5);comment:房屋用途编码" json:"housePurposeCode" `    // 房屋用途编码
    Orientation   string   `gorm:"orientation;type:varchar(5);comment:房屋朝向" json:"orientation" `    // 房屋朝向
    PersonNumber   int   `gorm:"person_number;type:int;comment:最大居住/办公人数" json:"personNumber" `    // 最大居住/办公人数
    Note   string   `gorm:"note;type:varchar(128);comment:备注" json:"note" `    // 备注
    UsageType   string   `gorm:"usage_type;type:varchar(5);comment:1 自用 2 出租 3 闲置" json:"usageType" binding:"required"`    // 1 自用 2 出租 3 闲置
    LocationCode   string   `gorm:"location_code;type:varchar(64);comment:地址编码" json:"locationCode" `    // 地址编码
    EmployerId   string   `gorm:"employer_id;type:varchar(36);comment:实有单位id" json:"employerId" `    // 实有单位id
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目ID" json:"projectUuid" `    // 项目ID
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseHouse) TableName() string {
	return "base_house"
}
