package middleware

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
)

/**
 * @Description swagger文档配置
 * @Author aurine
 * @Date 2022/8/3 9:16
 **/

func SwaggerConfig(wsContainer *restful.Container) {
	config := restfulspec.Config{
		WebServices:                   wsContainer.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	wsContainer.Add(restfulspec.NewOpenAPIService(config))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "边缘开发 框架的API文档",
			Description: "这是 边缘开发 框架文档，根据文档调用API",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "边缘开发",
					Email: "aurinerd@aurine.cn",
					URL:   "http://www.aurine.cn",
				},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "MIT",
					URL:  "http://www.aurine.cn",
				},
			},
			Version: "1.0.0",
		},
	}
}
