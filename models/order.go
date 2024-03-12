package models

import (
	"time"
)

type Order struct {
	ID            uint      `gorm:"primaryKey"`
	Customer_name string    `gorm:"type:varchar(100)"`
	Order_date    time.Time `gorm:"type:date;timestamptz;not null;default:CURRENT_TIMESTAMP"`
	Items         []Item    `gorm:"foreignKey:Order_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
