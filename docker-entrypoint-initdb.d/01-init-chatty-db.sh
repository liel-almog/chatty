#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE chatty;
EOSQL

# Switch to the 'chatty' database and create tables
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "chatty" <<-EOSQL
    -- Create sequences
    CREATE SEQUENCE IF NOT EXISTS messages_id_seq;
    CREATE SEQUENCE IF NOT EXISTS rooms_room_id_seq;

    -- Create rooms table
    CREATE TABLE IF NOT EXISTS public.rooms
    (
        id integer NOT NULL DEFAULT nextval('rooms_room_id_seq'::regclass),
        name text COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL,
        CONSTRAINT rooms_pkey PRIMARY KEY (id)
    )
    TABLESPACE pg_default;

    ALTER TABLE public.rooms
        OWNER to postgres;

    -- Create messages table
    CREATE TABLE IF NOT EXISTS public.messages
    (
        id integer NOT NULL DEFAULT nextval('messages_id_seq'::regclass),
        room_id integer,
        content text COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        username text COLLATE pg_catalog."default" NOT NULL,
        CONSTRAINT messages_pkey PRIMARY KEY (id),
        CONSTRAINT messages_room_id_fkey FOREIGN KEY (room_id)
            REFERENCES public.rooms (id) MATCH SIMPLE
            ON UPDATE NO ACTION
            ON DELETE NO ACTION
    )
    TABLESPACE pg_default;
    
    ALTER TABLE public.messages
        OWNER to postgres;

EOSQL