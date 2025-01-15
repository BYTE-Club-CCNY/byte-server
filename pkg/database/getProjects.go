package database

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type GetProjects struct {
	ID		 		uuid.UUID	 	`gorm:"column:id;type:uuid;default:gen_random_uuid()"`
	Member1			string			`gorm:"column:member1;type:varchar(255)"`
	Member2			string			`gorm:"column:member2;type:varchar(255)"`
	Member3			string			`gorm:"column:member3;type:varchar(255)"`
	Member4			string			`gorm:"column:member4;type:varchar(255)"`
	ProjectName 	string			`gorm:"column:name;type:varchar(255)"`
	ShortDesc		string			`gorm:"column:short_desc;type:varchar(255)"`
	LongDesc		string			`gorm:"column:long_desc;type:varchar(1000)"`
	Link			string			`gorm:"column:link;type:varchar(255)"`
	Image			string			`gorm:"column:image;type:varchar(255)"`
	TechStack		pq.StringArray	`gorm:"column:tech_stack;type:text[]"`
	Topic			pq.StringArray	`gorm:"column:topic;type:text[]"`
	Cohort			string			`gorm:"column:cohort;type:varchar(255)"`
}

var GetProjectsQuery string = 
"SELECT p.id, p.name, p.topic, " +
"p.name, p.short_desc, p.long_desc, " + 
"p.link, p.image, p.tech_stack, " +
"CONCAT(u1.first_name, ' ', u1.middle_name, ' ', u1.last_name) AS member1, " +
"CONCAT(u2.first_name, ' ', u2.middle_name, ' ', u2.last_name) AS member2, " +
"CONCAT(u3.first_name, ' ', u3.middle_name, ' ', u3.last_name) AS member3, " +
"CONCAT(u4.first_name, ' ', u4.middle_name, ' ', u4.last_name) AS member4, " +
"c.cohort_name as cohort " +
"FROM projects.project p " +
"INNER JOIN projects.teams t ON p.id = t.id " +
"LEFT JOIN users.people u1 ON t.member1 = u1.uid " +
"LEFT JOIN users.people u2 ON t.member2 = u2.uid " +
"LEFT JOIN users.people u3 ON t.member3 = u3.uid " +
"LEFT JOIN users.people u4 ON t.member4 = u4.uid " +
"INNER JOIN users.cohort c ON p.cohort_id = c.cohort_id;"