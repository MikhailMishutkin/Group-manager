BEGIN;
DELETE FROM persons;
UPDATE groups SET members = 0;
COMMIT;