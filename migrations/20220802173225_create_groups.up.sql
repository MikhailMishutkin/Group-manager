CREATE TABLE groups (
    id smallserial not null primary key,
    groupname varchar not null,
    members varchar[],
    subgroups varchar[]
);