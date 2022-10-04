CREATE TABLE persons (
    id smallserial not null primary key,
    person_name varchar not null,
    surname varchar not null,
    year_of_birth integer not null,
    groupname varchar not null,
    group_id integer,
    FOREIGN KEY (group_id) REFERENCES groups (id)
);