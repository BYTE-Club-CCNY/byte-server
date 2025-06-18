package gorm

type CohortHistory struct {
	UserId		string		 	`gorm:"column:user_id;type:uuid"`
	CohortId	string			`gorm:"column:cohort_id;type:int"`
}

func (CohortHistory) TableName() string { return "users.cohort_history" }