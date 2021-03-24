package schema

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

type SchemaExtension struct {
}

func (t *SchemaExtension) Init(ctx context.Context, p *graphql.Params) context.Context {
	return ctx
}

func (t *SchemaExtension) Name() string {
	return "SchemaExtension"
}

func (t *SchemaExtension) HasResult() bool {
	return false
}

func (t *SchemaExtension) GetResult(ctx context.Context) interface{} {
	return nil
}

func (t *SchemaExtension) ParseDidStart(ctx context.Context) (context.Context, graphql.ParseFinishFunc) {
	return ctx, func(err error) {}
}

func (t *SchemaExtension) ValidationDidStart(ctx context.Context) (context.Context, graphql.ValidationFinishFunc) {
	return ctx, func(errs []gqlerrors.FormattedError) {}
}

func (t *SchemaExtension) ExecutionDidStart(ctx context.Context) (context.Context, graphql.ExecutionFinishFunc) {
	return ctx, func(*graphql.Result) {}
}

func (t *SchemaExtension) ResolveFieldDidStart(ctx context.Context, i *graphql.ResolveInfo) (context.Context, graphql.ResolveFieldFinishFunc) {
	return ctx, func(v interface{}, err error) {}
}
