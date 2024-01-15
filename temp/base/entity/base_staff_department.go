// ==========================================================================
// 生成日期：2024-01-11 23:18:43 +0800 CST
// 生成路径: temp/base/entity/base_staff_department.go
// 生成人：aurine
// ==========================================================================
package entity
import (
  "mod.miligc.com/edge-common/CommonKit/model"
  
)


type BaseStaffDepartment struct {
    Id  int64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`    // 序列
    Uuid   string   `gorm:"uuid;type:varchar(36);comment:Uuid" json:"uuid" `    // Uuid
    PUuid   string   `gorm:"p_uuid;type:varchar(36);comment:上级id" json:"pUuid" `    // 上级id
    DepartmentName   string   `gorm:"department_name;type:varchar(64);comment:部门名称" json:"departmentName" binding:"required"`    // 部门名称
    EmployeeCount   int   `gorm:"employee_count;type:int;comment:员工数量" json:"employeeCount" `    // 员工数量
    Desc   string   `gorm:"desc;type:varchar(500);comment:描述" json:"desc" `    // 描述
    ContactPhone   string   `gorm:"contact_phone;type:varchar(32);comment:联系电话" json:"contactPhone" `    // 联系电话
    ProjectUuid   string   `gorm:"project_uuid;type:varchar(36);comment:项目ID" json:"projectUuid" `    // 项目ID
    CreateBy   string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
    UpdateBy   string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
    model.BaseModel
}

func (BaseStaffDepartment) TableName() string {
	return "base_staff_department"
}
