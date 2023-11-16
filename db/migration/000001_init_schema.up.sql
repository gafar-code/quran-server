CREATE TABLE "surah" (
  "id" BIGSERIAL PRIMARY KEY,
  "page" bigint NOT NULL,
  "ar" varchar NOT NULL,
  "number" bigint NOT NULL,
  "tr" varchar NOT NULL,
  "idn" varchar NOT NULL,
  "type" varchar NOT NULL,
  "total_ayah" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ayah" (
  "id" BIGSERIAL PRIMARY KEY,
  "page" bigint NOT NULL,
  "surah_number" bigint NOT NULL,
  "number" bigint NOT NULL,
  "number_in_surah" bigint NOT NULL,
  "ar" varchar NOT NULL,
  "tr" varchar NOT NULL,
  "idn" varchar NOT NULL,
  "audio_url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "surah" ("id");
