package schema

import (
	"log"
	"vwood/app/graphql/application"
	"vwood/app/graphql/application/queryType"

	"github.com/graphql-go/graphql"
)

// 初始化schema
func GetSchema() graphql.Schema {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType.GetRootQueryType(application.ApplicationContext.GatewayEndpoint()),
		// Mutation: MutationType,
		Extensions: GetExtensions(),
	})
	if err != nil {
		log.Fatal(err)
	}
	return schema
}

func GetExtensions() []graphql.Extension {
	var extensions []graphql.Extension
	extensions = append(extensions, SchemaExtension{})

	return extensions
}
