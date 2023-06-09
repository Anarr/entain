--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5
-- Dumped by pg_dump version 14.4

-- Started on 2022-10-03 17:12:33 +04

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2 (class 3079 OID 16938)
-- Name: adminpack; Type: EXTENSION; Schema: -; Owner: -
--

CREATE
EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;


--
-- TOC entry 3647 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION adminpack; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';

DROP SCHEMA IF EXISTS entain;
DROP SCHEMA IF EXISTS public;

CREATE SCHEMA IF NOT EXISTS entain;
ALTER SCHEMA entain OWNER TO postgres;
