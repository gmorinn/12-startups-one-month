BEGIN;

-- Change password column to text
ALTER TABLE users
ALTER COLUMN password TYPE text;

COMMIT;