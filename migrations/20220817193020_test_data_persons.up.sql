BEGIN;
INSERT INTO persons (person_name, surname, year_of_birth, groupname) VALUES ('Vasya', 'F', 1990, 'testgroup');
INSERT INTO persons (person_name, surname, year_of_birth, groupname) VALUES ('Kolya', 'S', 1991, 'testgroup1');
INSERT INTO persons (person_name, surname, year_of_birth, groupname) VALUES ('Petya', 'T', 1992, 'subtestgroup');
INSERT INTO persons (person_name, surname, year_of_birth, groupname) VALUES ('Katya', 'Fo', 1993, 'subtestgroup1');
INSERT INTO persons (person_name, surname, year_of_birth, groupname) VALUES ('Olya', 'Fi', 1994, 'subtestgroup2');
UPDATE groups SET members = 1;
COMMIT;