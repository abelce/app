package queryType

import (
	gen_md "vwood/app/common/code-gen/models"
	"github.com/graphql-go/graphql"
)

{{$entityName := .Name}}
var {{$entityName}}Type = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
        {{range $i, $field := .Fields}}
		"{{$field.Name}}": &graphql.Field{
			Type: {{getGraphqlType $field}},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source == nil {
					return nil, nil
				}
				entity := p.Source.(gen_md.{{$entityName}})
                return entity.{{proccessFieldName $field.Name}}, nil
			},
		},
        {{end}}
	},
})