CREATE TABLE groups (
    id smallserial not null primary key,
    groupname varchar not null,
    subgroup boolean,
    mothergroup varchar,
    members integer not null       
);