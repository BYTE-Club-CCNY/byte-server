package schema

import (
	"github.com/lib/pq"
)

type Project struct {
	Uid       int            `gorm:"column:uid;primaryKey;autoIncrement"`
	Name      string         `gorm:"column:name;type:varchar;not null"`
	ShortDesc string         `gorm:"column:short-desc;type:varchar;not null"`
	LongDesc  string         `gorm:"column:long-desc;type:varchar"`
	Team      pq.StringArray `gorm:"column:team;type:text[]"`
	Link      string         `gorm:"column:link;type:varchar"`
	Image     string         `gorm:"column:image;type:varchar"`
	TechStack pq.StringArray `gorm:"column:tech-stack;type:text[]"`
	Cohort    string         `gorm:"column:cohort;type:varchar"`
	Topic     pq.StringArray `gorm:"column:topic;type:text[]"`
}

func (Project) TableName() string {
	return "public.projects"
}
