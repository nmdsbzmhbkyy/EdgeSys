package plugin

import (
	"EdgeSys/pkg/global"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"io/fs"
	"mod.miligc.com/edge-common/business-common/business/pkg"
	public "mod.miligc.com/edge-common/edgesys-common/pkg"
	"path/filepath"
	"plugin"
	"strings"
)

const address = "./plugin/"

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
	//		pkg.Log.Errorf("插件注册失败 pluginName: %s, 版本: %s, 原因: %s", pluginObj.PluginName, pluginObj.PluginVersion, e)
	//	} else {
	//		pkg.Log.Infof("插件注册成功 pluginName: %s, 版本: %s", pluginObj.PluginName, pluginObj.PluginVersion)
	//	}
	//	//加载路由
	//	p.InitRouter()
	//	pluginMap[pluginObj.PluginName] = pluginObj.PluginVersion
	//	//校验版本
	//	err := p.CheckVersion(pluginMap)
	//	if err != nil {
	//		pkg.Log.Errorf("插件校验失败 pluginName: %s, 版本: %s, 原因: %s", pluginObj.PluginName, pluginObj.PluginVersion, e)
	//	} else {
	//		pkg.Log.Infof("插件校验成功 pluginName: %s, 版本: %s", pluginObj.PluginName, pluginObj.PluginVersion)
	//	}
	//}
}

// InitPlugin 注册所有已存在的插件
func (_ InitModule) InitPlugin(c *restful.Container) {
	pkg.Container = c
	//需要加载的插件
	loadingModuleArr := global.Conf.Module
	for _, plugName := range loadingModuleArr {
		pkg.Log.Infoln("加载的插件: ", plugName)
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
			pkg.Log.Errorf("注册插件发生错误:%s", err.Error())
		}
	}
}

func register(path string, plugName string, c *restful.Container) error {
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

	obj := public.Module{}

	//加载插件
	pluginObj, e := p.InitPlugin(obj)

	//校验签名
	res := checkSignature(pluginObj)
	if !res {
		return errors.New("插件签名异常 pluginName: " + pluginObj.PluginName + ", 版本: " + pluginObj.PluginVersion)
	}

	pluginMap[pluginObj.PluginName] = pluginObj.PluginVersion

	//加载路由
	p.InitRouter()
	//校验版本
	requiresPluginMap := p.GetDependentVersion()
	for k, v := range requiresPluginMap {
		value := pluginMap[k]
		if pluginMap[k] == "" {
			return errors.New("插件校验失败 pluginName: " + pluginObj.PluginName + ", 版本: " + pluginObj.PluginVersion + ", 原因: 缺少依赖插件" + k + "-" + v)
		} else if value != v {
			return errors.New("插件校验失败 pluginName: " + pluginObj.PluginName + ", 版本: " + pluginObj.PluginVersion + ", 原因: 依赖的插件版本匹配错误 " + k + "-" + v)
		}
	}

	if e != nil {
		pkg.Log.Errorf("插件注册失败 pluginName: %s, 版本: %s, 原因: %s", pluginObj.PluginName, pluginObj.PluginVersion, e)
	} else {
		pkg.Log.Infof("插件注册成功 pluginName: %s, 版本: %s", pluginObj.PluginName, pluginObj.PluginVersion)
	}

	return nil
}

func checkSignature(p public.Plugin) bool {
	s := p.PluginName + "-" + p.PluginVersion
	// 使用 HMAC-SHA256 算法计算签名
	h := hmac.New(sha256.New, []byte(s))
	signature := hex.EncodeToString(h.Sum(nil))
	if signature == p.Secret {
		return true
	}
	return false
}
