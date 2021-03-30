package application

import (
	"fmt"
	"log"
	"os"
)

type Rule struct {
	All  bool
	Urls []string
	Regs []string
}

type VerifyMethod struct {
	All     bool
	Methods map[string]*Rule
}

type APIContext struct {
	Key    string
	Scheme string
	Host   string
	Path   string
}

type Config struct {
	SSLChainCrtPath string
	SSLKeyPath      string
	APIs            []*APIContext
	NoVerify        map[string]*VerifyMethod
}

func (c *Config) GetAPI(key string) *APIContext {
	for _, api := range c.APIs {
		if api.Key == key {
			return api
		}
	}
	return nil
}

const (
	DEVELOPMENT string = "development"
	PRODUCTTION string = "production"
	TESTING     string = "testing"
)

func (c *Config) GetENV() string {
	envs := []string{DEVELOPMENT, PRODUCTTION, TESTING}
	env := os.Getenv("ENV")
	for _, val := range envs {
		if env == val {
			return val
		}
	}
	log.Print(fmt.Sprintf("unsuport ENV %s", env))
	return ""
}

func (c *Config) IsDevepoment() bool {
	if c.GetENV() == DEVELOPMENT {
		return true
	}
	return false
}

func (c *Config) IsProduction() bool {
	if c.GetENV() == PRODUCTTION {
		return true
	}
	return false
}

func (c *Config) IsTesting() bool {
	if c.GetENV() == TESTING {
		return true
	}
	return false
}
