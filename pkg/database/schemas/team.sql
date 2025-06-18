create table if not exists projects.team (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    member1 UUID NOT NULL REFERENCES users.user(uuid) ON DELETE SET NULL,
    member2 UUID REFERENCES users.user(uuid) ON DELETE SET NULL,
    member3 UUID REFERENCES users.user(uuid) ON DELETE SET NULL,
    member4 UUID REFERENCES users.user(uuid) ON DELETE SET NULL
);