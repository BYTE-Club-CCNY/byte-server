package gorm

type Project struct {
	Uuid		string		`gorm:"primaryKey;column:uuid;type:uuid"`
	Name		*string		`gorm:"column:name;type:varchar(255)"`
	ShortDesc	*string		`gorm:"column:short_desc;type:varchar(1000)"`
	LongDesc	*string		`gorm:"column:long_desc;type:varchar(1000)"`
	Github		*string		`gorm:"column:github;type:varchar(255)"`
	Image		*string		`gorm:"column:image;type:varchar(255)"`
	TechStack	*[]string	`gorm:"column:tech_stack;type:text[]"`
	Topic		*[]string	`gorm:"column:topic;type:text[]"`
	CohortId	string		`gorm:"column:cohort_id;type:string"`
}

func (Project) TableName() string { return "projects.project" }