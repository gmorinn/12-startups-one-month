BEGIN;

-- rename verified to is_premium
ALTER TABLE "users" RENAME COLUMN "verified" TO "is_premium";

COMMIT;