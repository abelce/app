package script

import (
	"strings"
)

const (
	fieldsPath = "interface-fields"
)

func GenerateConstant(codeGenPath string, entity *Entity) {
	RemovePath(codeGenPath)
	Mkdir(codeGenPath)
	// 生成fields
	generateFields(codeGenPath, entity)
}

func generateFields(codeGenPath string, entity *Entity) {
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

	Mkdir(codeGenPath + "/" + fieldsPath)
	err := WriteFile(codeGenPath+"/"+fieldsPath+"/"+entityName+".go", strings.Join(result, "\n"))
	if err != nil {
		panic(err)
	}

}
