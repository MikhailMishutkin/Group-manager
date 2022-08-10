BEGIN;

ALTER TABLE groups DROP COLUMN members;

ALTER TABLE groups ADD COLUMN members integer;

ALTER TABLE groups ADD COLUMN subgroup boolean;

ALTER TABLE groups ADD COLUMN mothergroup varchar;

COMMIT;
