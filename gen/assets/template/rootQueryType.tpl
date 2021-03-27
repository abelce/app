package queryType

import (
	gen_md "vwood/app/common/code-gen/models"
	"vwood/app/common/request"

	"github.com/graphql-go/graphql"
	"encoding/json"
)

func GetRootQueryType(endpoint string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			{{range $i, $entity := .Entities}}
			"{{$entity.Name}}": &graphql.Field{
				Type: {{$entity.Name}}Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 请求转发到具体的服务，并获取数据
					if id, ok := p.Args["id"].(string); ok && id != "" {
						req := request.Request{
							Url: endpoint + "/v1/{{lowerCase $entity.Name}}s/" + id,
							Method: "GET",
						}
						result, err := req.Do()
						if err != nil {
							return nil, err
						}
						var entity gen_md.Product
						err = json.Unmarshal(result, &entity)
						if err != nil {
							return nil, nil
						}
						
						return entity, nil	
					}

					return nil, nil
				},
			},
			"{{$entity.Name}}-list": &graphql.Field{
				Type: graphql.NewList({{$entity.Name}}Type),
				Args: graphql.FieldConfigArgument{
					"queryStr": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 请求转发到具体的服务，并获取数据
					// 考虑将服务器上返回的其他信息写入到context上，在ExecutionDidStart中写到exensions结构中，最后在graphql执行完后重新组装起来
					if queryStr, ok := p.Args["queryStr"].(string); ok && queryStr != "" {
						req := request.Request{
							Url: endpoint + "/v1/{{lowerCase $entity.Name}}s/list",
							Method: "GET",
							Params: queryStr,
						}
						result, err := req.Do()
						if err != nil {
							return nil, err
						}
						var entity gen_md.Product
						json.Unmarshal(result, &entity)
						
						return entity, nil	
					}

					return nil, nil
				},
			},
			{{end}}
		},
	})
}
