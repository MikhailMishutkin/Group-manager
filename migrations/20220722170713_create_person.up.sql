CREATE TABLE persons (
    id bigserial not null primary key,
    person_name varchar not null,
    surname varchar not null,
    year_of_birth integer not null,
    groupname varchar not null
);

