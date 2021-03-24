package gen_md

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// {{.Title}}
type {{.Name}} struct {
	{{getStructBody .}}
}

func (entity *{{.Name}}) Valid() error {
	_, err := govalidator.ValidateStruct(entity)
	if err != nil {
		return err
	}

	return nil
}

func New{{.Name}}(
	{{getCreateFuncParams .}}
	) (*{{.Name}}, error) {
	entity := &{{.Name}} {
		 {{getCreateFuncBody .}} 
		}

    entity.IsDeleted = true
    entity.CreatedTime = time.Now().Unix()
	entity.UpdatedTime = time.Now().Unix()
	
    if err := entity.Valid(); err != nil {
		return nil, err
	}

	return entity, nil
}

func (entity *{{.Name}}) Delete() {
	entity.IsDeleted = true
	entity.UpdatedTime = time.Now().Unix()
}

func (entity *{{.Name}}) Update(
	{{getUpdateParams .}}
) error {
	{{getUpdateBody .}}

	entity.UpdatedTime = time.Now().Unix()

	return entity.Valid()
}
