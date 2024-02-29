package classification

import (
	"encoding/json"
	"errors"
)

type Classification struct {
	ClassId int    `json:"classId" db:"class_id"`
	Name    string `json:"name" db:"name"`
	Enabled bool   `json:"enabled" db:"enabled"`
}

func (c *Classification) Scan(v any) error {
	var bytes, ok = v.([]byte)

	if !ok {
		return errors.New("invalid data type")
	}

	return json.Unmarshal(bytes, &c)
}

type Classifications []*Classification

func (c *Classifications) Scan(v any) error {
	var bytes, ok = v.([]byte)

	if !ok {
		return errors.New("invalid data type")
	}

	return json.Unmarshal(bytes, &c)
}
