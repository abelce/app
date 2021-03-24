package queryType

import (
	"github.com/graphql-go/graphql"
)

var RootQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
        
		"Product": &graphql.Field{
			Type: ProductType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 请求转发到具体的服务，并获取数据
				return nil, nil
			},
		},
        
		"User": &graphql.Field{
			Type: UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 请求转发到具体的服务，并获取数据
				return nil, nil
			},
		},
        
	},
})
