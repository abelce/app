package queryType

import (
	gen_md "vwood/app/common/code-gen/models"
	"github.com/graphql-go/graphql"
)


var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
        
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.Product)
                return entity.Id, nil
			},
		},
        
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.Product)
                return entity.Name, nil
			},
		},
        
		"categoryId": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.Product)
                return entity.CategoryId, nil
			},
		},
        
		"isDeleted": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.Product)
                return entity.IsDeleted, nil
			},
		},
        
		"updatedTime": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.Product)
                return entity.UpdatedTime, nil
			},
		},
        
		"createdTime": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.Product)
                return entity.CreatedTime, nil
			},
		},
        
		"operateID": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.Product)
                return entity.OperateID, nil
			},
		},
        
	},
})
