package cmds

import (
	"vwood/app/gen/assets/utils"
	"vwood/app/gen/command"
	"vwood/app/gen/domain/model"
)

func Run() {
	// 读取entity中的json文件
	fileNames := utils.ReadJsonFiles(utils.GetRealPath(utils.EntityPath))
	// 存储所有的entity， 方便后面需要所有的entity一起才能处理的任务使用
	var entites []*model.Entity

	for _, fileName := range fileNames {
		entity := utils.ReadOneJsonFile(utils.GetRealPath(utils.EntityPath + "/" + fileName))
		entites = append(entites, entity)
	}

	runMainCommand(entites)
}

// 统一执行所有的命令
func runMainCommand(entities []*model.Entity) {
	var mainCommand []command.Command
	constantCommand := command.NewConstantCommand(utils.GetRealPath(utils.CodeGenPath), entities)
	modelCommand := command.NewModelCommand(utils.GetRealPath(utils.CodeGenPath), entities)
	gqlCommand := command.NewGqlCommand(utils.GetRealPath(utils.GqlPath), entities)
	databaseCommand := command.NewDatabaseCommand(utils.GetRealPath(utils.DatabasePath), entities)

	mainCommand = append(mainCommand, constantCommand)
	mainCommand = append(mainCommand, modelCommand)
	mainCommand = append(mainCommand, gqlCommand)
	// databaseCommand 该命令中的建库脚本、docker.sh脚本、数据库的docker配置可以考虑使用子命令来组合（组合模式）, 暂时在一个脚本集中处理
	mainCommand = append(mainCommand, databaseCommand)

	for _, command := range mainCommand {
		command.Execute()
	}
}