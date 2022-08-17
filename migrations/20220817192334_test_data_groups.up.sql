BEGIN;
INSERT INTO groups (groupname, subgroup, members) VALUES ('testgroup', false, 0);
INSERT INTO groups (groupname, subgroup, members) VALUES ('testgroup1', false, 0);
INSERT INTO groups (groupname, subgroup, mothergroup, members) VALUES ('subtestgroup', true, 'testgroup', 0);
INSERT INTO groups (groupname, subgroup, mothergroup, members) VALUES ('subtestgroup1', true, 'testgroup', 0);
INSERT INTO groups (groupname, subgroup, mothergroup, members) VALUES ('subtestgroup2', true, 'testgroup1', 0);
COMMIT;