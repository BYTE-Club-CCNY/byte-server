-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists projects.project (
	id UUID PRIMARY KEY NOT NULL,
	name varchar(255),
	short_desc varchar(255),
	long_desc varchar(1000),
	link varchar(255),
	image varchar(255),
	tech_stack text[],
	topic text[],
	cohort_id int not null,
	CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES projects.teams(id) 
)