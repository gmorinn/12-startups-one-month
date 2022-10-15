BEGIN;

-- constraint user_id_viewer is different from profil_id_viewed
ALTER TABLE "viewers" ADD CONSTRAINT viewchk CHECK ("user_id_viewer" != "profil_id_viewed");


COMMIT;