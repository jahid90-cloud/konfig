package config

import (
	"errors"
)

type config struct {
	Id          string                 `json:"id"`
	ParentId    string                 `json:"parentId"`
	Description string                 `json:"description"`
	Data        map[string]interface{} `json:"data"`
	Metadata    map[string]interface{} `json:"metadata"`
}

var ErrAlreadyExists = errors.New("already exists")

func NewConfig(id, parentId string) *config {
	return &config{
		Id:          id,
		ParentId:    parentId,
		Description: "",
		Data:        make(map[string]interface{}),
		Metadata:    make(map[string]interface{}),
	}
}

func (c *config) UpdateDescription(desc string) {
	c.Description = desc
}

func (c *config) AddData(key string, value interface{}) error {
	if _, contains := c.Data[key]; contains {
		return ErrAlreadyExists
	}

	c.Data[key] = value

	return nil
}

func (c *config) AddMetadata(key string, value interface{}) error {
	if _, contains := c.Metadata[key]; contains {
		return ErrAlreadyExists
	}

	c.Metadata[key] = value

	return nil
}
