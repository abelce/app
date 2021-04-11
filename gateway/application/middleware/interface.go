package middleware

import (
	normal_context "abelce/app/gateway/application/context"

	"github.com/gorilla/mux"
)

type Middleware interface {
	Register(ctx *normal_context.Context) mux.MiddlewareFunc
}
