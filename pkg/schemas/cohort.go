package schema

type Cohort struct {
	CohortID   int    `gorm:"column:cohort_id;primaryKey"`
	CohortName string `gorm:"column:cohort_name;type:varchar(255)"`
}

func (Cohort) TableName() string {
	return "users.cohort"
}
