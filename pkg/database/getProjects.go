package database

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type GetProjects struct {
	ID		 		uuid.UUID	 	`gorm:"column:id;type:uuid;default:gen_random_uuid()"`
	Member1			string			`gorm:"column:member1;type:varchar(255)"`
	Member2			string			`gorm:"column:member2;type:varchar(255)"`
	Member3			string			`gorm:"column:member3;type:varchar(255)"`
	Member4			string			`gorm:"column:member4;type:varchar(255)"`
	ProjectName 	string			`gorm:"column:name;type:varchar(255)"`
	ShortDesc		string			`gorm:"column:short_desc;type:varchar(255)"`
	LongDesc		string			`gorm:"column:long_desc;type:varchar(1000)"`
	Link			string			`gorm:"column:link;type:varchar(255)"`
	Image			string			`gorm:"column:image;type:varchar(255)"`
	TechStack		pq.StringArray	`gorm:"column:tech_stack;type:text[]"`
	Topic			pq.StringArray	`gorm:"column:topic;type:text[]"`
	Cohort			string			`gorm:"column:cohort;type:varchar(255)"`
}