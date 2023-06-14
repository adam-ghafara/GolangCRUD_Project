-- Database: panamora

-- DROP DATABASE IF EXISTS panamora;

-- BUAT DATABASE MENGGUNAKAN POSTGRESQL

CREATE DATABASE panamora
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Indonesian_Indonesia.1252'
    LC_CTYPE = 'Indonesian_Indonesia.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE IF NOT EXISTS public.jenis_kendaraan
(
    id_jenis integer NOT NULL,
    jenis_kendaraan character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT jenis_kendaraan_pkey PRIMARY KEY (id_jenis)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.jenis_kendaraan
    OWNER to postgres;

CREATE TABLE IF NOT EXISTS public.data_servis_kendaraan
(
    id integer NOT NULL DEFAULT nextval('data_servis_kendaraan_id_seq'::regclass),
    id_jenis integer NOT NULL,
    nama_pemilik character varying(255) COLLATE pg_catalog."default" NOT NULL,
    nama_kendaraan character varying(255) COLLATE pg_catalog."default" NOT NULL,
    nomor_kendaraan integer NOT NULL,
    detail_servis text COLLATE pg_catalog."default",
    tanggal_servis date,
    status_servis character varying(50) COLLATE pg_catalog."default",
    CONSTRAINT data_servis_kendaraan_pkey PRIMARY KEY (id),
    CONSTRAINT data_servis_kendaraan_id_jenis_fkey FOREIGN KEY (id_jenis)
        REFERENCES public.jenis_kendaraan (id_jenis) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.data_servis_kendaraan
    OWNER to postgres;