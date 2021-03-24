package script

import (
	"html/template"
	"path/filepath"
)

const (
	queryTypePath = "queryType"
)

func GenerateGql(gqlPath string, entities []*Entity) {

	// 通过模版来渲染，字符串不好拼接代码
	path, err := filepath.Abs("./script/template/queryType.tpl")
	if err != nil {
		panic(err)
	}
	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"getGraphqlType": getGraphqlType})
	funcMaps = append(funcMaps, template.FuncMap{"proccessFieldName": proccessFieldName})

	for _, entity := range entities {
		result := RenderTemplate("queryType.tpl", path, entity, funcMaps)
		MkdirAll(gqlPath + "/" + queryTypePath)
		err := WriteFile(gqlPath+"/"+queryTypePath+"/"+entity.Name+".go", result)
		if err != nil {
			panic(err)
		}
	}
	// rootQueryType
	GenerateRootGql(gqlPath, entities)
}

func getGraphqlType(field Field) string {

	if field.Type == "string" {
		return "graphql.String"
	} else if CoerceInt(field.Type) == "int" {
		return "graphql.Int"
	} else if CoerceFloat(field.Type) == "float" {
		return "graphql.Float"
	} else if field.Type == "bool" {
		return "graphql.Boolean"
	}

	return ""
}

func getUppercaseName(field Field) string {
	return ProccessFieldName(field.Name)
}

type RootQuery struct {
	Entities []*Entity
}

// 生成rootQueryType
func GenerateRootGql(gqlPath string, entities []*Entity) {
	path, err := filepath.Abs("./script/template/rootQueryType.tpl")
	if err != nil {
		panic(err)
	}

	rt := RootQuery{
		Entities: entities,
	}

	var funcMaps []template.FuncMap
	result := RenderTemplate("rootQueryType.tpl", path, rt, funcMaps)
	MkdirAll(gqlPath + "/" + queryTypePath)
	err = WriteFile(gqlPath+"/"+queryTypePath+"/rootQueryType.go", result)
	if err != nil {
		panic(err)
	}
}
