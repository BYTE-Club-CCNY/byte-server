CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table if not exists projects.teams(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	member1 UUID NOT NULL,
	member2 UUID,
	member3 UUID,
	member4 UUID
)