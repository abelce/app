package queryType

import (
	gen_md "vwood/app/common/code-gen/models"
	"vwood/app/common/request"

	"encoding/json"

	"github.com/graphql-go/graphql"
)

func GetRootQueryType(endpoint string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{

			"Product": &graphql.Field{
				Type: ProductType,
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
							Url:    endpoint + "/v1/products/" + id,
							Method: "GET",
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

			"User": &graphql.Field{
				Type: UserType,
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
							Url:    endpoint + "/v1/users/" + id,
							Method: "GET",
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
		},
	})
}
