package models

import "time"

type TodoItem struct {
	Id        int        `json:"id" gorm:"primary_key;auto_increment;not null;unique;index;"`
	Title     string     `json:"title" gorm:"size:255; not null;"`
	Status    string     `json:"status" gorm:"size:30; not null; default:'doing';"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"` // why using *time.Time?
	UpdatedAt *time.Time `json:"updated_at"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"default:null;"`
}
