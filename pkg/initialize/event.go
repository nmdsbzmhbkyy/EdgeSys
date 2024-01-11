/*
 * @Author: chensl chensl@aurine.cn
 * @Date: 2023-11-18 18:17:30
 * @LastEditors: chensl chensl@aurine.cn
 * @LastEditTime: 2023-11-19 22:16:40
 * @FilePath: \EdgeSys\pkg\initialize\event.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package initialize

import (
	"mod.miligc.com/edge-common/business-common/business/pkg"
)

// 初始化事件监听
func InitEvents() {
	// 监听**链改变 更新**规则
	pkg.EventBus.On("", func() {
		pkg.Log.Infof("**链变更")
		// ... 更新
	})
}
