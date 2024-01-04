package service

import "EdgeSys/pkg/global"

type TestService struct {
}

func (t *TestService) Test() {
	global.Log.Infoln("被组件调用了")
}
