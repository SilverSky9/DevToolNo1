--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

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
-- Name: option_product; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.option_product AS ENUM (
    'buy',
    'sell'
);


ALTER TYPE public.option_product OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: interested_post; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.interested_post (
    post_id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public.interested_post OWNER TO postgres;

--
-- Name: interested_post_post_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.interested_post_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.interested_post_post_id_seq OWNER TO postgres;

--
-- Name: interested_post_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.interested_post_post_id_seq OWNED BY public.interested_post.post_id;


--
-- Name: interested_post_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.interested_post_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.interested_post_user_id_seq OWNER TO postgres;

--
-- Name: interested_post_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.interested_post_user_id_seq OWNED BY public.interested_post.user_id;


--
-- Name: pin; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pin (
    user_id integer,
    pin_id integer NOT NULL,
    pin character varying(255)
);


ALTER TABLE public.pin OWNER TO postgres;

--
-- Name: pin_pin_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pin_pin_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pin_pin_id_seq OWNER TO postgres;

--
-- Name: pin_pin_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pin_pin_id_seq OWNED BY public.pin.pin_id;


--
-- Name: post; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.post (
    post_id integer NOT NULL,
    pin_id integer,
    product_name character varying(255),
    post_date date,
    product_option public.option_product,
    price integer,
    amount integer,
    tag_id integer
);


ALTER TABLE public.post OWNER TO postgres;

--
-- Name: post_post_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.post_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_post_id_seq OWNER TO postgres;

--
-- Name: post_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.post_post_id_seq OWNED BY public.post.post_id;


--
-- Name: tag; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tag (
    tag_id integer NOT NULL,
    tag_name character varying
);


ALTER TABLE public.tag OWNER TO postgres;

--
-- Name: tag_tag_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tag_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tag_tag_id_seq OWNER TO postgres;

--
-- Name: tag_tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tag_tag_id_seq OWNED BY public.tag.tag_id;


--
-- Name: user1; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user1 (
    user_id integer NOT NULL,
    name character varying(255),
    address character varying(255),
    contact character varying(255),
    phone_number character varying(10)
);


ALTER TABLE public.user1 OWNER TO postgres;

--
-- Name: user1_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user1_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user1_user_id_seq OWNER TO postgres;

--
-- Name: user1_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user1_user_id_seq OWNED BY public.user1.user_id;


--
-- Name: interested_post post_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interested_post ALTER COLUMN post_id SET DEFAULT nextval('public.interested_post_post_id_seq'::regclass);


--
-- Name: interested_post user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interested_post ALTER COLUMN user_id SET DEFAULT nextval('public.interested_post_user_id_seq'::regclass);


--
-- Name: pin pin_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pin ALTER COLUMN pin_id SET DEFAULT nextval('public.pin_pin_id_seq'::regclass);


--
-- Name: post post_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post ALTER COLUMN post_id SET DEFAULT nextval('public.post_post_id_seq'::regclass);


--
-- Name: tag tag_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag ALTER COLUMN tag_id SET DEFAULT nextval('public.tag_tag_id_seq'::regclass);


--
-- Name: user1 user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user1 ALTER COLUMN user_id SET DEFAULT nextval('public.user1_user_id_seq'::regclass);


--
-- Name: pin pin_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pin
    ADD CONSTRAINT pin_pkey PRIMARY KEY (pin_id);


--
-- Name: post post_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_pkey PRIMARY KEY (post_id);


--
-- Name: tag tag_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (tag_id);


--
-- Name: user1 user1_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user1
    ADD CONSTRAINT user1_pkey PRIMARY KEY (user_id);


--
-- Name: post fk_pin_post; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT fk_pin_post FOREIGN KEY (pin_id) REFERENCES public.pin(pin_id);


--
-- Name: pin pin_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pin
    ADD CONSTRAINT pin_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.user1(user_id);


--
-- Name: interested_post post_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interested_post
    ADD CONSTRAINT post_id FOREIGN KEY (post_id) REFERENCES public.post(post_id);


--
-- Name: interested_post user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.interested_post
    ADD CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES public.user1(user_id);


--
-- PostgreSQL database dump complete
--

