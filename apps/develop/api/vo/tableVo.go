package vo

import "EdgeSys/apps/develop/entity"

/**
 * @Description
 * @Author aurine
 * @Date 2022/8/4 15:52
 **/

type TableInfoVo struct {
	List []entity.DevGenTableColumn `json:"list"`
	Info entity.DevGenTable         `json:"info"`
}
