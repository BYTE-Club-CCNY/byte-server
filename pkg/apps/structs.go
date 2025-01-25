package apps

type Cohort struct {
	Cohort_id string `json:"cohort_id" form:"cohort_id" validate:"required"`
}