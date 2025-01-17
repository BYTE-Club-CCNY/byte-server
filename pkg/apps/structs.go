package apps

type EditCohortDraftBody struct {
	ProjectID		string		`json:"project_id" form:"project_id" validate:"required"`
	ProjectSchema 	string		`json:"project_schema" form:"project_schema" validate:"required"`
}