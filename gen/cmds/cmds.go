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

func runMainCommand(entities []*model.Entity) {
	var mainCommand []command.Command
	constantCommand := command.NewConstantCommand(utils.GetRealPath(utils.CodeGenPath), entities)
	modelCommand := command.NewModelCommand(utils.GetRealPath(utils.CodeGenPath), entities)
	gqlCommand := command.NewGqlCommand(utils.GetRealPath(utils.GqlPath), entities)

	mainCommand = append(mainCommand, constantCommand)
	mainCommand = append(mainCommand, modelCommand)
	mainCommand = append(mainCommand, gqlCommand)

	for _, command := range mainCommand {
		command.Execute()
	}
}