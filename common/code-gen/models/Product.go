package gen_md

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// 产品
type Product struct {
	  //产品ID
  Id string `json:"id" valid:"required"`
  //产品名称
  Name string `json:"name" valid:"required"`
  //产品类型
  CategoryId string `json:"categoryId" valid:"required"`
  //是否删除
  IsDeleted bool `json:"isDeleted"`
  //更新时间
  UpdatedTime int64 `json:"updatedTime"`
  //创建时间
  CreatedTime int64 `json:"createdTime"`
  //用户ID
  OperateID string `json:"operateID"`
}

func (entity *Product) Valid() error {
	_, err := govalidator.ValidateStruct(entity)
	if err != nil {
		return err
	}

	return nil
}

func NewProduct(
	  id string,
  name string,
  categoryId string,
  updatedTime int64,
  createdTime int64,
  operateID string,
	) (*Product, error) {
	entity := &Product {
		   Id: id,
  Name: name,
  CategoryId: categoryId,
  UpdatedTime: updatedTime,
  CreatedTime: createdTime,
  OperateID: operateID, 
		}

    entity.IsDeleted = true
    entity.CreatedTime = time.Now().Unix()
	entity.UpdatedTime = time.Now().Unix()
	
    if err := entity.Valid(); err != nil {
		return nil, err
	}

	return entity, nil
}

func (entity *Product) Delete() {
	entity.IsDeleted = true
	entity.UpdatedTime = time.Now().Unix()
}

func (entity *Product) Update(
	  name string,
  categoryId string,
  isDeleted bool,
  updatedTime int64,
  createdTime int64,
  operateID string,
) error {
	  entity.Name=name
  entity.CategoryId=categoryId
  entity.IsDeleted=isDeleted
  entity.UpdatedTime=updatedTime
  entity.CreatedTime=createdTime
  entity.OperateID=operateID

	entity.UpdatedTime = time.Now().Unix()

	return entity.Valid()
}

