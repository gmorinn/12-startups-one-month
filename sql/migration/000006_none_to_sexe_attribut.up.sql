BEGIN;

-- add 'none' to sexe enum
ALTER TYPE sexe ADD VALUE 'none';

-- set "sexe" sexe DEFAULT 'none'
ALTER TABLE "users" ALTER COLUMN "sexe" SET DEFAULT 'none';

-- add 'none' to formule enum
ALTER TYPE formule ADD VALUE 'none';

-- set "formule" formule DEFAULT 'none'
ALTER TABLE "users" ALTER COLUMN "formule" SET DEFAULT 'none';

COMMIT;