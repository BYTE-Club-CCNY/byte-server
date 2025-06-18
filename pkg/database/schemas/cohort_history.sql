create table if not exists users.cohort_history (
	user_id uuid REFERENCES users.user(uuid) ON DELETE CASCADE,
	cohort_id int REFERENCES users.cohort(cohort_id)
);