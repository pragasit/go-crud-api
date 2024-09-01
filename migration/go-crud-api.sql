CREATE TABLE "basketball_court" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "location" varchar
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar,
  "email" varchar,
  "created_at" timestamp
);

CREATE TABLE "reservations" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "court_id" integer,
  "reservation_date" date,
  "reservation_time" time,
  "status" varchar,
  "created_at" timestamp
);

CREATE TABLE "queue" (
  "id" integer PRIMARY KEY,
  "reservation_id" integer,
  "queue_number" integer,
  "created_at" timestamp
);

COMMENT ON COLUMN "reservations"."status" IS 'Pending, Paid, Cancelled';

ALTER TABLE "reservations" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reservations" ADD FOREIGN KEY ("court_id") REFERENCES "basketball_court" ("id");

ALTER TABLE "queue" ADD FOREIGN KEY ("reservation_id") REFERENCES "reservations" ("id");
