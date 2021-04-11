package application

import (
	"abelce/app/gateway/domain"
	"abelce/at"
	"net/url"

	"github.com/gorilla/mux"
)

var (
	ErrAPINotExist = at.NewError("gateway:application/ErrAPINotExist", "api not exist")
)

type Context struct {
	config *domain.Config // 系统配置文件
	port int
	apiProxies map[string]*url.URL // 考虑用service的方式接入
	// service
}

func NewContext(cfg *domain.Config) *Context {
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

func (c *Context) Config() *domain.Config {
	return c.config
}

func (c *Context) APIProxyURL(key string) (*url.URL, error) {
	if c.apiProxies[key] == nil {
		return nil, ErrAPINotExist
	}

	return c.apiProxies[key], nil
}

func (c *Context) Router80() *mux.Router {
	c.port = 80
	return NewRouter(80)
}

func (c *Context) Router443() *mux.Router {
	c.port = 443
	return NewRouter(443)
}
