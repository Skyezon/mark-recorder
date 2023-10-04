-- check if database already created if not create it
SELECT 'CREATE DATABASE assignment_mezink' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'assignment_mezink')\gexec

CREATE TABLE IF NOT EXISTS records(
    id serial UNIQUE PRIMARY KEY,
    name varchar(256),
    marks integer[],
    createdAt timestamptz 
);

DO $$ 
BEGIN
  IF NOT EXISTS (SELECT * FROM records WHERE id = 1) THEN
    INSERT INTO records (name, marks, createdAt) VALUES ('TESTING', '{100,200,300}', '2016-02-12:T12:00:01Z');
  END IF;
END $$;


CREATE INDEX idx_records_date on records (createdAt);
