package models

import "github.com/jinzhu/gorm"

type ArtModel struct {
	gorm.Model         //Id is automatically provided by gorm.Model
	Artist_id   int    `json:"artist_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     []byte `json:"picture"`
}
