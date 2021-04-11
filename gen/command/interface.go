package command

import "app/gen/domain/model"

type Command interface {
	Add(Command)
	Execute()
}

type GenBase struct {
	BasePath string
	Entities []*model.Entity
}
