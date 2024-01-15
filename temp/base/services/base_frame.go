// ==========================================================================
// 生成日期：2024-01-11 23:13:08 +0800 CST
// 生成路径: temp/base/services/base_frame.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/base/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseFrameModel interface {
		Insert(data entity.BaseFrame) *entity.BaseFrame
		FindOne(id int64) *entity.BaseFrame
		FindListPage(page, pageSize int, data entity.BaseFrame) (*[]entity.BaseFrame, int64)
		FindList(data entity.BaseFrame) *[]entity.BaseFrame
		Update(data entity.BaseFrame) *entity.BaseFrame
		Delete(ids []int64)
	}

	frameModelImpl struct {
		table string
	}
)


var BaseFrameModelDao BaseFrameModel = &frameModelImpl{
	table: `base_frame`,
}

func (m *frameModelImpl) Insert(data entity.BaseFrame) *entity.BaseFrame {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加房屋管理2失败")
	return &data
}

func (m *frameModelImpl) FindOne(id int64) *entity.BaseFrame {
	resData := new(entity.BaseFrame)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询房屋管理2失败")
	return resData
}

func (m *frameModelImpl) FindListPage(page, pageSize int, data entity.BaseFrame) (*[]entity.BaseFrame, int64) {
	list := make([]entity.BaseFrame, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.PUuid != "" {
        db = db.Where("p_uuid = ?", data.PUuid)
    }
    if data.ThirdId != "" {
        db = db.Where("third_id = ?", data.ThirdId)
    }
    if data.EntityCode != "" {
        db = db.Where("entity_code = ?", data.EntityCode)
    }
    if data.FrameCode != "" {
        db = db.Where("frame_code = ?", data.FrameCode)
    }
    if data.EntityName != "" {
        db = db.Where("entity_name like ?", "%"+data.EntityName+"%")
    }
    if data.FrameName != "" {
        db = db.Where("frame_name like ?", "%"+data.FrameName+"%")
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    if data.LevelCode != 0 {
        db = db.Where("level_code = ?", data.LevelCode)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询房屋管理2分页列表失败")
	return &list, total
}

func (m *frameModelImpl) FindList(data entity.BaseFrame) *[]entity.BaseFrame {
	list := make([]entity.BaseFrame, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.PUuid != "" {
        db = db.Where("p_uuid = ?", data.PUuid)
    }
    if data.ThirdId != "" {
        db = db.Where("third_id = ?", data.ThirdId)
    }
    if data.EntityCode != "" {
        db = db.Where("entity_code = ?", data.EntityCode)
    }
    if data.FrameCode != "" {
        db = db.Where("frame_code = ?", data.FrameCode)
    }
    if data.EntityName != "" {
        db = db.Where("entity_name like ?", "%"+data.EntityName+"%")
    }
    if data.FrameName != "" {
        db = db.Where("frame_name like ?", "%"+data.FrameName+"%")
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    if data.LevelCode != 0 {
        db = db.Where("level_code = ?", data.LevelCode)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询房屋管理2列表失败")
	return &list
}

func (m *frameModelImpl) Update(data entity.BaseFrame) *entity.BaseFrame {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改房屋管理2失败")
	return &data
}

func (m *frameModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseFrame{}, "id in (?)", ids).Error, "删除房屋管理2失败")
}
