-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "base_calls" (
  "id" SERIAL,
  "uuid" uuid PRIMARY KEY,
  "username" varchar,
  "caller_id_name" varchar,
  "caller_id_number" varchar,
  "destination_number" varchar NOT NULL,
  "cj_direction" varchar NOT NULL,
  "start_stamp" timestamp NOT NULL,
  "duration" int NOT NULL,
  "billsec" int NOT NULL,
  "record_seconds" int NOT NULL,
  "record_name" varchar,
  "start_epoch" int8 NOT NULL,
  "answer_epoch" int8 NOT NULL,
  "end_epoch" int8 NOT NULL,
  "sip_hangup_disposition" varchar,
  "hangup_cause" varchar,
  "sip_term_status" varchar
);

COMMENT ON COLUMN "base_calls"."uuid" IS 'call uuid';
COMMENT ON COLUMN "base_calls"."username" IS 'username of caller';
COMMENT ON COLUMN "base_calls"."caller_id_name" IS 'caller id name';
COMMENT ON COLUMN "base_calls"."caller_id_number" IS 'caller id number';
COMMENT ON COLUMN "base_calls"."destination_number" IS 'destination number';
COMMENT ON COLUMN "base_calls"."cj_direction" IS 'call direction';
COMMENT ON COLUMN "base_calls"."start_stamp" IS 'call start time';
COMMENT ON COLUMN "base_calls"."duration" IS 'call duration include signals';
COMMENT ON COLUMN "base_calls"."billsec" IS 'call talk duration without signaling time';
COMMENT ON COLUMN "base_calls"."record_seconds" IS 'call record duration';
COMMENT ON COLUMN "base_calls"."record_name" IS 'call record file name';
COMMENT ON COLUMN "base_calls"."start_epoch" IS 'call start time unix';
COMMENT ON COLUMN "base_calls"."answer_epoch" IS 'call answer time unix';
COMMENT ON COLUMN "base_calls"."end_epoch" IS 'call end time unix';
COMMENT ON COLUMN "base_calls"."sip_hangup_disposition" IS 'call hangup initiator';
COMMENT ON COLUMN "base_calls"."hangup_cause" IS 'call hangup status';
COMMENT ON COLUMN "base_calls"."sip_term_status" IS 'call status code';


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "base_calls";
