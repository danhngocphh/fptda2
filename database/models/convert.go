package models

import (
	// "github.com/jinzhu/gorm"
	"github.com/danhngocphh/fptda2/lib/common"
)

// Post data model
type Convert struct {
	voice   string `sql:"type:text;"`
	text   string `sql:"type:text;"`
	speed   string `sql:"type:number;"`
	format   string `sql:"type:text;"`
}

// Serialize serializes post data
func (c Convert) Serialize() common.JSON {
	return common.JSON{
		"voice":       c.voice,
		"text":       c.text,
		"speed":       c.speed,
		"format":       c.format,
	}
}
