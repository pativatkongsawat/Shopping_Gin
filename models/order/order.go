package order

import "time"

type Order struct {
	Id         int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	CreateAt   *time.Time `json:"create_at" gorm:"create_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" gorm:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" gorm:"deleted_at"`
	UserId     string     `json:"user_id" gorm:"user_id"`
	TotalPrice float64    `json:"total_price" gorm:"total_price"`
	CreatedBy  string     `json:"created_by,omitempty" gorm:"created_by"`
	Status     string     `json:"status" gorm:"status" default:"unpaid"`
}
