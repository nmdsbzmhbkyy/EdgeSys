// ==========================================================================
// 生成日期：2024-01-12 00:25:58 +0800 CST
// 生成路径: temp/device/entity/base_device.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BaseDevice struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // Id
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    PUuid   string   `gorm:"p_uuid;type:varchar(36);comment:父id" json:"pUuid" `    // 父id
    ThirdId   string   `gorm:"third_id;type:varchar(64);comment:第三方id" json:"thirdId" `    // 第三方id
    Name   string   `gorm:"name;type:varchar(128);comment:设备名称" json:"name" binding:"required"`    // 设备名称
    Sn   string   `gorm:"sn;type:varchar(128);comment:sn" json:"sn" binding:"required"`    // sn
    Mac   string   `gorm:"mac;type:varchar(128);comment:mac" json:"mac" `    // mac
    Ipv4   string   `gorm:"ipv4;type:varchar(128);comment:ip" json:"ipv4" `    // ip
    Port   string   `gorm:"port;type:varchar(10);comment:端口" json:"port" `    // 端口
    Type   string   `gorm:"type;type:varchar(1);comment:设备类型" json:"type" binding:"required"`    // 设备类型
    Status   string   `gorm:"status;type:varchar(1);comment:状态：添加中，已添加，添加失败" json:"status" binding:"required"`    // 状态：添加中，已添加，添加失败
    Online   string   `gorm:"online;type:varchar(1);comment:是否在线" json:"online" `    // 是否在线
    Ability   string   `gorm:"ability;type:varchar(200);comment:能力" json:"ability" `    // 能力
    FrameUuid   string   `gorm:"frame_uuid;type:varchar(36);comment:框架id" json:"frameUuid" `    // 框架id
    ProductUuid   string   `gorm:"product_uuid;type:varchar(36);comment:产品" json:"productUuid" `    // 产品
    Username   string   `gorm:"username;type:varchar(100);comment:登陆账号" json:"username" `    // 登陆账号
    Password   string   `gorm:"password;type:varchar(100);comment:密码" json:"password" `    // 密码
    Source   string   `gorm:"source;type:varchar(5);comment:数据来源 手动,对接,自动" json:"source" `    // 数据来源 手动,对接,自动
    DlStatus   string   `gorm:"dl_status;type:varchar(1);comment:下发状态" json:"dlStatus" `    // 下发状态
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:ProjectUuid" json:"projectUuid" `    // ProjectUuid
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseDevice) TableName() string {
	return "base_device"
}
