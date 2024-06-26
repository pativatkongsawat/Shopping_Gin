package users

import "time"

type Users struct {
	ID           string     `gorm:"primaryKey" json:"id"`
	Firstname    string     `gorm:"not null" json:"firstname"`
	Lastname     string     `gorm:"not null" json:"lastname"`
	Address      string     `json:"address"`
	Email        string     `gorm:"unique" json:"email"`
	Password     string     `gorm:"not null" json:"password"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	PermissionID int        `json:"permission_id"`
	UpdatedBy    string     `json:"updated_by"`
}

func (Users) TableName() string {
	return "users"
}
