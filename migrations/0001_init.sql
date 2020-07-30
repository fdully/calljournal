-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "base_calls" (
  "id" SERIAL,
  "uuid" uuid PRIMARY KEY,
  "clid" varchar NOT NULL,
  "clna" varchar NOT NULL,
  "dest" varchar NOT NULL,
  "dirc" varchar NOT NULL,
  "stti" timestamptz NOT NULL,
  "durs" int,
  "bils" int,
  "recd" bool,
  "recs" int,
  "recl" varchar,
  "rtag" varchar,
  "epos" int8 NOT NULL,
  "epoa" int8,
  "epoe" int8 NOT NULL,
  "wbye" varchar,
  "hang" varchar,
  "code" varchar
);

COMMENT ON COLUMN "base_calls"."uuid" IS 'call uuid';
COMMENT ON COLUMN "base_calls"."clid" IS 'caller id number';
COMMENT ON COLUMN "base_calls"."clna" IS 'caller id name';
COMMENT ON COLUMN "base_calls"."dest" IS 'callee number';
COMMENT ON COLUMN "base_calls"."dirc" IS 'call direction';
COMMENT ON COLUMN "base_calls"."stti" IS 'call start time';
COMMENT ON COLUMN "base_calls"."durs" IS 'call duration include signals';
COMMENT ON COLUMN "base_calls"."bils" IS 'call talk duration without signaling time';
COMMENT ON COLUMN "base_calls"."recd" IS 'call recorded';
COMMENT ON COLUMN "base_calls"."recs" IS 'call record duration';
COMMENT ON COLUMN "base_calls"."recl" IS 'call record file';
COMMENT ON COLUMN "base_calls"."rtag" IS 'call record privacy tag';
COMMENT ON COLUMN "base_calls"."epos" IS 'call start time unix';
COMMENT ON COLUMN "base_calls"."epoa" IS 'call answer time unix';
COMMENT ON COLUMN "base_calls"."epoe" IS 'call end time unix';
COMMENT ON COLUMN "base_calls"."wbye" IS 'call hangup initiator';
COMMENT ON COLUMN "base_calls"."hang" IS 'call hangup status';
COMMENT ON COLUMN "base_calls"."code" IS 'call status code';


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "base_calls";
