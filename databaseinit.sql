SELECT 'CREATE DATABASE assignment_mezink' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'assignment_mezink')\gexec
CREATE TABLE IF NOT EXISTS records(
    id integer UNIQUE PRIMARY KEY,
    name varchar(256),
    marks integer[],
    createdAt datetime
)

CREATE INDEX idx_records_date on records (createdAt);
