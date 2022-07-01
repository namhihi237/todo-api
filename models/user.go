package models

import "time"

type User struct {
	Id        int        `json:"id" gorm:"primary_key;auto_increment;not null;unique;index;"`
	Name      string     `json:"name" gorm:"size:255; not null;"`
	Email     string     `json:"email" gorm:"size:255; not null;"`
	Password  string     `json:"password" gorm:"size:255; not null;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"default:null;"`
}
