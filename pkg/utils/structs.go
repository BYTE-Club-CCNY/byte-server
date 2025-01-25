package utils

type AddUsersBody struct{
	Name			string		`json:"name" form:"name" validate:"required"`
	CunyEmail		string		`json:"cuny_email" form:"cuny_email" validate:"required"`
	Emplid			string		`json:"emplid" form:"emplid" validate:"required"`
	PersonalEmail	string		`json:"personal_email" form:"personal_email"`
	Discord			string		`json:"discord" form:"discord"`
}

type EditDraft struct {
	Cohort_id string `json:"cohort_id" form:"cohort_id" validate:"required"`
	Deadline string `json:"deadline" form:"deadline" validate:"required"`
	Questions []Question `json:"questions" form:"questions" validate:"required"` 
}

type Question struct {
	Question   string   `json:"question"`
	AnswerType string   `json:"answerType"`
	Options    []string `json:"options"`
}
