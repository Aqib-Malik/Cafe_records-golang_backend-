package main

import "gorm.io/gorm"

type Cafe struct {
	gorm.Model
	Name     string `json:"name"`
	Location string `json:"location"`
	Type     string `json:"type"`
}
