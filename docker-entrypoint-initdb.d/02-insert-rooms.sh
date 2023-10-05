#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "chatty" <<-EOSQL
    INSERT INTO public.rooms (name, created_at) VALUES
    ('צאט כללי', CURRENT_TIMESTAMP),
    ('צאט טכני', CURRENT_TIMESTAMP),
    ('רנדומלי', CURRENT_TIMESTAMP);
EOSQL
