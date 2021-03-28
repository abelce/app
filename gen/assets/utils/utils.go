package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"vwood/app/gen/domain/model"
)

const (
	EntityPath  = "../entity"
	CodeGenPath = "../common/code-gen"
	ModelPath   = "../common/model"
	GqlPath     = "../graphql/application"
	DatabasePath = "../common/database"
)

// 删除制定目录下的所有文件
func RemovePath(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

// 文件是否存在
func FileIsExist(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

// 创建制定目录
func Mkdir(path string) error {
	return os.Mkdir(path, 0777)
}

// 创建练级目录
func MkdirAll(path string) error {
	return os.MkdirAll(path, 0777)
}

// 写入文件
func WriteFile(path string, data string) error {
	fmt.Println("write file...")
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()
	fmt.Fprintln(f, data)

	return nil
}

func IsRequiredFields(field string) bool {
	requiredFields := [...]string{"id", "isDeleted", "createdTime", "updatedTime"}

	for _, item := range requiredFields {
		return item == field
	}

	return false
}

func IsIncludeItem(items []string, target string) bool {
	for _, item := range items {
		return item == target
	}

	return false
}

func CoerceInt(fieldType string) string {
	switch fieldType {
	case "int":
		return "int"
	case "int8":
		return "int"
	case "int16":
		return "int"
	case "int32":
		return "int"
	case "int64":
		return "int"
	case "uint":
		return "int"
	case "uint8":
		return "int"
	case "uint16":
		return "int"
	case "uint32":
		return "int"
	case "uint64":
		return "int"
	}

	return fieldType
}

func CoerceFloat(fieldType string) string {
	switch fieldType {
	case "float32":
		return "float"
	case "float64":
		return "float"
	}

	return fieldType
}

func ProccessFieldName(name string) string {
	if name == "" {
		panic("name is can not nil")
	}
	capture := string([]byte(name)[:1]) // 这里不考虑中文的空间不止一个字节的问题， name通常都是中文
	others := string([]byte(name)[1:])
	return strings.ToUpper(capture) + others
}

func RenderTemplate(templateName string, tplPath string, data interface{}, funcMaps []template.FuncMap) string {
	t := template.New(templateName)
	for _, funcMap := range funcMaps {
		t.Funcs(funcMap)
	}

	file, err := filepath.Abs(tplPath)
	if err != nil {
		panic(err)
	}
	t, err = t.ParseFiles(file)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func LowerCase(name string) string {
	if name == "" {
		panic("name is can not nil")
	}
	capture := string([]byte(name)[:1]) // 这里不考虑中文的空间不止一个字节的问题， name通常都是中文
	others := string([]byte(name)[1:])
	return strings.ToLower(capture) + others
}

func getDirPath() (string, error) {
	path, err := os.Getwd()
	return path + "/", err
}

// 从执行目录获取文件的实际路径
func GetRealPath(path string) string {
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
func ReadJsonFiles(dir string) []string {
	rd, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var fileNames []string
	for _, fi := range rd {
		if !fi.IsDir() && strings.HasSuffix(fi.Name(), ".json") {
			fileNames = append(fileNames, fi.Name())
		}

	}

	return fileNames
}

// 读取一个文件的内容
func ReadOneJsonFile(path string) *model.Entity {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	entity := &model.Entity{}

	err = json.Unmarshal(data, entity)
	if err != nil {
		panic(err)
	}

	// 判断必须的字段有没有， 没有自动加上

	return entity
}
