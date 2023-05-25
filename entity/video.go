package entity

import "time"

type Person struct {
	ID        uint64
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

type Song struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title       string    `json:"title" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	URL         string    `json:"url" gorm:"type:varchar(255);unique"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignKey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
