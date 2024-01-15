// ==========================================================================
// 生成日期：2024-01-12 00:42:26 +0800 CST
// 生成路径: temp/cfg/services/base_project_cfg.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/cfg/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseProjectCfgModel interface {
		Insert(data entity.BaseProjectCfg) *entity.BaseProjectCfg
		FindOne(id int64) *entity.BaseProjectCfg
		FindListPage(page, pageSize int, data entity.BaseProjectCfg) (*[]entity.BaseProjectCfg, int64)
		FindList(data entity.BaseProjectCfg) *[]entity.BaseProjectCfg
		Update(data entity.BaseProjectCfg) *entity.BaseProjectCfg
		Delete(ids []int64)
	}

	projectCfgModelImpl struct {
		table string
	}
)


var BaseProjectCfgModelDao BaseProjectCfgModel = &projectCfgModelImpl{
	table: `base_project_cfg`,
}

func (m *projectCfgModelImpl) Insert(data entity.BaseProjectCfg) *entity.BaseProjectCfg {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加全局配置失败")
	return &data
}

func (m *projectCfgModelImpl) FindOne(id int64) *entity.BaseProjectCfg {
	resData := new(entity.BaseProjectCfg)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询全局配置失败")
	return resData
}

func (m *projectCfgModelImpl) FindListPage(page, pageSize int, data entity.BaseProjectCfg) (*[]entity.BaseProjectCfg, int64) {
	list := make([]entity.BaseProjectCfg, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询全局配置分页列表失败")
	return &list, total
}

func (m *projectCfgModelImpl) FindList(data entity.BaseProjectCfg) *[]entity.BaseProjectCfg {
	list := make([]entity.BaseProjectCfg, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询全局配置列表失败")
	return &list
}

func (m *projectCfgModelImpl) Update(data entity.BaseProjectCfg) *entity.BaseProjectCfg {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改全局配置失败")
	return &data
}

func (m *projectCfgModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseProjectCfg{}, "id in (?)", ids).Error, "删除全局配置失败")
}
