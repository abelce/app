package command

import (
	"abelce/app/gen/assets/utils"
	"abelce/app/gen/domain/model"
	"fmt"
)

type MainCommand struct {
	Entities []*model.Entity
	//CommandList []Command
}

func NewMainCommand(entities []*model.Entity) MainCommand {
	return MainCommand{
		Entities: entities,
	}
}

func (t MainCommand) Execute() {
	t.ExecuteRecord()
	t.ExecuteEnums()
}
func (t MainCommand) Add(cm Command) {}

// 实体
func (t MainCommand) ExecuteRecord() {
	var entities []*model.Entity
	for _, entity := range t.Entities {
		if entity.Type == "record" {
			entities = append(entities, entity)
		}
	}

	constantCommand := NewConstantCommand(utils.GetRealPath(utils.CodeGenPath), entities)
	modelCommand := NewModelCommand(utils.GetRealPath(utils.CodeGenPath), entities)
	gqlCommand := NewGqlCommand(utils.GetRealPath(utils.GqlPath), entities)
	databaseCommand := NewDatabaseCommand(utils.GetRealPath(utils.DatabasePath), entities)

	var CommandList []Command
	CommandList = append(CommandList, constantCommand)
	CommandList = append(CommandList, modelCommand)
	CommandList = append(CommandList, gqlCommand)
	// databaseCommand 该命令中的建库脚本、docker.sh脚本、数据库的docker配置可以考虑使用子命令来组合（组合模式）, 暂时在一个脚本集中处理
	CommandList = append(CommandList, databaseCommand)

	for _, cmd := range CommandList {
		cmd.Execute()
	}
}

// 枚举
func (t MainCommand) ExecuteEnums() {
	var entities []*model.Entity
	for _, entity := range t.Entities {
		if entity.Type == "enum" {
			entities = append(entities, entity)
		}
	}
	fmt.Println(len(entities))

	enumCommand := NewEnumCommand(utils.GetRealPath(utils.CodeGenPath), entities)

	var CommandList []Command
	CommandList = append(CommandList, enumCommand)

	for _, cmd := range CommandList {
		cmd.Execute()
	}
}
