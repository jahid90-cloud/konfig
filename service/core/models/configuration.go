package models

import (
	"errors"
)

type configuration struct {
	Id          string
	ParentId    string
	Description string
	Data        map[string]interface{}
	Metadata    map[string]interface{}
}

var ErrConfgurationAlreadyExists = errors.New("configuration already exists")

func NewConfiguration(id, parentId string) *configuration {
	return &configuration{
		Id:          id,
		ParentId:    parentId,
		Description: "",
		Data:        make(map[string]interface{}),
		Metadata:    make(map[string]interface{}),
	}
}

func (c *configuration) UpdateDescription(desc string) {
	c.Description = desc
}

func (c *configuration) AddData(key string, value interface{}) error {
	if _, contains := c.Data[key]; contains {
		return ErrConfgurationAlreadyExists
	}

	c.Data[key] = value

	return nil
}

func (c *configuration) AddMetadata(key string, value interface{}) error {
	if _, contains := c.Metadata[key]; contains {
		return ErrConfgurationAlreadyExists
	}

	c.Metadata[key] = value

	return nil
}
