-- modify "users" table
ALTER TABLE "users" ADD COLUMN "role" character varying(255) NOT NULL DEFAULT 'user';
