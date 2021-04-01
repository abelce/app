package command

import (
	"vwood/app/gen/assets/utils"
	"vwood/app/gen/domain/model"
)

type MainCommand struct {
	Entities []*model.Entity
	CommandList []Command
}

func NewMainCommand(entities []*model.Entity) MainCommand {
	return MainCommand{
		Entities: entities,
	}
}

func (t MainCommand) Execute() {
	constantCommand := NewConstantCommand(utils.GetRealPath(utils.CodeGenPath), t.Entities)
	modelCommand := NewModelCommand(utils.GetRealPath(utils.CodeGenPath), t.Entities)
	gqlCommand := NewGqlCommand(utils.GetRealPath(utils.GqlPath), t.Entities)
	databaseCommand := NewDatabaseCommand(utils.GetRealPath(utils.DatabasePath), t.Entities)

	t.CommandList = append(t.CommandList, constantCommand)
	t.CommandList = append(t.CommandList, modelCommand)
	t.CommandList = append(t.CommandList, gqlCommand)
	// databaseCommand 该命令中的建库脚本、docker.sh脚本、数据库的docker配置可以考虑使用子命令来组合（组合模式）, 暂时在一个脚本集中处理
	t.CommandList = append(t.CommandList, databaseCommand)

	for _, cmd := range t.CommandList {
		cmd.Execute()
	}
}
func (t MainCommand) Add(cm Command) {}
