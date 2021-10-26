CREATE TABLE "employee" (
  "id" serial PRIMARY KEY,
  "name" text NOT NULL,
  "age" integer NOT NULL,
  "gender" varchar(1) NOT NULL,
  "address" text NOT NULL,
  "department" text NOT NULL,
  "mobile_number" text NOT NULL,
  "created_at" timestamp with time zone NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp with time zone NOT NULL DEFAULT (CURRENT_TIMESTAMP)
)