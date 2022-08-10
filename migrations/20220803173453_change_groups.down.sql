BEGIN;

ALTER TABLE groups DROP COLUMN mothergroup varchar;
ALTER TABLE groups DROP COLUMN subgroup boolean;
ALTER TABLE groups DROP COLUMN members integer;
ALTER TABLE groups ADD COLUMN members varchar[];

COMMIT;
