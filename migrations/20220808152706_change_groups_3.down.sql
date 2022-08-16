BEGIN;

ALTER TABLE groups DROP COLUMN members;

ALTER TABLE groups ADD COLUMN members integer;

ALTER TABLE groups ADD COLUMN subgroups varchar;

COMMIT;