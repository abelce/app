package middleware

import (
	normal_context "abelce/app/gateway/application/context"

	"github.com/gorilla/mux"
)

// 这里所有的中间键暂时只考虑加到请求发生之前
func GetMiddlewares(ctx *normal_context.Context) (fns []mux.MiddlewareFunc) {
	var mdws []Middleware
	if ctx.GetPort() == 443 {
		mid := NewAuthMiddleware()
		mdws = append(mdws, mid)
	}

	for _, mid := range mdws {
		fns = append(fns, mid.Register(ctx))
	}

	return fns
}
