/*
 * @Author: chensl chensl@aurine.cn
 * @Date: 2023-11-18 18:17:29
 * @LastEditors: chensl chensl@aurine.cn
 * @LastEditTime: 2023-11-19 22:17:38
 * @FilePath: \EdgeSys\apps\develop\entity\dev_gen_table.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package entity

import "mod.miligc.com/edge-common/CommonKit/model"

type DevGenTable struct {
	TableId        int64               `gorm:"primaryKey;autoIncrement"   json:"tableId"` // 编号
	OrgId          int64               `gorm:"type:int;comment:机构ID"    json:"orgId"  `
	Owner          string              `gorm:"type:varchar(64);comment:创建者,所有者"  json:"owner"`
	TableName      string              `gorm:"table_name"         json:"tableName"`      // 表名称
	TableComment   string              `gorm:"table_comment"      json:"tableComment"`   // 表描述
	ClassName      string              `gorm:"class_name"         json:"className"`      // 实体类名称
	TplCategory    string              `gorm:"tpl_category"       json:"tplCategory"`    // 使用的模板（crud单表操作 tree树表操作）
	PackageName    string              `gorm:"package_name"       json:"packageName"`    // 生成包路径
	ModuleName     string              `gorm:"module_name"        json:"moduleName"`     // 生成模块名
	BusinessName   string              `gorm:"business_name"      json:"businessName"`   // 生成业务名
	FunctionName   string              `gorm:"function_name"      json:"functionName"`   // 生成功能名
	FunctionAuthor string              `gorm:"function_author"    json:"functionAuthor"` // 生成功能作者
	Options        string              `gorm:"options"            json:"options"`        // 其它生成选项
	Remark         string              `gorm:"remark"             json:"remark"`         // 备注
	PkColumn       string              `gorm:"pk_column;"         json:"pkColumn"`
	PkGoField      string              `gorm:"pk_go_field"        json:"pkGoField"`
	PkGoType       string              `gorm:"pk_go_type"         json:"pkGoType"`
	PkJsonField    string              `gorm:"pk_json_field"      json:"pkJsonField"`
	MenuGroup      string              `gorm:"menu_group"         json:"menuGroup"`
	Columns        []DevGenTableColumn `gorm:"-"                  json:"columns"` // 字段信息
	model.BaseModel

	RoleId int64 `gorm:"-"` // 角色数据权限
}

type DBTables struct {
	TableName      string `gorm:"column:TABLE_NAME" json:"tableName"`
	Engine         string `gorm:"column:ENGINE" json:"engine"`
	TableRows      string `gorm:"column:TABLE_ROWS" json:"tableRows"`
	TableCollation string `gorm:"column:TABLE_COLLATION" json:"tableCollation"`
	CreateTime     string `gorm:"column:CREATE_TIME" json:"createTime"`
	UpdateTime     string `gorm:"column:UPDATE_TIME" json:"updateTime"`
	TableComment   string `gorm:"column:TABLE_COMMENT" json:"tableComment"`
}
