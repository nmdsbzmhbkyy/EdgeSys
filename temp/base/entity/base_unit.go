// ==========================================================================
// 生成日期：2024-01-11 23:06:12 +0800 CST
// 生成路径: temp/base/entity/base_unit.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BaseUnit struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // 序列
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    UnitCode   string   `gorm:"unit_code;type:varchar(64);comment:单元编码" json:"unitCode" `    // 单元编码
    FrameCode   string   `gorm:"frame_code;type:varchar(64);comment:框架编码" json:"frameCode" `    // 框架编码
    UnitName   string   `gorm:"unit_name;type:varchar(32);comment:单元名称" json:"unitName" binding:"required"`    // 单元名称
    AddressCode   string   `gorm:"address_code;type:varchar(128);comment:地址编码" json:"addressCode" `    // 地址编码
    PicUrl1   string   `gorm:"pic_url1;type:varchar(128);comment:单元图片1" json:"picUrl1" `    // 单元图片1
    PicUrl2   string   `gorm:"pic_url2;type:varchar(128);comment:单元图片2" json:"picUrl2" `    // 单元图片2
    PicUrl3   string   `gorm:"pic_url3;type:varchar(128);comment:单元图片3" json:"picUrl3" `    // 单元图片3
    Remark   string   `gorm:"remark;type:varchar(128);comment:备注" json:"remark" `    // 备注
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目ID" json:"projectUuid" `    // 项目ID
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseUnit) TableName() string {
	return "base_unit"
}
