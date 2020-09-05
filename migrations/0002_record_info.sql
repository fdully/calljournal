-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "record_info" (
  "id" SERIAL,
  "uuid" uuid PRIMARY KEY,
  "addr" varchar NOT NULL,
  "dirc" varchar NOT NULL,
  "year" varchar NOT NULL,
  "mont" varchar NOT NULL,
  "rday" varchar NOT NULL,
  "rnam" varchar NOT NULL
);

COMMENT ON COLUMN "record_info"."uuid" IS 'record uuid';
COMMENT ON COLUMN "record_info"."addr" IS 'record storage address';
COMMENT ON COLUMN "record_info"."dirc" IS 'call direction';
COMMENT ON COLUMN "record_info"."year" IS 'record year';
COMMENT ON COLUMN "record_info"."mont" IS 'record month';
COMMENT ON COLUMN "record_info"."rday" IS 'record day';
COMMENT ON COLUMN "record_info"."rnam" IS 'record file name';



-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "record_info";
