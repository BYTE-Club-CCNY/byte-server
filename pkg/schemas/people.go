package schema

import (
	"github.com/google/uuid"
)

type User struct {
    UID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
    Active         bool      `gorm:"default:true"`
    FirstName      string    `gorm:"type:varchar(255);not null"`
    MiddleName     string    `gorm:"type:varchar(255)"`
    LastName       string    `gorm:"type:varchar(255);not null"`
    PersonalEmail  string    `gorm:"type:varchar(255)"`
    CunyEmail      string    `gorm:"type:varchar(255)"`
    Discord        string    `gorm:"type:varchar(255)"`
    Emplid         string    `gorm:"type:varchar(255);not null"`
}

func (User) TableName() string {
    return "users.people"
}
