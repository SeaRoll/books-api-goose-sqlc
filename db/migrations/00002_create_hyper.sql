-- +goose Up
CREATE TABLE conditions (
   time        TIMESTAMPTZ       NOT NULL,
   location    TEXT              NOT NULL,
   device      TEXT              NOT NULL,
   temperature DOUBLE PRECISION  NOT NULL,
   humidity    DOUBLE PRECISION  NOT NULL
);

SELECT create_hypertable('conditions', by_range('time'));

-- https://docs.timescale.com/use-timescale/latest/schema-management/about-indexing/
CREATE INDEX ON conditions (location, time DESC);


-- +goose Down
DROP TABLE conditions;
