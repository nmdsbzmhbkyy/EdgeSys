package api

import (
	"EdgeSys/apps/system/api/vo"
	entity "EdgeSys/apps/system/entity"
	services "EdgeSys/apps/system/services"

	"mod.miligc.com/edge-common/CommonKit/biz"
	"mod.miligc.com/edge-common/CommonKit/restfulx"
	"mod.miligc.com/edge-common/CommonKit/utils"
)

type MenuApi struct {
	MenuApp         services.SysMenuModel
	OrganizationApp services.SysOrganizationModel

	RoleMenuApp services.SysRoleMenuModel
	RoleApp     services.SysRoleModel
}

func (m *MenuApi) GetMenuTreeSelect(rc *restfulx.ReqCtx) {
	menuGroup := restfulx.QueryParam(rc, "menuGroup")
	label := m.MenuApp.SelectMenuLabel(entity.SysMenu{MenuGroup: menuGroup})
	rc.ResData = label
}

func (m *MenuApi) GetMenuRole(rc *restfulx.ReqCtx) {
	roleKey := restfulx.QueryParam(rc, "roleKey")
	biz.IsTrue(roleKey != "", "请传入角色Key")
	rc.ResData = Build(*m.MenuApp.SelectMenuRole(roleKey))
}

func (m *MenuApi) GetMenuTreeRoleSelect(rc *restfulx.ReqCtx) {
	roleId := restfulx.PathParamInt(rc, "roleId")

	result := m.MenuApp.SelectMenuLabel(entity.SysMenu{})
	menuIds := make([]int64, 0)
	if roleId != 0 {
		menuIds = m.RoleApp.GetRoleMeunId(entity.SysRole{RoleId: int64(roleId)})
	}
	rc.ResData = vo.MenuTreeVo{
		Menus:       *result,
		CheckedKeys: menuIds,
	}
}

func (m *MenuApi) GetMenuPaths(rc *restfulx.ReqCtx) {
	roleKey := restfulx.QueryParam(rc, "roleKey")
	biz.IsTrue(roleKey != "", "请传入角色Key")
	rc.ResData = m.RoleMenuApp.GetMenuPaths(entity.SysRoleMenu{RoleName: roleKey})
}

func (m *MenuApi) GetMenuList(rc *restfulx.ReqCtx) {
	menuName := restfulx.QueryParam(rc, "menuName")
	status := restfulx.QueryParam(rc, "status")
	menuGroup := restfulx.QueryParam(rc, "menuGroup")

	menu := entity.SysMenu{MenuName: menuName, Status: status, MenuGroup: menuGroup}
	if menu.MenuName == "" {
		rc.ResData = m.MenuApp.SelectMenu(menu)
	} else {
		rc.ResData = m.MenuApp.FindList(menu)
	}
}

func (m *MenuApi) GetMenu(rc *restfulx.ReqCtx) {
	menuId := restfulx.PathParamInt(rc, "menuId")

	rc.ResData = m.MenuApp.FindOne(int64(menuId))
}

func (m *MenuApi) InsertMenu(rc *restfulx.ReqCtx) {
	var menu entity.SysMenu
	restfulx.BindJsonAndValid(rc, &menu)
	menu.CreateBy = rc.LoginAccount.UserName
	m.MenuApp.Insert(menu)
	permis := m.RoleMenuApp.GetPermis(rc.LoginAccount.RoleId)
	menus := m.MenuApp.SelectMenuRole(rc.LoginAccount.RoleKey)
	rc.ResData = vo.MenuPermisVo{
		Menus:       Build(*menus),
		Permissions: permis,
	}
}

func (m *MenuApi) UpdateMenu(rc *restfulx.ReqCtx) {
	var menu entity.SysMenu
	restfulx.BindJsonAndValid(rc, &menu)
	menu.UpdateBy = rc.LoginAccount.UserName
	m.MenuApp.Update(menu)
	permis := m.RoleMenuApp.GetPermis(rc.LoginAccount.RoleId)
	menus := m.MenuApp.SelectMenuRole(rc.LoginAccount.RoleKey)
	rc.ResData = vo.MenuPermisVo{
		Menus:       Build(*menus),
		Permissions: permis,
	}
}

func (m *MenuApi) DeleteMenu(rc *restfulx.ReqCtx) {
	menuId := restfulx.PathParam(rc, "menuId")
	m.MenuApp.Delete(utils.IdsStrToIdsIntGroup(menuId))
}
