// ==========================================================================
// 生成日期：2024-01-11 23:15:00 +0800 CST
// 生成路径: temp/base/entity/base_region.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
  "time"
)


type BaseRegion struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    Pid   int64   `gorm:"pid;type:bigint;comment:Pid" json:"pid" binding:"required"`    // Pid
    Name   string   `gorm:"name;type:varchar(64);comment:区域名称" json:"name" binding:"required"`    // 区域名称
    PicName   string   `gorm:"pic_name;type:varchar(50);comment:平面图名称" json:"picName" `    // 平面图名称
    PicUrl   string   `gorm:"pic_url;type:varchar(128);comment:平面图url" json:"picUrl" `    // 平面图url
    UploadTime   time.Time   `gorm:"upload_time;type:datetime;comment:上传时间" json:"uploadTime" `    // 上传时间
    IsDefault   string   `gorm:"is_default;type:char(1);comment:是否默认区域 1 是 0 否" json:"isDefault" `    // 是否默认区域 1 是 0 否
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目id" json:"projectUuid" `    // 项目id
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseRegion) TableName() string {
	return "base_region"
}
