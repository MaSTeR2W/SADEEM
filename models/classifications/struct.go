package classifications

import (
	"encoding/json"
	"errors"
)

type Classification struct {
	ClassId int    `json:"class_id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

func (c *Classification) Scan(v any) error {
	var bytes, ok = v.([]byte)

	if !ok {
		return errors.New("invalid data type")
	}

	return json.Unmarshal(bytes, &c)
}
