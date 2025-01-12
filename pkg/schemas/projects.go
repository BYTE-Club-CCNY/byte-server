package schema

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Project struct {
	ID         uuid.UUID      `gorm:"column:id;type:uuid;primaryKey"`
	Name       string         `gorm:"column:name;type:varchar(255)"`
	ShortDesc  string         `gorm:"column:short_desc;type:varchar(255)"`
	LongDesc   string         `gorm:"column:long_desc;type:varchar(1000)"`
	Link       string         `gorm:"column:link;type:varchar(255)"`
	Image      string         `gorm:"column:image;type:varchar(255)"`
	TechStack  pq.StringArray `gorm:"column:tech_stack;type:text[]"`
	Topic      pq.StringArray `gorm:"column:topic;type:text[]"`
	CohortID   int            `gorm:"column:cohort_id;not null"`
}

func (Project) TableName() string {
	return "projects.project"
}
