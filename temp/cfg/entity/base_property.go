// ==========================================================================
// 生成日期：2024-01-12 00:41:50 +0800 CST
// 生成路径: temp/cfg/entity/base_property.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BaseProperty struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // 序列
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    PropertyCompany   string   `gorm:"property_company;type:varchar(32);comment:物业公司" json:"propertyCompany" `    // 物业公司
    PropertyPrincipal   string   `gorm:"property_principal;type:varchar(32);comment:物业负责人" json:"propertyPrincipal" `    // 物业负责人
    PropertyCode   string   `gorm:"property_code;type:varchar(32);comment:物业组织机构代码" json:"propertyCode" `    // 物业组织机构代码
    PropertyPhone   string   `gorm:"property_phone;type:varchar(20);comment:物业公司电话" json:"propertyPhone" `    // 物业公司电话
    PropertyAddress   string   `gorm:"property_address;type:varchar(128);comment:物业公司地址" json:"propertyAddress" `    // 物业公司地址
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目ID" json:"projectUuid" `    // 项目ID
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseProperty) TableName() string {
	return "base_property"
}
