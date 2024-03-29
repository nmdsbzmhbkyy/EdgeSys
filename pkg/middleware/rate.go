package middleware

import (
	"EdgeSys/pkg/global"

	"github.com/didip/tollbooth"
	"github.com/emicklei/go-restful/v3"
)

/**
 * @Description 限流中间件
 * @Author aurine
 * @Date 2022/1/19 8:28
 **/

// Rate 限流中间件
func Rate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	lmt := tollbooth.NewLimiter(global.Conf.Server.Rate.RateNum, nil)
	lmt.SetMessage("已经超出接口请求限制，稍后再试.")
	httpError := tollbooth.LimitByRequest(lmt, resp, req.Request)
	if httpError != nil {
		resp.WriteErrorString(httpError.StatusCode, httpError.Message)
		return
	}
	chain.ProcessFilter(req, resp)
}
