package command

import (
	"app/gen/assets/utils"
	"app/gen/domain/model"
	"fmt"
	"strings"
)

const (
	enumsPath = "enum-names"
)

type EnumCommand struct {
	BasePath string
	Entities []*model.Entity
}

// 只有type == enum的才执行
func NewEnumCommand(basePath string, entities []*model.Entity) EnumCommand {
	return EnumCommand{
		BasePath: basePath,
		Entities: entities,
	}
}

func (t EnumCommand) Execute() {
	for _, entity := range t.Entities {
		fmt.Println("[generate enums-------------------]" + entity.Name)
		GenerateEnum(t.BasePath, entity)
	}
}
func (t EnumCommand) Add(cm Command) {}

func GenerateEnum(codeGenPath string, entity *model.Entity) {
	utils.RemovePath(codeGenPath)
	utils.Mkdir(codeGenPath)
	// 生成fields
	generateEnumFields(codeGenPath, entity)
}

func generateEnumFields(codeGenPath string, entity *model.Entity) {
	entityName := entity.Name

	var result []string

	result = append(result, "package gen_enums")  // package
	result = append(result, "")                // 空行
	result = append(result, "//"+entity.Title) // 实体名称
	result = append(result, "const (")         // 使用常量
	result = append(result, "  ENUM_"+entityName + " = \""+entityName+"\"")
	for _, field := range entity.Fields {
		result = append(result, "  //"+field.Title)
		result = append(result, "  ENUM_"+entityName+"_"+field.Name+" = \""+entityName +"."+field.Name+"\"")
	}
	result = append(result, ")")

	utils.Mkdir(codeGenPath + "/" + enumsPath)
	err := utils.WriteFile(codeGenPath+"/"+enumsPath+"/"+entityName+".go", strings.Join(result, "\n"))
	if err != nil {
		panic(err)
	}

}
