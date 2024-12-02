package models

import (
	"fmt"
	"net/http"
)

type CatRequest struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Breed string `json:"breed"`
}

func (c *CatRequest) Bind(r *http.Request) error {
	if c.Name == "" || c.Age <= 0 || c.Breed == "" {
		return fmt.Errorf("invalid data")
	}
	return nil
}

type CatResponse struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Breed  string  `json:"breed"`
	Weight float64 `json:"weight,omitempty"`
}
