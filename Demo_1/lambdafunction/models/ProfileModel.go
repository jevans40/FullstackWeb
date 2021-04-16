package models

import "github.com/jinzhu/gorm"

type ProfileModel struct {
	gorm.Model         //Id is provided automatically by gorm.Model
	Artist_name string `json:"artist_name"`
	Likes       int    `json:"likes"`
	Picture     []byte `json:"picture"`
	Description string `json:"description"`
}
