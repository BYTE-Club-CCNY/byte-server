package database

type GetProjects struct {
	ID				string
	Member1			string
	Member2			string
	Member3			string
	Member4			string
	ProjectName 	string
	ShortDesc		string
	LongDesc		string
	Link			string
	Image			string
	TechStack		[]string
	Topic			[]string
	Cohort			string
}
