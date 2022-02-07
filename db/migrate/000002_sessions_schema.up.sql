CREATE TABLE "Sessions" (
  "id" bigserial,
  "user_id" bigserial NOT NULL,
  "token" varchar NOT NULL,
  "time" timestamp NOT NULL DEFAULT (now()),
  "device" varchar NOT NULL,
  "status" varchar,
  FOREIGN KEY ("user_id") REFERENCES "Users" ("id"),
  PRIMARY KEY ("id", "user_id"),

);

