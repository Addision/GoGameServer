package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"utils/file"
)

var (
	path_not_exist = errors.New("config path is not exit")
)

type IConfigure interface {
	Load(configPath string) error
	AssignStruct(st interface{}) error
	GetConfig() interface{}
}

type Configure struct {
	st   interface{}
	data []byte
}

func NewConfigure() *Configure {
	return &Configure{
		st:   nil,
		data: nil,
	}
}

func (c *Configure) Load(configPath string) error {
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "."
	}
	configPath = pwd + "/" + configPath
	if isExists := file.Exists(configPath); !isExists {
		return path_not_exist
	}
	c.data, err = ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	return nil
}

func (c *Configure) AssignStruct(st interface{}) error {
	if c.data == nil {
		return errors.New("call load function first")
	}
	if err := json.Unmarshal(c.data, st); err != nil {
		return err
	}
	c.st = st
	return nil
}

func (c *Configure) GetConfig() interface{} {
	return c.st
}
