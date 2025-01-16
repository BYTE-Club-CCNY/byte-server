package schema

import (
	"github.com/google/uuid"
)

type Team struct {
	ID      uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Member1 uuid.UUID `gorm:"column:member1;type:uuid;not null"`
	Member2 uuid.UUID `gorm:"column:member2;type:uuid"`
	Member3 uuid.UUID `gorm:"column:member3;type:uuid"`
	Member4 uuid.UUID `gorm:"column:member4;type:uuid"`
}

func (Team) TableName() string {
	return "projects.teams"
}
