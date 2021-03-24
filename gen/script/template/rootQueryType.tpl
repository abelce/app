package queryType

import (
	"github.com/graphql-go/graphql"
)

var RootQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
        {{range $i, $entity := .Entities}}
		"{{$entity.Name}}": &graphql.Field{
			Type: {{$entity.Name}}Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 请求转发到具体的服务，并获取数据
				return nil, nil
			},
		},
        {{end}}
	},
})