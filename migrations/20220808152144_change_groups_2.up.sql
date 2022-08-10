BEGIN;

ALTER TABLE groups DROP COLUMN members;

ALTER TABLE groups DROP COLUMN subgroups;

ALTER TABLE groups ADD COLUMN members integer NOT NULL;

COMMIT;