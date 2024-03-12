package models

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	Item_code   string `gorm:"type:varchar(100); not null"`
	Description string `gorm:"type:varchar(100)"`
	Quantity    int    `gorm:"type:integer"`
	Order_id    uint
}
