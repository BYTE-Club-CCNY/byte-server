package projects

import (
	"byteserver/pkg/database/gorm"

	"github.com/google/uuid"
)

type GetProjectsBody struct {
	Cohort 		int
	Name 		string
	N_records 	int
	Page 		int
}

type getProjectsReturn struct {
	gorm.Project
	Member1		string	`gorm:"column:member1;type:text"`
	Member2 	string	`gorm:"column:member2;type:text"`
	Member3		string	`gorm:"column:member3;type:text"`
	Member4		string	`gorm:"column:member4;type:text"`
	Cohort		string	`gorm:"column:cohort;type:text"`
}

type AddProjectsBody struct {
	Name		string			`json:"name" form:"name" validate:"required"`
	Short_Desc 	string			`json:"short_desc" form:"short_desc" validate:"required"`
	Long_Desc 	string			`json:"long_desc" form:"long_desc" validate:"required"`
	Member1 	uuid.UUID		`json:"member1" form:"member1" validate:"required"`
	Member2		uuid.UUID		`json:"member2" form:"member2"`
	Member3		uuid.UUID		`json:"member3" form:"member3"`
	Member4 	uuid.UUID		`json:"member4" form:"member4"`
	Link		string			`json:"link" form:"link" validate:"required"`
	Image		string			`json:"image" form:"image" validate:"required"`
	Tech_Stack	string			`json:"tech_stack" form:"tech_stack" validate:"required"`
	Topic		string			`json:"topic" form:"topic" validate:"required"`
	Cohort		int				`json:"cohort" form:"cohort" validate:"required"`
}
