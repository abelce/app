package application

import (
	"fmt"
	"net/url"
	"regexp"
	"vwood/app/gateway/application/middleware"

	"github.com/gorilla/mux"
)

var (
	ErrAPINotExist = at.NewError("gateway:application/ErrAPINotExist", "api not exist")
)

type Context struct {
	config     *Config
	apiProxies map[string]*url.URL
	middlewares []middleware.Middleware
}

func NewContext(cfg *Config) *Context {
	ctx := &Context{
		config: cfg,
	}

	ctx.apiProxies = make(map[string]*url.URL)
	for _, aApi := range cfg.APIs {
		ctx.apiProxies[aApi.Key] = &url.URL{
			Scheme: aApi.Scheme,
			Host:   aApi.Host,
			Path:   aApi.Path,
		}
	}
	return ctx
}

func (c *Context) Config() *Config {
	return c.config
}

func (c *Context) APIProxyURL(key string) (*url.URL, error) {
	if c.apiProxies[key] == nil {
		return nil, ErrAPINotExist
	}

	return c.apiProxies[key], nil
}

//key: 服务, url: url， action: 方法(get, post, delete, put)
func (c *Context) NoVerify(key, url, action string) bool {
	fmt.Println("#v", c.Config().NoVerify)
	verifyMethod := c.Config().NoVerify[key]
	if verifyMethod == nil {
		return false
	}

	// all表示所以接口都不验证
	if verifyMethod.All == true {
		return true
	}

	method := verifyMethod.Methods[action]
	if method == nil {
		return false
	}

	if method.All == true {
		return true
	}

	if len(method.Urls) > 0 {
		for _, v := range method.Urls {
			if v == url {
				return true
			}
		}
	}

	if len(method.Regs) > 0 {
		for _, v := range method.Regs {
			ok, err := regexp.MatchString(v, url)
			if at.Ensure(&err) {
				return false
			}
			if ok {
				return true
			}
		}
	}

	return false
}

func (c *Context) Router80() *mux.Router {
	return NewRouter(80)
}

func (c *Context) Router443() *mux.Router {
	return NewRouter(443)
}
