package command

import (
	"app/gen/assets/utils"
	"app/gen/domain/model"
	"fmt"
	"strings"
)

const (
	fieldsPath = "interface-fields"
)

type ConstantCommand struct {
	BasePath string
	Entities []*model.Entity
}

func NewConstantCommand(basePath string, entities []*model.Entity) ConstantCommand {
	return ConstantCommand{
		BasePath: basePath,
		Entities: entities,
	}
}

func (t ConstantCommand) Execute() {
	for _, entity := range t.Entities {
		fmt.Println("[generate constants-------------------]" + entity.Name)
		GenerateConstant(t.BasePath, entity)
	}
}
func (t ConstantCommand) Add(cm Command) {}

func GenerateConstant(codeGenPath string, entity *model.Entity) {
	utils.RemovePath(codeGenPath)
	utils.Mkdir(codeGenPath)
	// 生成fields
	generateFields(codeGenPath, entity)
}

func generateFields(codeGenPath string, entity *model.Entity) {
	entityName := entity.Name

	var result []string

	result = append(result, "package gen_ef")  // package
	result = append(result, "")                // 空行
	result = append(result, "//"+entity.Title) // 实体名称
	result = append(result, "const (")         // 使用常量
	for _, field := range entity.Fields {
		result = append(result, "  //"+field.Title)
		result = append(result, "  F_"+entityName+"_"+field.Name+" = \""+field.Name+"\"")
	}
	result = append(result, ")")

	utils.Mkdir(codeGenPath + "/" + fieldsPath)
	err := utils.WriteFile(codeGenPath+"/"+fieldsPath+"/"+entityName+".go", strings.Join(result, "\n"))
	if err != nil {
		panic(err)
	}

}
