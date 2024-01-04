package business

import (
	"EdgeSys/pkg/global"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"io/fs"
	"mod.miligc.com/edge-common/edgesys-common/interface"
	public "mod.miligc.com/edge-common/edgesys-common/pkg"
	"path/filepath"
	"plugin"
	"strings"
)

const address = "./plugin/"

var PersonComponentInterface personInterface.PersonInterface

type InitModule struct {
}

func (_ InitModule) InitPluginDebug(c *restful.Container) {
	//module := []public.RegisterPlugin{
	//	{
	//		Name: "PersonModule",
	//		//S:    person.NewModule(),
	//	},
	//}
	//pluginMap := make(map[string]string)
	//
	//obj := public.Module{
	//	Container: c,
	//	Db:        global.MiliDb,
	//}
	//
	//for _, d := range module {
	//	p := d.S.(public.Interface)
	//	//加载插件
	//	pluginObj, e := p.InitPlugin(obj)
	//	if e != nil {
	//		global.Log.Errorf("插件注册失败 pluginName: %s, 版本: %s, 原因: %s", pluginObj.PluginName, pluginObj.PluginVersion, e)
	//	} else {
	//		global.Log.Infof("插件注册成功 pluginName: %s, 版本: %s", pluginObj.PluginName, pluginObj.PluginVersion)
	//	}
	//	//加载路由
	//	p.InitRouter()
	//	pluginMap[pluginObj.PluginName] = pluginObj.PluginVersion
	//	//校验版本
	//	err := p.CheckVersion(pluginMap)
	//	if err != nil {
	//		global.Log.Errorf("插件校验失败 pluginName: %s, 版本: %s, 原因: %s", pluginObj.PluginName, pluginObj.PluginVersion, e)
	//	} else {
	//		global.Log.Infof("插件校验成功 pluginName: %s, 版本: %s", pluginObj.PluginName, pluginObj.PluginVersion)
	//	}
	//}
}

// InitPlugin 注册所有已存在的插件
func (_ InitModule) InitPlugin(c *restful.Container) {
	//需要加载的插件
	loadingModuleArr := global.Conf.Module
	for _, plugName := range loadingModuleArr {
		global.Log.Infoln("加载的插件: ", plugName)
		err := filepath.Walk(address, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			name := strings.Replace(info.Name(), ".so", "", -1)
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".so") && name == plugName {
				err := register(path, name, c)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			global.Log.Errorf("注册插件发生错误:%s", err.Error())
		}
	}
}

func register(path string, plugName string, c *restful.Container) error {
	global.Log.Infoln("plugName", plugName)

	// 加载插件 参数为插件路径
	open, err := plugin.Open(path)
	if err != nil {
		return err
	}
	// 查找插件符号，可以是函数，结构体，属性等
	lookup, err := open.Lookup("NewModule")
	if err != nil {
		return err
	}
	// 断言为无参函数返回值为interface{}
	sym, ok := lookup.(func() interface{})
	if !ok {
		return fmt.Errorf("unexpected type from module symbol")
	}
	p := sym().(public.Interface)

	pluginMap := make(map[string]string)

	obj := public.Module{
		Container: c,
		Db:        global.MiliDb,
	}
	//加载插件
	pluginObj, e := p.InitPlugin(obj)
	if e != nil {
		global.Log.Errorf("插件注册失败 pluginName: %s, 版本: %s, 原因: %s", pluginObj.PluginName, pluginObj.PluginVersion, e)
	} else {
		global.Log.Infof("插件注册成功 pluginName: %s, 版本: %s", pluginObj.PluginName, pluginObj.PluginVersion)
	}
	//加载路由
	p.InitRouter()
	pluginMap[pluginObj.PluginName] = pluginObj.PluginVersion
	//校验版本
	checkVersionError := p.CheckVersion(pluginMap)
	if checkVersionError != nil {
		global.Log.Errorf("插件校验失败 pluginName: %s, 版本: %s, 原因: %s", pluginObj.PluginName, pluginObj.PluginVersion, e)
	} else {
		global.Log.Infof("插件校验成功 pluginName: %s, 版本: %s", pluginObj.PluginName, pluginObj.PluginVersion)
	}

	//加载组件接口
	err = initComponentInterFace(open, plugName)
	if err != nil {
		return err
	}

	return nil
}

func initComponentInterFace(open *plugin.Plugin, name string) error {
	lookup, e := open.Lookup("NewComponentInterFace")
	if e != nil {
		return e
	}
	sym, err := lookup.(func() interface{})
	if !err {
		return fmt.Errorf("unexpected type from module symbol")
	}
	if name == "PersonComponent-v1.0.0" {
		PersonComponentInterface = sym().(personInterface.PersonInterface)
	}
	return nil
}
