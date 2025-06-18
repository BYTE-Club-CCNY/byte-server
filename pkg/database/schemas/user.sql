create table if not exists users.user (
    uuid UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    active boolean NOT NULL DEFAULT true,
    first_name varchar(255) NOT NULL,
    middle_name varchar(255),
    last_name varchar(255) NOT NULL,
    personal_email varchar(255),
    cuny_email varchar(255),
    discord varchar(255),
    emplid varchar(255) NOT NULL
);