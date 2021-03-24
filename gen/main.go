package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"vwood/app/gen/script"
)

func main() {
	run()
}

func getDirPath() (string, error) {
	path, err := os.Getwd()
	return path + "/", err
}

// 从执行目录获取文件的实际路径
func getRealPath(path string) string {
	dirPath, err := getDirPath()
	if err != nil {
		panic(err)
	}
	realPath, err := filepath.Abs(dirPath + path)
	if err != nil {
		panic(err)
	}
	return realPath
}

// 读取文件名称
func readJsonFiles(path string) []string {
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var fileNames []string
	for _, fi := range rd {
		if !fi.IsDir() {
			fileNames = append(fileNames, fi.Name())
		}
	}

	return fileNames
}

// 读取一个文件的内容
func readOneJsonFile(path string) *script.Entity {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	entity := &script.Entity{}

	err = json.Unmarshal(data, entity)
	if err != nil {
		panic(err)
	}

	// 判断必须的字段有没有， 没有自动加上

	return entity
}

func run() {
	// 读取entity中的json文件
	fileNames := readJsonFiles(getRealPath(script.EntityPath))
	// 存储所有的entity， 方便后面需要所有的entity一起才能处理的任务使用
	var entites []*script.Entity

	for _, fileName := range fileNames {
		entity := readOneJsonFile(getRealPath(script.EntityPath + "/" + fileName))
		entites = append(entites, entity)
		// 生成常量
		fmt.Println("[generate constants-------------------]")
		script.GenerateConstant(getRealPath(script.CodeGenPath), entity)
		// 生成 model
		fmt.Println("[generate models-------------------]")
		script.GenerateModel(getRealPath(script.CodeGenPath), entity)
		// 生成gql代码
		// script.GenerateGql(entity)
	}
	fmt.Println("[generate gqls-------------------]")
	script.GenerateGql(getRealPath(script.GqlPath), entites)
}
