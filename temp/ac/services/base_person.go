// ==========================================================================
// 生成日期：2024-01-11 23:57:43 +0800 CST
// 生成路径: temp/ac/services/base_person.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/ac/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BasePersonModel interface {
		Insert(data entity.BasePerson) *entity.BasePerson
		FindOne(id int64) *entity.BasePerson
		FindListPage(page, pageSize int, data entity.BasePerson) (*[]entity.BasePerson, int64)
		FindList(data entity.BasePerson) *[]entity.BasePerson
		Update(data entity.BasePerson) *entity.BasePerson
		Delete(ids []int64)
	}

	personModelImpl struct {
		table string
	}
)


var BasePersonModelDao BasePersonModel = &personModelImpl{
	table: `base_person`,
}

func (m *personModelImpl) Insert(data entity.BasePerson) *entity.BasePerson {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加员工登记2失败")
	return &data
}

func (m *personModelImpl) FindOne(id int64) *entity.BasePerson {
	resData := new(entity.BasePerson)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询员工登记2失败")
	return resData
}

func (m *personModelImpl) FindListPage(page, pageSize int, data entity.BasePerson) (*[]entity.BasePerson, int64) {
	list := make([]entity.BasePerson, 0)
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
    if data.Name != "" {
        db = db.Where("name like ?", "%"+data.Name+"%")
    }
    if data.Phone != "" {
        db = db.Where("phone = ?", data.Phone)
    }
    if data.Status != "" {
        db = db.Where("status = ?", data.Status)
    }
    if data.Gender != "" {
        db = db.Where("gender = ?", data.Gender)
    }
    if data.CredentialType != "" {
        db = db.Where("credential_type = ?", data.CredentialType)
    }
    if data.CredentialNo != "" {
        db = db.Where("credential_no = ?", data.CredentialNo)
    }
    if data.Enable != "" {
        db = db.Where("enable = ?", data.Enable)
    }
    if data.Source != "" {
        db = db.Where("source = ?", data.Source)
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询员工登记2分页列表失败")
	return &list, total
}

func (m *personModelImpl) FindList(data entity.BasePerson) *[]entity.BasePerson {
	list := make([]entity.BasePerson, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.Uuid != "" {
        db = db.Where("uuid = ?", data.Uuid)
    }
    if data.ThirdId != "" {
        db = db.Where("third_id = ?", data.ThirdId)
    }
    if data.Name != "" {
        db = db.Where("name like ?", "%"+data.Name+"%")
    }
    if data.Phone != "" {
        db = db.Where("phone = ?", data.Phone)
    }
    if data.Status != "" {
        db = db.Where("status = ?", data.Status)
    }
    if data.Gender != "" {
        db = db.Where("gender = ?", data.Gender)
    }
    if data.CredentialType != "" {
        db = db.Where("credential_type = ?", data.CredentialType)
    }
    if data.CredentialNo != "" {
        db = db.Where("credential_no = ?", data.CredentialNo)
    }
    if data.Enable != "" {
        db = db.Where("enable = ?", data.Enable)
    }
    if data.Source != "" {
        db = db.Where("source = ?", data.Source)
    }
    if data.ProjectUuid != "" {
        db = db.Where("project_uuid = ?", data.ProjectUuid)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询员工登记2列表失败")
	return &list
}

func (m *personModelImpl) Update(data entity.BasePerson) *entity.BasePerson {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改员工登记2失败")
	return &data
}

func (m *personModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BasePerson{}, "id in (?)", ids).Error, "删除员工登记2失败")
}
