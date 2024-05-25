-- Create "users" table
CREATE TABLE "users" ("id" character varying(36) NOT NULL, "name" character varying(200) NULL, "email" character varying(200) NULL, "password_hash" character varying(500) NULL, "username" character varying(200) NULL, PRIMARY KEY ("id"));
