package dtos

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model

    ID          int64     	`gorm:"primary_key" json:"id"`
    Category    string    	`gorm:"column:category" json:"category"`
    Type        int       	`gorm:"column:type" json:"type"`
    Description string    	`gorm:"column:description" json:"description"`
    Amount      float64   	 `gorm:"column:amount" json:"amount"`
    Date        time.Time   `gorm:"column:date" json:"date"`
    CreatedAt   time.Time   `gorm:"column:created_at" json:"createdAt"`
    DeletedAt   gorm.DeletedAt  `gorm:"column:deleted_at" json:"deletedAt"`
}

// TableName specifies the table name for the Expense entity
func (Expense) TableName() string {
    return "expenses"
}