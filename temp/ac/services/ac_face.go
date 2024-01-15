// ==========================================================================
// 生成日期：2024-01-12 00:02:25 +0800 CST
// 生成路径: temp/ac/services/ac_face.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/ac/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	AcFaceModel interface {
		Insert(data entity.AcFace) *entity.AcFace
		FindOne(id int64) *entity.AcFace
		FindListPage(page, pageSize int, data entity.AcFace) (*[]entity.AcFace, int64)
		FindList(data entity.AcFace) *[]entity.AcFace
		Update(data entity.AcFace) *entity.AcFace
		Delete(ids []int64)
	}

	faceModelImpl struct {
		table string
	}
)


var AcFaceModelDao AcFaceModel = &faceModelImpl{
	table: `ac_face`,
}

func (m *faceModelImpl) Insert(data entity.AcFace) *entity.AcFace {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加人员权限3失败")
	return &data
}

func (m *faceModelImpl) FindOne(id int64) *entity.AcFace {
	resData := new(entity.AcFace)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询人员权限3失败")
	return resData
}

func (m *faceModelImpl) FindListPage(page, pageSize int, data entity.AcFace) (*[]entity.AcFace, int64) {
	list := make([]entity.AcFace, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.ThirdId != "" {
        db = db.Where("third_id = ?", data.ThirdId)
    }
    if data.PersonUuid != "" {
        db = db.Where("person_uuid = ?", data.PersonUuid)
    }
    if data.ImageUrl != "" {
        db = db.Where("image_url = ?", data.ImageUrl)
    }
    if data.ImageCode != "" {
        db = db.Where("image_code = ?", data.ImageCode)
    }
    if data.Status != "" {
        db = db.Where("status = ?", data.Status)
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询人员权限3分页列表失败")
	return &list, total
}

func (m *faceModelImpl) FindList(data entity.AcFace) *[]entity.AcFace {
	list := make([]entity.AcFace, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.ThirdId != "" {
        db = db.Where("third_id = ?", data.ThirdId)
    }
    if data.PersonUuid != "" {
        db = db.Where("person_uuid = ?", data.PersonUuid)
    }
    if data.ImageUrl != "" {
        db = db.Where("image_url = ?", data.ImageUrl)
    }
    if data.ImageCode != "" {
        db = db.Where("image_code = ?", data.ImageCode)
    }
    if data.Status != "" {
        db = db.Where("status = ?", data.Status)
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询人员权限3列表失败")
	return &list
}

func (m *faceModelImpl) Update(data entity.AcFace) *entity.AcFace {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改人员权限3失败")
	return &data
}

func (m *faceModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.AcFace{}, "id in (?)", ids).Error, "删除人员权限3失败")
}
