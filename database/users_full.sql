--
-- PostgreSQL database dump
--

\restrict NTrSE107M5g0wM3KLZCLhbbSyiLoTPR1ZsQbANoTogPNQcD8aPr5Q4QhUf1Ls0O

-- Dumped from database version 18.1
-- Dumped by pg_dump version 18.1

-- Started on 2026-01-29 12:12:25

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 220 (class 1259 OID 16487)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    email text NOT NULL,
    google_uid text,
    password text,
    role public.role DEFAULT 'Owner'::public.role NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 5058 (class 0 OID 16487)
-- Dependencies: 220
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, email, google_uid, password, role, created_at, updated_at) VALUES ('f2685ec7-e3ee-4be2-9c53-93f9c15e37b1', 'owner@laundry.com', NULL, '$2a$10$7qybndyMPsq4V2mTNsCrMuXVf/cd1KwWgWaKx.QVFuoFhmj3AAZka', 'Owner', '2026-01-25 19:09:29.150618+07', '2026-01-25 19:09:29.150618+07');
INSERT INTO public.users (id, email, google_uid, password, role, created_at, updated_at) VALUES ('f5e0ff3d-2868-48ec-b228-7b2f5d49ef2b', 'tedymenong@gmail.com', '100154710522124386299', NULL, 'Cashier', '2026-01-29 02:28:21.749529+07', '2026-01-29 02:28:21.749529+07');


--
-- TOC entry 4905 (class 2606 OID 16502)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 4907 (class 2606 OID 16504)
-- Name: users users_google_uid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_google_uid_key UNIQUE (google_uid);


--
-- TOC entry 4909 (class 2606 OID 16500)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4902 (class 1259 OID 16520)
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_email ON public.users USING btree (email);


--
-- TOC entry 4903 (class 1259 OID 16521)
-- Name: idx_users_google_uid; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_google_uid ON public.users USING btree (google_uid);


--
-- TOC entry 4910 (class 2620 OID 16523)
-- Name: users trg_users_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trg_users_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


-- Completed on 2026-01-29 12:12:25

--
-- PostgreSQL database dump complete
--

\unrestrict NTrSE107M5g0wM3KLZCLhbbSyiLoTPR1ZsQbANoTogPNQcD8aPr5Q4QhUf1Ls0O

