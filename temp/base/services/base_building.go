// ==========================================================================
// 生成日期：2024-01-11 23:02:38 +0800 CST
// 生成路径: temp/base/services/base_building.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/base/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseBuildingModel interface {
		Insert(data entity.BaseBuilding) *entity.BaseBuilding
		FindOne(id int64) *entity.BaseBuilding
		FindListPage(page, pageSize int, data entity.BaseBuilding) (*[]entity.BaseBuilding, int64)
		FindList(data entity.BaseBuilding) *[]entity.BaseBuilding
		Update(data entity.BaseBuilding) *entity.BaseBuilding
		Delete(ids []int64)
	}

	buildingModelImpl struct {
		table string
	}
)


var BaseBuildingModelDao BaseBuildingModel = &buildingModelImpl{
	table: `base_building`,
}

func (m *buildingModelImpl) Insert(data entity.BaseBuilding) *entity.BaseBuilding {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加楼栋设置失败")
	return &data
}

func (m *buildingModelImpl) FindOne(id int64) *entity.BaseBuilding {
	resData := new(entity.BaseBuilding)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询楼栋设置失败")
	return resData
}

func (m *buildingModelImpl) FindListPage(page, pageSize int, data entity.BaseBuilding) (*[]entity.BaseBuilding, int64) {
	list := make([]entity.BaseBuilding, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.BuildingName != "" {
        db = db.Where("building_name like ?", "%"+data.BuildingName+"%")
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询楼栋设置分页列表失败")
	return &list, total
}

func (m *buildingModelImpl) FindList(data entity.BaseBuilding) *[]entity.BaseBuilding {
	list := make([]entity.BaseBuilding, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.BuildingName != "" {
        db = db.Where("building_name like ?", "%"+data.BuildingName+"%")
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询楼栋设置列表失败")
	return &list
}

func (m *buildingModelImpl) Update(data entity.BaseBuilding) *entity.BaseBuilding {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改楼栋设置失败")
	return &data
}

func (m *buildingModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseBuilding{}, "id in (?)", ids).Error, "删除楼栋设置失败")
}
