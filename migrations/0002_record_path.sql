-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "records_path" (
  "id" SERIAL,
  "uuid" uuid PRIMARY KEY,
  "storage_address" varchar NOT NULL,
  "direction" varchar NOT NULL,
  "year" varchar NOT NULL,
  "month" varchar NOT NULL,
  "day" varchar NOT NULL,
  "name" varchar NOT NULL
);


COMMENT ON COLUMN "records_path"."uuid" IS 'call uuid';
COMMENT ON COLUMN "records_path"."storage_address" IS 'record storage host address';
COMMENT ON COLUMN "records_path"."direction" IS 'inc or out';
COMMENT ON COLUMN "records_path"."year" IS 'year of the call record';
COMMENT ON COLUMN "records_path"."month" IS 'month of the call record';
COMMENT ON COLUMN "records_path"."day" IS 'day of the call record';
COMMENT ON COLUMN "records_path"."name" IS 'record file name';



-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "records_path";
