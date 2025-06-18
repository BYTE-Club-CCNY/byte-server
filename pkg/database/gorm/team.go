package gorm

type Team struct {
	uuid		string		 	`gorm:"primaryKey;column:uuid;type:uuid"`
	Member1		string			`gorm:"column:member1;type:uuid"`
	Member2		*string			`gorm:"column:member2;type:uuid"`
	Member3		*string			`gorm:"column:member3;type:uuid"`
	Member4		*string			`gorm:"column:member4;type:uuid"`
}

func (Team) TableName() string { return "projects.team" }