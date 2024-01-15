// ==========================================================================
// 生成日期：2024-01-12 00:42:03 +0800 CST
// 生成路径: temp/cfg/services/base_frame_cfg.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/cfg/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseFrameCfgModel interface {
		Insert(data entity.BaseFrameCfg) *entity.BaseFrameCfg
		FindOne(id int64) *entity.BaseFrameCfg
		FindListPage(page, pageSize int, data entity.BaseFrameCfg) (*[]entity.BaseFrameCfg, int64)
		FindList(data entity.BaseFrameCfg) *[]entity.BaseFrameCfg
		Update(data entity.BaseFrameCfg) *entity.BaseFrameCfg
		Delete(ids []int64)
	}

	framecfgModelImpl struct {
		table string
	}
)


var BaseFrameCfgModelDao BaseFrameCfgModel = &framecfgModelImpl{
	table: `base_frame_cfg`,
}

func (m *framecfgModelImpl) Insert(data entity.BaseFrameCfg) *entity.BaseFrameCfg {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加框架信息失败")
	return &data
}

func (m *framecfgModelImpl) FindOne(id int64) *entity.BaseFrameCfg {
	resData := new(entity.BaseFrameCfg)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询框架信息失败")
	return resData
}

func (m *framecfgModelImpl) FindListPage(page, pageSize int, data entity.BaseFrameCfg) (*[]entity.BaseFrameCfg, int64) {
	list := make([]entity.BaseFrameCfg, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.LevelCode != 0 {
        db = db.Where("level_code = ?", data.LevelCode)
    }
    if data.CodeRule != "" {
        db = db.Where("code_rule = ?", data.CodeRule)
    }
    if data.LevelDesc != "" {
        db = db.Where("level_desc = ?", data.LevelDesc)
    }
    if data.LevelType != "" {
        db = db.Where("level_type = ?", data.LevelType)
    }
    if data.IsDisable != "" {
        db = db.Where("is_disable = ?", data.IsDisable)
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询框架信息分页列表失败")
	return &list, total
}

func (m *framecfgModelImpl) FindList(data entity.BaseFrameCfg) *[]entity.BaseFrameCfg {
	list := make([]entity.BaseFrameCfg, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.LevelCode != 0 {
        db = db.Where("level_code = ?", data.LevelCode)
    }
    if data.CodeRule != "" {
        db = db.Where("code_rule = ?", data.CodeRule)
    }
    if data.LevelDesc != "" {
        db = db.Where("level_desc = ?", data.LevelDesc)
    }
    if data.LevelType != "" {
        db = db.Where("level_type = ?", data.LevelType)
    }
    if data.IsDisable != "" {
        db = db.Where("is_disable = ?", data.IsDisable)
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询框架信息列表失败")
	return &list
}

func (m *framecfgModelImpl) Update(data entity.BaseFrameCfg) *entity.BaseFrameCfg {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改框架信息失败")
	return &data
}

func (m *framecfgModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseFrameCfg{}, "id in (?)", ids).Error, "删除框架信息失败")
}
