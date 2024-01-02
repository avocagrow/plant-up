-- Create function "nuuid()"
CREATE OR REPLACE FUNCTION nuuid()
RETURNS TEXT
LANGUAGE SQL
AS $$
    SELECT
    (REPLACE(CAST(gen_random_uuid() as TEXT), '-', ''))
$$;
-- Create "plants" table
CREATE TABLE "public"."plants" (
 "id" text NOT NULL DEFAULT nuuid(),
 "name" text NOT NULL,
 "botanical_name" text NULL,
 "description" text NULL,
 "water_pref" text NULL,
 "light_pref" text NULL,
 "humidity_pref" text NULL,
 "created_at" timestamp NOT NULL DEFAULT now(),
 "updated_at" timestamp NOT NULL DEFAULT now(),
 "deleted_at" timestamp NULL,
 PRIMARY KEY ("id")
);
-- Create "users" table
CREATE TABLE "public"."users" (
 "id" text NOT NULL DEFAULT nuuid(),
 "name" text NULL,
 "display_name" text NOT NULL, 
 "email" text NOT NULL,
 "phone" text NOT NULL,
 "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 "deleted_at" timestamp NULL,
 PRIMARY KEY ("id"),
 UNIQUE ("email", "deleted_at"),
 UNIQUE ("phone", "deleted_at")
);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "public"."users" ("email");
-- Create index "users_phone_key" to table: "users"
CREATE UNIQUE INDEX "users_phone_key" ON "public"."users" ("phone");
