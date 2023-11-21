package model

import (
	"EdgeSys/apps/system/entity"
	"EdgeSys/apps/system/services"
	"encoding/base64"
	"math/rand"
	"strings"
	"time"

	"github.com/PandaXGO/PandaKit/biz"
	"gorm.io/gorm"
)

func OrgAuthSet(tx *gorm.DB, roleId int64, owner string) {
	if roleId == 0 {
		return
	}
	role, err := services.SysRoleModelDao.FindOrganizationsByRoleId(roleId)
	biz.ErrIsNil(err, "查询角色数据权限失败")
	if role.DataScope != entity.SELFDATASCOPE {
		biz.IsTrue(len(role.Org) > 0, "该角色下未分配组织权限")
		tx.Where("org_id in (?)", strings.Split(role.Org, ","))
	} else {
		tx.Where("owner = ?", owner)
	}
}
func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	id := make([]byte, 7) // 由于base64编码会增加字符数，这里使用7个字节生成10位ID
	_, err := rand.Read(id)
	if err != nil {
		panic(err) // 错误处理，根据实际情况进行处理
	}
	return base64.URLEncoding.EncodeToString(id)[:10]
}
