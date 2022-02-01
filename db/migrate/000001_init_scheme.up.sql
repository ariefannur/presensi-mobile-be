
CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "pasword" varchar NOT NULL,
  "user_type" varchar NOT NULL
);

CREATE TABLE "Presensi" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "foto" varchar NOT NULL,
  "time" timestamp NOT NULL DEFAULT (now()),
  "lat" decimal NOT NULL,
  "lng" decimal NOT NULL,
  "alamat" varchar NOT NULL
);

ALTER TABLE "Presensi" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

CREATE INDEX ON "Users" ("id");

CREATE INDEX ON "Presensi" ("user_id");
