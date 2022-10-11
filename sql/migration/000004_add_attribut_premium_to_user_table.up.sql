BEGIN;

-- remove badgechk constraint
ALTER TABLE "premium" DROP CONSTRAINT "badgechk";

-- remove askchk constraint
ALTER TABLE "premium" DROP CONSTRAINT "askchk";

-- remove formule to premium table
ALTER TABLE "premium" DROP COLUMN "formule";

-- remove ask to premium table
ALTER TABLE "premium" DROP COLUMN "ask";

-- remove badge to premium table
ALTER TABLE "premium" DROP COLUMN "badge";

-- add "ask" int NOT NULL DEFAULT 15 CONSTRAINT askchk CHECK (ask >= 0 AND ask <= 70) to table "user"
ALTER TABLE "users" ADD COLUMN "ask" int NOT NULL DEFAULT 15 CONSTRAINT askchk CHECK (ask >= 0 AND ask <= 70);

--  add "badge" boolean NOT NULL DEFAULT false,
ALTER TABLE "users" ADD COLUMN "badge" boolean NOT NULL DEFAULT false;

-- add "formule" formule DEFAULT NULL
ALTER TABLE "users" ADD COLUMN "formule" formule DEFAULT NULL;

COMMIT;