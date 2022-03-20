#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$postgres" --dbname "postgres" <<-EOSQL
  CREATE DATABASE testdb;
  \connect testdb
  BEGIN;
    CREATE TABLE users (
        id varchar(50) primary key,
        name varchar(50),
        birthday varchar(50)
    );
   COMMIT;
EOSQL