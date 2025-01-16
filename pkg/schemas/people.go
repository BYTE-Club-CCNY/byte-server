package schema

import (
	"github.com/google/uuid"
)

type User struct {
    UID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
    FirstName      string    `gorm:"type:varchar(255);not null"`
    LastName       string    `gorm:"type:varchar(255);not null"`
    CunyEmail      string    `gorm:"type:varchar(255);not null"`
    Emplid         string    `gorm:"type:varchar(255);not null"`
    Active         bool      `gorm:"column=active;default:false"`
    MiddleName     string    `gorm:"type:varchar(255);default:null"`
    PersonalEmail  string    `gorm:"type:varchar(255);default:null"`
    Discord        string    `gorm:"type:varchar(255);default:null"`
}

func (User) TableName() string {
    return "users.people"
}
