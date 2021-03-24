package application

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
)

type Context struct {
	config   *Config
	pgClient *sql.DB

	service *Service
	// productRepository   model.ProductRepository
	// productQueryService *ProductQueryService
}

func NewContext(cfgPath string) (*Context, error) {
	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	c := &Context{}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	c.config = cfg

	return c, nil
}

func (c *Context) GetConfig() *Config {
	return c.config
}

func (c *Context) PgClient() (*sql.DB, error) {
	if c.pgClient != nil {
		return c.pgClient, nil
	}

	client, err := sql.Open(c.config.DB.DriverName, c.config.DB.DataSourceName)
	if err != nil {
		return nil, err
	}

	c.pgClient = client

	return c.pgClient, err
}

// func (c *Context) Service() (*Service, error) {
// 	if c.service != nil {
// 		return c.service, nil
// 	}
// 	pr, err := c.ProductRepository()
// 	if err != nil {
// 		return nil, err
// 	}
// 	service := NewService(pr)
// 	c.service = service

// 	return service, nil
// }

// func (c *Context) ProductRepository() (model.ProductRepository, error) {
// 	if c.productRepository != nil {
// 		return c.productRepository, nil
// 	}

// 	client, err := c.PgClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	ar := &pgsql.ProductRepository{
// 		Client:    client,
// 		TableName: c.config.ProductTableName,
// 	}

// 	c.productRepository = ar
// 	return c.productRepository, nil
// }

// func (c *Context) ProductQueryService() (*ProductQueryService, error) {
// 	if c.productQueryService != nil {
// 		return c.productQueryService, nil
// 	}
// 	client, err := c.PgClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	pr := NewProductQueryService(client, c.config.ProductTableName)

// 	c.productQueryService = pr
// 	return c.productQueryService, nil
// }
