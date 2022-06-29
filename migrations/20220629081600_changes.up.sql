-- create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "age" bigint NOT NULL, "name" character varying NOT NULL DEFAULT 'unknown', PRIMARY KEY ("id"));
-- create "cars" table
CREATE TABLE "cars" ("id" uuid NOT NULL, "model" character varying NOT NULL, "registered_at" timestamptz NOT NULL, "user_cars" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "cars_users_cars" FOREIGN KEY ("user_cars") REFERENCES "users" ("id") ON DELETE NO ACTION);
-- create "groups" table
CREATE TABLE "groups" ("id" uuid NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- create "group_users" table
CREATE TABLE "group_users" ("group_id" uuid NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("group_id", "user_id"), CONSTRAINT "group_users_group_id" FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON DELETE CASCADE, CONSTRAINT "group_users_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE);
