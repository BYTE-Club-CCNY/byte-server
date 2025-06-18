package gorm

type Cohort struct {
	CohortId		string		 	`gorm:"primaryKey;column:cohort_id;type:int"`
	CohortName		string			`gorm:"column:cohort_name;type:string"`
}

func (Cohort) TableName() string { return "users.cohort" }