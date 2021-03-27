package command

import (
	"fmt"
	"html/template"
	"path/filepath"
	"vwood/app/gen/assets/utils"
	"vwood/app/gen/domain/model"
)

const (
	queryTypePath = "queryType"
)

type GqlCommand struct {
	BasePath string
	Entities []*model.Entity
}

func NewGqlCommand(basePath string, entities []*model.Entity) GqlCommand {
	return GqlCommand{
		BasePath: basePath,
		Entities: entities,
	}
}

func (t GqlCommand) Execute() {
	//for _, entity := range t.Entities {
		fmt.Println("[generate graphql-------------------]")
		GenerateGql(t.BasePath, t.Entities)
	//}
}
func (t GqlCommand) Add(cm Command) {}

func GenerateGql(gqlPath string, entities []*model.Entity) {

	// 通过模版来渲染，字符串不好拼接代码
	path, err := filepath.Abs("./assets/template/queryType.tpl")
	if err != nil {
		panic(err)
	}
	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"getGraphqlType": getGraphqlType})
	funcMaps = append(funcMaps, template.FuncMap{"proccessFieldName": proccessFieldName})

	for _, entity := range entities {
		fmt.Println("[generate graphql-------------------]" + entity.Name)
		result := utils.RenderTemplate("queryType.tpl", path, entity, funcMaps)
		utils.MkdirAll(gqlPath + "/" + queryTypePath)
		err := utils.WriteFile(gqlPath+"/"+queryTypePath+"/"+entity.Name+".go", result)
		if err != nil {
			panic(err)
		}
	}
	// rootQueryType
	GenerateRootGql(gqlPath, entities)
}

func getGraphqlType(field model.Field) string {

	if field.Type == "string" {
		return "graphql.String"
	} else if utils.CoerceInt(field.Type) == "int" {
		return "graphql.Int"
	} else if utils.CoerceFloat(field.Type) == "float" {
		return "graphql.Float"
	} else if field.Type == "bool" {
		return "graphql.Boolean"
	}

	return ""
}

func getUppercaseName(field model.Field) string {
	return utils.ProccessFieldName(field.Name)
}

type RootQuery struct {
	Entities []*model.Entity
}

// 生成rootQueryType
func GenerateRootGql(gqlPath string, entities []*model.Entity) {

	fmt.Println("[generate graphql-------------------] rootQueryType")

	path, err := filepath.Abs("./assets/template/rootQueryType.tpl")
	if err != nil {
		panic(err)
	}

	rt := RootQuery{
		Entities: entities,
	}

	var funcMaps []template.FuncMap
	funcMaps = append(funcMaps, template.FuncMap{"lowerCase": utils.LowerCase})

	result := utils.RenderTemplate("rootQueryType.tpl", path, rt, funcMaps)
	utils.MkdirAll(gqlPath + "/" + queryTypePath)
	err = utils.WriteFile(gqlPath+"/"+queryTypePath+"/rootQueryType.go", result)
	if err != nil {
		panic(err)
	}
}
