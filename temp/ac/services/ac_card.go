// ==========================================================================
// 生成日期：2024-01-12 00:05:03 +0800 CST
// 生成路径: temp/ac/services/ac_card.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/ac/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	AcCardModel interface {
		Insert(data entity.AcCard) *entity.AcCard
		FindOne(id int64) *entity.AcCard
		FindListPage(page, pageSize int, data entity.AcCard) (*[]entity.AcCard, int64)
		FindList(data entity.AcCard) *[]entity.AcCard
		Update(data entity.AcCard) *entity.AcCard
		Delete(ids []int64)
	}

	cardModelImpl struct {
		table string
	}
)


var AcCardModelDao AcCardModel = &cardModelImpl{
	table: `ac_card`,
}

func (m *cardModelImpl) Insert(data entity.AcCard) *entity.AcCard {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加人员发卡失败")
	return &data
}

func (m *cardModelImpl) FindOne(id int64) *entity.AcCard {
	resData := new(entity.AcCard)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询人员发卡失败")
	return resData
}

func (m *cardModelImpl) FindListPage(page, pageSize int, data entity.AcCard) (*[]entity.AcCard, int64) {
	list := make([]entity.AcCard, 0)
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
    if data.CardNumber != "" {
        db = db.Where("card_number = ?", data.CardNumber)
    }
    if data.Type != "" {
        db = db.Where("type = ?", data.Type)
    }
    if data.Status != "" {
        db = db.Where("status = ?", data.Status)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询人员发卡分页列表失败")
	return &list, total
}

func (m *cardModelImpl) FindList(data entity.AcCard) *[]entity.AcCard {
	list := make([]entity.AcCard, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.ThirdId != "" {
        db = db.Where("third_id = ?", data.ThirdId)
    }
    if data.CardNumber != "" {
        db = db.Where("card_number = ?", data.CardNumber)
    }
    if data.Type != "" {
        db = db.Where("type = ?", data.Type)
    }
    if data.Status != "" {
        db = db.Where("status = ?", data.Status)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询人员发卡列表失败")
	return &list
}

func (m *cardModelImpl) Update(data entity.AcCard) *entity.AcCard {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改人员发卡失败")
	return &data
}

func (m *cardModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.AcCard{}, "id in (?)", ids).Error, "删除人员发卡失败")
}
