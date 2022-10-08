
BEGIN;

CREATE EXTENSION pgcrypto;

CREATE TYPE "role" AS ENUM (
  'admin',
  'pro',
  'user'
);


CREATE TYPE "status_payment" AS ENUM (
  'active',
  'cancel',
  'past_due'
);

CREATE TYPE "formule" AS ENUM (
  'basic',
  'gold',
  'diamond'
);


CREATE TYPE "sexe" AS ENUM (
  'man',
  'woman',
  'other'
);

CREATE TYPE "goals" AS ENUM (
  'prise_de_masse',
  'perte_de_poids',
  'prise_de_force',
  'sèche',
  'cardio'
);

CREATE TYPE "performance" AS ENUM (
  'developpe_couche',
  'squat',
  'developpe_incline',
  'developpe_epaule',
  'souleve_terre',
  'leg_press',
  'leg_curl',
  'leg_extension',
  'curl_barre',
  'hip_thrust'
);

CREATE TABLE "viewers" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp CONSTRAINT deletedchk CHECK (deleted_at > created_at),
  "user_id_viewer" uuid NOT NULL,
  "profil_id_viewed" uuid NOT NULL,
  "date_viewed" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp CONSTRAINT deletedchk CHECK (deleted_at > created_at),
  "email" text NOT NULL CONSTRAINT emailchk CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
  "password" varchar(255) CONSTRAINT passwordchk CHECK (char_length(password) >= 6) NOT NULL,
  "firstname" varchar(25) CONSTRAINT firstnamechk CHECK (char_length(firstname) > 1 AND char_length(firstname) <= 25 AND firstname ~ '^[^0-9]*$') DEFAULT NULL,
  "lastname" varchar(25) CONSTRAINT lastnamechk CHECK (char_length(lastname) >1 AND char_length(lastname) <= 25 AND  lastname ~ '^[^0-9]*$') DEFAULT NULL,
  "role" role ARRAY NOT NULL DEFAULT '{user}',
  "age" int CONSTRAINT agechk CHECK (age > 0 AND age < 150) DEFAULT NULL,
  "sexe" sexe DEFAULT NULL,
  "goals" goals ARRAY DEFAULT NULL,
  "ideal_partners" text DEFAULT NULL,
  "profile_picture" text DEFAULT NULL,
  "verified" boolean DEFAULT false,
  "city" text DEFAULT NULL
);

CREATE TABLE "premium" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp CONSTRAINT deletedchk CHECK (deleted_at > created_at),
  "user_id" uuid NOT NULL,
  "start_date" timestamp NOT NULL DEFAULT (now()),
  "status_payment" status_payment NOT NULL DEFAULT 'active',
  "stripe_customer_id" text NOT NULL,
  "stripe_subscription_id" text NOT NULL,
  "formule" formule NOT NULL DEFAULT 'basic',
  "badge" boolean NOT NULL DEFAULT false CONSTRAINT badgechk CHECK (badge = (formule = 'diamond' OR formule = 'gold')),
  "ask" int NOT NULL DEFAULT 15 CONSTRAINT askchk CHECK (ask >= 0 AND ask <= 30)
);

CREATE TABLE "avis" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp CONSTRAINT deletedchk CHECK (deleted_at > created_at),
  "user_id_target" uuid NOT NULL,
  "user_id_writer" uuid NOT NULL,
  "note" int NOT NULL CONSTRAINT notechk CHECK (note >= 0 AND note <= 5),
  "comment" text NOT NULL
);


CREATE TABLE "refresh_token" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
  "deleted_at" timestamptz,
  "token" text NOT NULL,
  "ip" text NOT NULL,
  "user_agent" text NOT NULL,
  "expir_on" timestamptz NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "files" (
    "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
    "created_at" timestamptz NOT NULL DEFAULT (NOW()),
    "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
    "deleted_at" timestamptz,
    "name" varchar(255),
    "url" text,
    "mime" text,
    "size" bigint
);

ALTER TABLE "refresh_token" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "viewers" ADD FOREIGN KEY ("user_id_viewer") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "viewers" ADD FOREIGN KEY ("profil_id_viewed") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "premium" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "avis" ADD FOREIGN KEY ("user_id_target") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "avis" ADD FOREIGN KEY ("user_id_writer") REFERENCES "users" ("id") ON DELETE CASCADE;

COMMIT;