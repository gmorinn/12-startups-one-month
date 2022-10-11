BEGIN;

-- remove "sèche" from enum "goals"
ALTER TYPE "goals" RENAME TO "goals_old";
CREATE TYPE "goals" AS ENUM (
  'prise_de_masse',
  'perte_de_poids',
  'prise_de_force',
  'garder_la_forme',
  'prise_de_muscle',
  'cardio'
);

COMMIT;