CREATE SCHEMA fizzbuzz;

CREATE TABLE fizzbuzz.stats
(
    id          SERIAL PRIMARY KEY,
    int1        INTEGER             NOT NULL,
    int2        INTEGER             NOT NULL,
    max_limit   INTEGER             NOT NULL,
    str1        VARCHAR             NOT NULL    DEFAULT '',
    str2        VARCHAR             NOT NULL    DEFAULT '',

    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX stats_idx ON fizzbuzz.stats(int1, int2, max_limit, str1, str2);