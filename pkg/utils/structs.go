package utils

type AddUsersBody struct{
	Name			string		`json:"name" form:"name" validate:"required"`
	CunyEmail		string		`json:"cuny_email" form:"cuny_email" validate:"required"`
	Emplid			string		`json:"emplid" form:"emplid" validate:"required"`
	PersonalEmail	string		`json:"personal_email" form:"personal_email"`
	Discord			string		`json:"discord" form:"discord"`
}