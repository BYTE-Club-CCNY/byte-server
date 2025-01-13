package database

type GetProjects struct {
	ID				string
	Names			[]string
	ProjectName 	string
	ShortDesc		string
	LongDesc		string
	Link			string
	Image			string
	TechStack		[]string
	Topic			[]string
	Cohort			string
}
