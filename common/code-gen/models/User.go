package gen_md

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// 用户
type User struct {
	  //用户ID
  Id string `json:"id" valid:"required"`
  //姓名
  Name string `json:"name" valid:"required"`
  //性别
  Sex string `json:"sex" valid:"required"`
  //是否删除
  IsDeleted bool `json:"isDeleted"`
  //更新时间
  UpdatedTime int64 `json:"updatedTime"`
  //创建时间
  CreatedTime int64 `json:"createdTime"`
}

func (entity *User) Valid() error {
	_, err := govalidator.ValidateStruct(entity)
	if err != nil {
		return err
	}

	return nil
}

func NewUser(
	  id string,
  name string,
  sex string,
  updatedTime int64,
  createdTime int64,
	) (*User, error) {
	entity := &User {
		   Id: id,
  Name: name,
  Sex: sex,
  UpdatedTime: updatedTime,
  CreatedTime: createdTime, 
		}

    entity.IsDeleted = false
    entity.CreatedTime = time.Now().Unix()
	entity.UpdatedTime = time.Now().Unix()
	
    if err := entity.Valid(); err != nil {
		return nil, err
	}

	return entity, nil
}

func (entity *User) Delete() {
	entity.IsDeleted = false
	entity.UpdatedTime = time.Now().Unix()
}

func (entity *User) Update(
	  name string,
  sex string,
  isDeleted bool,
  updatedTime int64,
  createdTime int64,
) error {
	  entity.Name=name
  entity.Sex=sex
  entity.IsDeleted=isDeleted
  entity.UpdatedTime=updatedTime
  entity.CreatedTime=createdTime

	entity.UpdatedTime = time.Now().Unix()

	return entity.Valid()
}

