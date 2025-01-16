package projects

type GetProjectsBody struct {
	Cohort 		int
	Name 		string
	N_records 	int
	Page 		int
}

type AddProjectsBody struct {
	Name		string			`json:"name" form:"name" validate:"required"`
	Short_Desc 	string			`json:"short_desc" form:"short_desc" validate:"required"`
	Long_Desc 	string			`json:"long_desc" form:"long_desc" validate:"required"`
	Member1 	string			`json:"member1" form:"member1" validate:"required"`
	Member2		string			`json:"member2" form:"member2"`
	Member3		string			`json:"member3" form:"member3"`
	Member4 	string			`json:"member4" form:"member4"`
	Link		string			`json:"link" form:"link" validate:"required"`
	Image		string			`json:"image" form:"image" validate:"required"`
	Tech_Stack	string		`json:"tech_stack" form:"tech_stack" validate:"required"`
	Topic		string		`json:"topic" form:"topic" validate:"required"`
	Cohort		int				`json:"cohort" form:"cohort" validate:"required"`
}
