package initialize

import (
	devEntity "EdgeSys/apps/develop/entity"
	jobEntity "EdgeSys/apps/job/entity"
	logEntity "EdgeSys/apps/log/entity"
	systemEntity "EdgeSys/apps/system/entity"
	"EdgeSys/pkg/global"
	"mod.miligc.com/edge-common/CommonKit/biz"
	"mod.miligc.com/edge-common/business-common/business/pkg"
)

// 初始化时如果没有表创建表
func InitTable() {
	m := global.Conf.Server
	if m.IsInitTable {
		biz.ErrIsNil(
			pkg.Db.AutoMigrate(
				//casbin.CasbinRule{},
				systemEntity.SysOrganization{},
				systemEntity.SysApi{},
				systemEntity.SysConfig{},
				systemEntity.SysDictType{},
				systemEntity.SysDictData{},
				systemEntity.SysUser{},
				systemEntity.SysTenants{},
				systemEntity.SysRole{},
				systemEntity.SysMenu{},
				systemEntity.SysPost{},
				systemEntity.SysRoleMenu{},
				systemEntity.SysRoleOrganization{},
				systemEntity.SysNotice{},

				logEntity.LogLogin{},
				logEntity.LogOper{},
				jobEntity.SysJob{},
				devEntity.DevGenTable{},
				devEntity.DevGenTableColumn{},
			),
			"初始化表失败")
	}
	// err := global.TdDb.CreateEventTable()
	// if err != nil {
	// 	pkg.Log.Panic(err)
	// }
}
