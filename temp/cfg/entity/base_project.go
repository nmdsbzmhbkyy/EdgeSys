// ==========================================================================
// 生成日期：2024-01-12 00:41:39 +0800 CST
// 生成路径: temp/cfg/entity/base_project.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
  
  "time"
)


type BaseProject struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // 序列
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:项目ID" json:"uuid" `    // 项目ID
    ThirdId   string   `gorm:"third_id;type:varchar(64);comment:第三方项目编号" json:"thirdId" `    // 第三方项目编号
    ProjectCode   string   `gorm:"project_code;type:varchar(64);comment:第三方编码" json:"projectCode" `    // 第三方编码
    ProjectName   string   `gorm:"project_name;type:varchar(255);comment:项目名称" json:"projectName" `    // 项目名称
    ShortName   string   `gorm:"short_name;type:varchar(32);comment:项目简称" json:"shortName" `    // 项目简称
    ProjectType   string   `gorm:"project_type;type:varchar(5);comment:项目类型 1 正式 2 测试 3 演示" json:"projectType" `    // 项目类型 1 正式 2 测试 3 演示
    ContactPerson   string   `gorm:"contact_person;type:varchar(64);comment:联系人" json:"contactPerson" `    // 联系人
    ContactPhone   string   `gorm:"contact_phone;type:varchar(32);comment:联系电话" json:"contactPhone" `    // 联系电话
    FixPhone   string   `gorm:"fix_phone;type:varchar(20);comment:固定电话" json:"fixPhone" `    // 固定电话
    Acreage   float64   `gorm:"acreage;type:decimal(10,2);comment:建筑面积" json:"acreage" `    // 建筑面积
    PicPath   string   `gorm:"pic_path;type:varchar(255);comment:图片路径" json:"picPath" `    // 图片路径
    PoliceStation   string   `gorm:"police_station;type:varchar(32);comment:所属派出所" json:"policeStation" `    // 所属派出所
    PoliceCode   string   `gorm:"police_code;type:varchar(32);comment:公安机关编码" json:"policeCode" `    // 公安机关编码
    Remark   string   `gorm:"remark;type:varchar(128);comment:备注" json:"remark" `    // 备注
    EffTime   time.Time   `gorm:"eff_time;type:datetime;comment:生效时间" json:"effTime" `    // 生效时间
    ExpTime   time.Time   `gorm:"exp_time;type:datetime;comment:失效时间" json:"expTime" `    // 失效时间
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseProject) TableName() string {
	return "base_project"
}
