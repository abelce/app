package normal_context

import (
	"abelce/app/gateway/domain"
)

type NormalContext interface {
	GetConfig() *domain.Config
	GetParams() map[string]interface{}
	GetBody() interface{}
	GetResponse() interface{}
	GetRequestPipeline()
	GetParamsPipeline()
	GetResponsePipeline()
}

type Context struct {
	port             int
	config           *domain.Config
	params           map[string]interface{}
	reponse          interface{}
	body             interface{}
	requestPipeline  []interface{}
	responsePipeline []interface{}
	paramsPipeline   []interface{}
}

func NewNormalContext(port int, config *domain.Config) *Context {
	return &Context{
		port: port,
		config: config,
	}
}

// 获取端口
func (ctx *Context) GetPort() int {
	return ctx.port
}

// 获取系统配置
func (ctx *Context) GetConfig() *domain.Config {
	return ctx.config
} /*
 */

// 获取请求的参数
func (ctx *Context) GetParams() map[string]interface{} {
	return ctx.params
}

// 获取请求体的数据
func (ctx *Context) GetBody() interface{} {
	return ctx.body
}

// 获取请求成功的原始数据
func (ctx *Context) GetResponse() interface{} {
	return ctx.reponse
}

// 获取发起请求的pipeline
func (ctx *Context) GetRequestPipeline() []interface{} {
	return ctx.requestPipeline
}

// 获取发起请求参数的pipeline
func (ctx *Context) GetParamsPipeline() []interface{} {
	return ctx.paramsPipeline
}

// 获取请求成功后处理数据的pipeline
func (ctx *Context) GetResponsePipeline() []interface{} {
	return ctx.responsePipeline
} /*
 */
