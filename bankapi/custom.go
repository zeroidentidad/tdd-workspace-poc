package main

import (
	"bankcore"
	"encoding/json"
)

// CustomAccount ...
type CustomAccount struct {
	*bankcore.Account
}

// Statement ...
func (c *CustomAccount) Statement() string {
	json, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}

	return string(json)
}
