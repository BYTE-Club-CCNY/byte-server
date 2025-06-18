-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists projects.project (
	uuid UUID PRIMARY KEY NOT NULL REFERENCES projects.team(uuid) ON DELETE SET NULL,
	name varchar(255),
	short_desc varchar(1000),
	long_desc varchar(1000),
	github varchar(255),
	image varchar(255),
	tech_stack text[],
	topic text[],
	cohort_id int NOT NULL REFERENCES users.cohort(cohort_id)
);