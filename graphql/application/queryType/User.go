package queryType

import (
	gen_md "vwood/app/common/code-gen/models"
	"github.com/graphql-go/graphql"
)


var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
        
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.User)
                return entity.Id, nil
			},
		},
        
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.User)
                return entity.Name, nil
			},
		},
        
		"sex": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.User)
                return entity.Sex, nil
			},
		},
        
		"isDeleted": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.User)
                return entity.IsDeleted, nil
			},
		},
        
		"updatedTime": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.User)
                return entity.UpdatedTime, nil
			},
		},
        
		"createdTime": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.User)
                return entity.CreatedTime, nil
			},
		},
        
	},
})
