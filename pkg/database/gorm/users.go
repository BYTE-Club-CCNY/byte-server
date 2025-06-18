package gorm

type User struct {
	ID		 		string		 	`gorm:"primaryKey;column:uuid;type:uuid;default:gen_random_uuid()"`
	Active			bool			`gorm:"column:active;type:boolean"`
	FirstName		string			`gorm:"column:first_name;type:varchar(255)"`
	MiddleName		*string			`gorm:"column:middle_name;type:varchar(255)"`
	LastName		string			`gorm:"column:last_name;type:varchar(255)"`
	PersonalEmail 	*string			`gorm:"column:personal_email;type:varchar(255)"`
	CUNYEmail		*string			`gorm:"column:cuny_email;type:varchar(255)"`
	Discord			*string			`gorm:"column:discord;type:varchar(1000)"`
	Emplid			string			`gorm:"column:emplid;type:varchar(255)"`
}

func (User) TableName() string { return "users.user" }