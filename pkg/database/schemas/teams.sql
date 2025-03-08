-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists projects.teams(
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    member1 UUID not null,
    member2 UUID,
    member3 UUID,
    member4 UUID
);