// ==========================================================================
// 生成日期：2024-01-11 23:18:43 +0800 CST
// 生成路径: temp/base/services/base_staff_department.go
// 生成人：aurine
// ==========================================================================

package services

import (
    "EdgeSys/temp/base/entity"
    "mod.miligc.com/edge-common/CommonKit/biz"
    "mod.miligc.com/edge-common/business-common/business/pkg"
)

type (
	BaseStaffDepartmentModel interface {
		Insert(data entity.BaseStaffDepartment) *entity.BaseStaffDepartment
		FindOne(id int64) *entity.BaseStaffDepartment
		FindListPage(page, pageSize int, data entity.BaseStaffDepartment) (*[]entity.BaseStaffDepartment, int64)
		FindList(data entity.BaseStaffDepartment) *[]entity.BaseStaffDepartment
		Update(data entity.BaseStaffDepartment) *entity.BaseStaffDepartment
		Delete(ids []int64)
	}

	staffDepartmentModelImpl struct {
		table string
	}
)


var BaseStaffDepartmentModelDao BaseStaffDepartmentModel = &staffDepartmentModelImpl{
	table: `base_staff_department`,
}

func (m *staffDepartmentModelImpl) Insert(data entity.BaseStaffDepartment) *entity.BaseStaffDepartment {
	err := pkg.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加部门管理失败")
	return &data
}

func (m *staffDepartmentModelImpl) FindOne(id int64) *entity.BaseStaffDepartment {
	resData := new(entity.BaseStaffDepartment)
	db := pkg.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询部门管理失败")
	return resData
}

func (m *staffDepartmentModelImpl) FindListPage(page, pageSize int, data entity.BaseStaffDepartment) (*[]entity.BaseStaffDepartment, int64) {
	list := make([]entity.BaseStaffDepartment, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.DepartmentName != "" {
        db = db.Where("department_name like ?", "%"+data.DepartmentName+"%")
    }
    if data.EmployeeCount != 0 {
        db = db.Where("employee_count = ?", data.EmployeeCount)
    }
    if data.Desc != "" {
        db = db.Where("desc = ?", data.Desc)
    }
    if data.ContactPhone != "" {
        db = db.Where("contact_phone = ?", data.ContactPhone)
    }
    db.Where("delete_time IS NULL")
	err := db.Count(&total).Error

	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询部门管理分页列表失败")
	return &list, total
}

func (m *staffDepartmentModelImpl) FindList(data entity.BaseStaffDepartment) *[]entity.BaseStaffDepartment {
	list := make([]entity.BaseStaffDepartment, 0)
	db := pkg.Db.Table(m.table)
	// 此处填写 where参数判断
    if data.DepartmentName != "" {
        db = db.Where("department_name like ?", "%"+data.DepartmentName+"%")
    }
    if data.EmployeeCount != 0 {
        db = db.Where("employee_count = ?", data.EmployeeCount)
    }
    if data.Desc != "" {
        db = db.Where("desc = ?", data.Desc)
    }
    if data.ContactPhone != "" {
        db = db.Where("contact_phone = ?", data.ContactPhone)
    }
    db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询部门管理列表失败")
	return &list
}

func (m *staffDepartmentModelImpl) Update(data entity.BaseStaffDepartment) *entity.BaseStaffDepartment {
	biz.ErrIsNil(pkg.Db.Table(m.table).Updates(&data).Error, "修改部门管理失败")
	return &data
}

func (m *staffDepartmentModelImpl) Delete(ids []int64) {
	biz.ErrIsNil(pkg.Db.Table(m.table).Delete(&entity.BaseStaffDepartment{}, "id in (?)", ids).Error, "删除部门管理失败")
}
