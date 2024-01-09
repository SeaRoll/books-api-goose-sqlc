-- +goose Up
CREATE TABLE conditions (
   time        TIMESTAMPTZ       NOT NULL,
   location    TEXT              NOT NULL,
   device      TEXT              NOT NULL,
   temperature DOUBLE PRECISION  NULL,
   humidity    DOUBLE PRECISION  NULL
);

SELECT create_hypertable('conditions', by_range('time'));

-- Learn more about indexing
-- https://docs.timescale.com/use-timescale/latest/schema-management/about-indexing/

-- +goose Down
DROP TABLE conditions;
