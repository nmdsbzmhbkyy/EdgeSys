// ==========================================================================
// 生成日期：2024-01-11 23:02:38 +0800 CST
// 生成路径: temp/base/entity/base_building.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BaseBuilding struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // 序列
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    BuildingCode   string   `gorm:"building_code;type:varchar(64);comment:楼栋编码" json:"buildingCode" `    // 楼栋编码
    FrameCode   string   `gorm:"frame_code;type:varchar(64);comment:框架编码" json:"frameCode" `    // 框架编码
    BuildingName   string   `gorm:"building_name;type:varchar(32);comment:楼栋名称" json:"buildingName" binding:"required"`    // 楼栋名称
    FloorTotal   int   `gorm:"floor_total;type:smallint;comment:楼层数" json:"floorTotal" `    // 楼层数
    HouseTotal   int   `gorm:"house_total;type:int;comment:户数" json:"houseTotal" `    // 户数
    UnitTotal   int   `gorm:"unit_total;type:smallint;comment:单元数" json:"unitTotal" `    // 单元数
    UnitNameType   string   `gorm:"unit_name_type;type:varchar(5);comment:单元名称类型 1 数字 2 字母" json:"unitNameType" binding:"required"`    // 单元名称类型 1 数字 2 字母
    BuildingArea   float64   `gorm:"building_area;type:decimal(12,2);comment:楼栋面积" json:"buildingArea" `    // 楼栋面积
    BuildingEra   string   `gorm:"building_era;type:varchar(4);comment:楼栋年代,YYYY" json:"buildingEra" `    // 楼栋年代,YYYY
    BuildingType   string   `gorm:"building_type;type:varchar(5);comment:建筑类别" json:"buildingType" `    // 建筑类别
    FloorGround   int   `gorm:"floor_ground;type:smallint;comment:地面建筑物层数" json:"floorGround" `    // 地面建筑物层数
    FloorUnderGround   int   `gorm:"floor_under_ground;type:smallint;comment:地下建筑物层数" json:"floorUnderGround" `    // 地下建筑物层数
    LiveUnderGround   int   `gorm:"live_under_ground;type:smallint;comment:地下层居住层数" json:"liveUnderGround" `    // 地下层居住层数
    PicUrl1   string   `gorm:"pic_url1;type:varchar(128);comment:楼栋图片" json:"picUrl1" `    // 楼栋图片
    PicUrl2   string   `gorm:"pic_url2;type:varchar(128);comment:楼栋图片" json:"picUrl2" `    // 楼栋图片
    PicUrl3   string   `gorm:"pic_url3;type:varchar(128);comment:楼栋图片" json:"picUrl3" `    // 楼栋图片
    RegionId   string   `gorm:"region_id;type:varchar(32);comment:区域id" json:"regionId" `    // 区域id
    OverHead   string   `gorm:"over_head;type:char(1);comment:是否架空 1：有 ；0：无;" json:"overHead" `    // 是否架空 1：有 ；0：无;
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目ID" json:"projectUuid" `    // 项目ID
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseBuilding) TableName() string {
	return "base_building"
}
