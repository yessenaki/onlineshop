--
-- PostgreSQL database dump
--

-- Dumped from database version 12.4
-- Dumped by pg_dump version 12.4

-- Started on 2020-10-20 22:04:21

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

DROP DATABASE onlineshop;
--
-- TOC entry 2970 (class 1262 OID 16393)
-- Name: onlineshop; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE onlineshop WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'English_United States.1252' LC_CTYPE = 'English_United States.1252';


ALTER DATABASE onlineshop OWNER TO postgres;

\connect onlineshop

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 209 (class 1259 OID 24742)
-- Name: brands; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.brands (
    id smallint NOT NULL,
    name character varying(30),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.brands OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 24740)
-- Name: brands_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.brands ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.brands_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 220 (class 1259 OID 32939)
-- Name: cart_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cart_items (
    id smallint NOT NULL,
    cart_id smallint,
    product_id smallint,
    quantity smallint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.cart_items OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 32944)
-- Name: cart_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.cart_items ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.cart_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 223 (class 1259 OID 32948)
-- Name: carts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.carts (
    id smallint NOT NULL,
    user_id smallint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.carts OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 32946)
-- Name: carts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.carts ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.carts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 206 (class 1259 OID 24732)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id smallint NOT NULL,
    name character varying(30),
    parent_id smallint DEFAULT 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    gender smallint DEFAULT 0,
    is_kids smallint DEFAULT 0
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 2971 (class 0 OID 0)
-- Dependencies: 206
-- Name: COLUMN categories.is_kids; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.categories.is_kids IS '0 - all, 1 - yes, 2 - no';


--
-- TOC entry 207 (class 1259 OID 24737)
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.categories ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 215 (class 1259 OID 24766)
-- Name: shoe_sizes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shoe_sizes (
    id smallint NOT NULL,
    size character varying(10),
    created_at time without time zone,
    updated_at timestamp without time zone,
    type smallint
);


ALTER TABLE public.shoe_sizes OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 24764)
-- Name: clothes_sizes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.shoe_sizes ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.clothes_sizes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 211 (class 1259 OID 24749)
-- Name: colors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.colors (
    id smallint NOT NULL,
    name character varying(30),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.colors OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 24747)
-- Name: colors_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.colors ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.colors_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 219 (class 1259 OID 32929)
-- Name: files; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.files (
    id smallint NOT NULL,
    name character varying(255),
    path character varying(255),
    product_id smallint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.files OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 32927)
-- Name: files_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.files ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.files_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 227 (class 1259 OID 32965)
-- Name: post_categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.post_categories (
    id smallint NOT NULL,
    name character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.post_categories OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 32955)
-- Name: posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.posts (
    id smallint NOT NULL,
    title character varying(255),
    body text,
    image_path character varying(255),
    category_id smallint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    image_name character varying(255),
    author character varying(255)
);


ALTER TABLE public.posts OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 32953)
-- Name: post_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.posts ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.post_categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 226 (class 1259 OID 32963)
-- Name: post_categories_id_seq1; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.post_categories ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.post_categories_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 231 (class 1259 OID 32979)
-- Name: post_tag_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.post_tag_items (
    id smallint NOT NULL,
    post_id smallint,
    tag_id smallint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.post_tag_items OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 32977)
-- Name: post_tag_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.post_tag_items ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.post_tag_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 229 (class 1259 OID 32972)
-- Name: post_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.post_tags (
    id smallint NOT NULL,
    name character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.post_tags OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 32970)
-- Name: post_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.post_tags ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.post_tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 217 (class 1259 OID 24784)
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id smallint NOT NULL,
    title character varying(50),
    price integer,
    old_price integer,
    gender smallint,
    is_kids smallint,
    is_new smallint,
    is_discount smallint,
    dsc_percent smallint,
    brand_id smallint,
    color_id smallint,
    category_id smallint,
    size_id smallint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    description text
);


ALTER TABLE public.products OWNER TO postgres;

--
-- TOC entry 2972 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN products.gender; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.products.gender IS '0 - male, 1 - female';


--
-- TOC entry 2973 (class 0 OID 0)
-- Dependencies: 217
-- Name: COLUMN products.is_kids; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.products.is_kids IS '0 - no, 1 - yes';


--
-- TOC entry 216 (class 1259 OID 24782)
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.products ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 203 (class 1259 OID 16498)
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id integer NOT NULL,
    session_id character varying(50) NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 16509)
-- Name: sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.sessions ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 213 (class 1259 OID 24759)
-- Name: sizes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sizes (
    id smallint NOT NULL,
    size character varying(10),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    type smallint
);


ALTER TABLE public.sizes OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 24757)
-- Name: shoe_sizes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.sizes ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.shoe_sizes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 202 (class 1259 OID 16394)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    first_name character varying(50),
    last_name character varying(50),
    email character varying(50),
    password character varying(80),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    role smallint DEFAULT 0
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 16511)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 2942 (class 0 OID 24742)
-- Dependencies: 209
-- Data for Name: brands; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (1, 'Whistles', '2020-09-18 15:44:03', '2020-10-02 18:57:49');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (2, 'Asos', '2020-09-18 16:28:39', '2020-10-02 18:58:14');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (4, 'Virgos Lounge', '2020-09-21 15:27:07', '2020-10-02 19:22:14');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (5, 'The North Face', '2020-10-02 19:41:16', '2020-10-02 19:41:16');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (6, 'Calvin Klein', '2020-10-02 19:47:21', '2020-10-02 19:47:21');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (7, 'Tommy Hilfiger', '2020-10-02 19:47:55', '2020-10-02 19:47:55');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (8, 'Dr Denim', '2020-10-02 20:01:34', '2020-10-02 20:01:34');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (9, 'JDY', '2020-10-02 20:16:24', '2020-10-02 20:16:24');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (10, 'Topman', '2020-10-02 20:19:59', '2020-10-02 20:19:59');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (11, 'Celio', '2020-10-05 10:41:56', '2020-10-05 10:41:56');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (12, 'Closet London', '2020-10-05 10:45:30', '2020-10-05 10:45:30');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (13, 'Lamoda', '2020-10-05 11:09:13', '2020-10-05 11:09:13');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (14, 'Public Desire', '2020-10-05 11:18:19', '2020-10-05 11:18:19');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (15, 'Hudson London', '2020-10-05 11:22:48', '2020-10-05 11:22:48');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (16, 'UGG', '2020-10-05 11:33:59', '2020-10-05 11:33:59');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (17, 'Nike', '2020-10-05 11:41:10', '2020-10-05 11:41:10');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (18, 'River Island', '2020-10-05 11:45:22', '2020-10-05 11:45:22');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (19, 'Hugo', '2020-10-05 12:27:40', '2020-10-05 12:27:40');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (20, '& Other Stories', '2020-10-05 12:31:08', '2020-10-05 12:31:27');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (21, 'St. Tropez', '2020-10-05 12:36:05', '2020-10-05 12:36:05');
INSERT INTO public.brands (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (22, 'Essie', '2020-10-05 12:40:01', '2020-10-05 12:40:01');


--
-- TOC entry 2953 (class 0 OID 32939)
-- Dependencies: 220
-- Data for Name: cart_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.cart_items (id, cart_id, product_id, quantity, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (3, 2, 0, 1, '2020-10-11 10:52:39', '2020-10-11 10:52:39');


--
-- TOC entry 2956 (class 0 OID 32948)
-- Dependencies: 223
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.carts (id, user_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (1, 1, '2020-10-10 21:03:29', '2020-10-10 21:03:29');
INSERT INTO public.carts (id, user_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (2, 0, '2020-10-11 10:52:39', '2020-10-11 10:52:39');


--
-- TOC entry 2939 (class 0 OID 24732)
-- Dependencies: 206
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (2, 'Shoes', 0, '2020-09-16 19:59:16', '2020-10-05 17:32:48', 2, 2);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (1, 'Clothing', 0, '2020-09-12 16:20:16', '2020-10-05 17:32:55', 2, 2);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (4, 'Cosmetics', 0, '2020-09-16 20:00:43', '2020-10-05 17:33:34', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (3, 'Accessories', 0, '2020-09-16 19:59:52', '2020-10-05 17:33:41', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (5, 'Coats & Jackets', 1, '2020-09-16 20:03:09', '2020-10-05 17:34:02', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (6, 'Dresses', 1, '2020-09-16 20:03:35', '2020-10-05 17:34:21', 1, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (7, 'Hoodies & Sweatshirts', 1, '2020-09-17 21:51:11', '2020-10-05 17:34:33', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (9, 'Jeans', 1, '2020-09-21 12:10:20', '2020-10-05 17:34:46', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (10, 'Jumpers & Cardigans', 1, '2020-09-21 12:31:59', '2020-10-05 17:35:06', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (12, 'Skirts', 1, '2020-09-26 12:00:40', '2020-10-05 17:35:25', 1, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (13, 'Polo shirts', 1, '2020-09-26 12:04:02', '2020-10-05 17:35:46', 0, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (15, 'T-Shirts & Vests', 1, '2020-09-26 12:04:54', '2020-10-05 17:36:01', 0, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (16, 'Boots', 2, '2020-09-26 12:06:51', '2020-10-05 17:36:13', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (17, 'Heels', 2, '2020-09-26 12:07:38', '2020-10-05 17:36:27', 1, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (18, 'Sandals', 2, '2020-09-26 12:07:52', '2020-10-05 17:36:37', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (19, 'Slippers', 2, '2020-09-26 12:08:26', '2020-10-05 17:36:59', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (20, 'Trainers', 2, '2020-09-26 12:08:46', '2020-10-05 17:37:05', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (21, 'Bags & Purses', 3, '2020-09-26 12:09:46', '2020-10-05 17:37:13', 1, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (22, 'Belts', 3, '2020-09-26 12:10:04', '2020-10-05 17:37:23', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (23, 'Caps & Hats', 3, '2020-09-26 12:10:36', '2020-10-05 17:37:35', 2, 2);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (30, 'Wallets', 3, '2020-09-26 12:13:34', '2020-10-05 17:37:53', 0, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (31, 'Body care', 4, '2020-09-26 12:14:24', '2020-10-05 17:38:02', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (32, 'Skin care', 4, '2020-09-26 12:14:49', '2020-10-05 17:38:14', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (33, 'Hair care', 4, '2020-09-26 12:15:17', '2020-10-05 17:38:30', 2, 0);
INSERT INTO public.categories (id, name, parent_id, created_at, updated_at, gender, is_kids) OVERRIDING SYSTEM VALUE VALUES (34, 'Nails', 4, '2020-09-26 12:15:43', '2020-10-05 17:38:42', 1, 0);


--
-- TOC entry 2944 (class 0 OID 24749)
-- Dependencies: 211
-- Data for Name: colors; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (1, 'Beige', '2020-09-18 21:53:10', '2020-09-18 21:53:10');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (2, 'Black', '2020-09-18 21:53:16', '2020-09-18 21:53:16');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (3, 'Blue', '2020-09-18 21:53:26', '2020-09-18 21:53:26');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (4, 'Brown', '2020-09-18 21:53:33', '2020-09-18 21:53:33');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (5, 'Copper', '2020-09-18 21:53:49', '2020-09-18 21:53:49');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (6, 'Cream', '2020-09-18 21:53:54', '2020-09-18 21:53:54');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (7, 'Gold', '2020-09-18 21:54:08', '2020-09-18 21:54:08');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (8, 'Green', '2020-09-18 21:54:18', '2020-09-18 21:54:18');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (9, 'Grey', '2020-09-18 21:54:25', '2020-09-18 21:54:25');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (10, 'Multi', '2020-09-18 21:55:28', '2020-09-18 21:55:28');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (11, 'Navy', '2020-09-18 21:55:34', '2020-09-18 21:55:34');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (12, 'Orange', '2020-09-18 21:55:50', '2020-09-18 21:55:50');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (13, 'Pink', '2020-09-18 21:56:02', '2020-09-18 21:56:02');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (14, 'Purple', '2020-09-18 21:56:09', '2020-09-18 21:56:09');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (15, 'Red', '2020-09-18 21:56:22', '2020-09-18 21:56:22');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (16, 'Silver', '2020-09-18 21:56:29', '2020-09-18 21:56:29');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (17, 'Stone', '2020-09-18 21:56:44', '2020-09-18 21:56:44');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (18, 'Tan', '2020-09-18 21:57:24', '2020-09-18 21:57:24');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (19, 'White', '2020-09-18 21:57:41', '2020-09-18 21:57:41');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (20, 'Yellow', '2020-09-18 21:57:47', '2020-09-18 21:57:47');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (21, 'Cyan', '2020-09-21 16:12:48', '2020-09-21 16:13:06');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (22, 'Khaki', '2020-10-02 19:15:48', '2020-10-02 19:15:48');
INSERT INTO public.colors (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (23, 'Oyster', '2020-10-05 11:31:29', '2020-10-05 11:31:29');


--
-- TOC entry 2952 (class 0 OID 32929)
-- Dependencies: 219
-- Data for Name: files; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (6, '20609979-1-black.jpg', '/assets/uploads/ce7974e5049b9cc3275738ebb7460a6df11be50a6ef9f24e9555462917cb48d9.jpg', 4, '2020-10-09 19:05:38', '2020-10-09 19:05:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (7, '20609979-2.jpg', '/assets/uploads/374d26b49e760f807c14b6dbde83856feb9dc10f10d53ce257e8c0454fcdab6f.jpg', 4, '2020-10-09 19:05:38', '2020-10-09 19:05:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (8, '20609979-3.jpg', '/assets/uploads/657d97ed0cbe1fc12867e0beab6f421771e522b160867bbb1c82c12cb51b1dab.jpg', 4, '2020-10-09 19:05:38', '2020-10-09 19:05:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (9, '20609979-4.jpg', '/assets/uploads/41e70c952e2c620ba91dee34c2134a141763edec43b8208c676c79c208e28a4a.jpg', 4, '2020-10-09 19:05:38', '2020-10-09 19:05:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (10, '14695701-1-brown.jpg', '/assets/uploads/6b8715d2ab6331b00aa4b030b7574096eb8e8deaa3320169eb19bd50de5089aa.jpg', 5, '2020-10-09 19:09:08', '2020-10-09 19:09:08');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (11, '14695701-2.jpg', '/assets/uploads/7e2b3b3d44877538e981aa32ec27fbd19d095327b2b343ce64ba0ae8c3496001.jpg', 5, '2020-10-09 19:09:08', '2020-10-09 19:09:08');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (12, '14695701-3.jpg', '/assets/uploads/dd0441d20c954982d96f7e1660efc6d41e0b2a36470e657ce8ea9403110f5258.jpg', 5, '2020-10-09 19:09:08', '2020-10-09 19:09:08');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (13, '14695701-4.jpg', '/assets/uploads/66d25d18f97c1ec6edc5644fe4ce374952c996772879fd099902b7cfaa3a6184.jpg', 5, '2020-10-09 19:09:08', '2020-10-09 19:09:08');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (14, '13201055-1-black.jpg', '/assets/uploads/0ef9fb17bd827e6b6834602d4b2578a4b2b9e7454b096aacb97ef7540c5991ae.jpg', 6, '2020-10-09 19:11:13', '2020-10-09 19:11:13');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (15, '13201055-2.jpg', '/assets/uploads/4df6e250369ba4814f6c439a1a8431c361ee6c0767bda09eb70210389250ffea.jpg', 6, '2020-10-09 19:11:13', '2020-10-09 19:11:13');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (16, '13201055-3.jpg', '/assets/uploads/b4d299d8eb2cd7d770ecd46e392d95cbe03b59ba81dcd74c9841bf8d895d22af.jpg', 6, '2020-10-09 19:11:13', '2020-10-09 19:11:13');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (17, '13201055-4.jpg', '/assets/uploads/56a3ab23fad20bc54b1ed3cb692bfae2597c68cdcd3ec8adfa2efa61375f533a.jpg', 6, '2020-10-09 19:11:13', '2020-10-09 19:11:13');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (18, '20508110-1-khaki.jpg', '/assets/uploads/165af94aad688f8b33174451deca6fe20621312c0ead90cc88fb3f54593c4a0b.jpg', 7, '2020-10-09 19:11:26', '2020-10-09 19:11:26');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (19, '20508110-2.jpg', '/assets/uploads/7bf8485d0a4461640e5b95556be955d0182b7e8875ed876f93e4937fd61e46cf.jpg', 7, '2020-10-09 19:11:26', '2020-10-09 19:11:26');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (20, '20508110-3.jpg', '/assets/uploads/57d79ca9804872a1ed2752f8d945c706dee0909aad670619dad7702a81212806.jpg', 7, '2020-10-09 19:11:26', '2020-10-09 19:11:26');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (21, '20508110-4.jpg', '/assets/uploads/2e786cdff3368bb3ed755b137ff4ca0a30398e9f11930a11dc4b357adfe62418.jpg', 7, '2020-10-09 19:11:26', '2020-10-09 19:11:26');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (22, '14941769-1-navy.jpg', '/assets/uploads/0a2e99a1665b235b6ec800394cf5f31a866e514355ecf0074fe07f06938570e2.jpg', 8, '2020-10-09 19:12:01', '2020-10-09 19:12:01');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (23, '14941769-2.jpg', '/assets/uploads/526f22222217234db14892f080e3d25668ba5bed1eff8b11fb0b8a76e3c5a07c.jpg', 8, '2020-10-09 19:12:01', '2020-10-09 19:12:01');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (24, '14941769-3.jpg', '/assets/uploads/93fc5980a1338bd6e3503eed9cb2e483043223dd06a8968f0105828d58ea604b.jpg', 8, '2020-10-09 19:12:01', '2020-10-09 19:12:01');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (25, '14941769-4.jpg', '/assets/uploads/025f8fd65f02bb7e726cc8bf069f1a7008d653b99873f6fa4e939d9709fd617a.jpg', 8, '2020-10-09 19:12:01', '2020-10-09 19:12:01');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (26, '20821575-1-green.jpg', '/assets/uploads/28ce9e957b8ae0b431084105963f47c3553dc9e4409c65157e447ed5e1d3c9e5.jpg', 9, '2020-10-09 19:12:50', '2020-10-09 19:12:50');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (27, '20821575-2.jpg', '/assets/uploads/0938cce0be4a97a2ede2d6c33a5cd505110759dd349fba25e73f469b55eb39f2.jpg', 9, '2020-10-09 19:12:50', '2020-10-09 19:12:50');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (28, '20821575-3.jpg', '/assets/uploads/7873a4a56bb3411ca1afb9fc0217b3906b0c6e66db0a06de359ce80d1f749a02.jpg', 9, '2020-10-09 19:12:50', '2020-10-09 19:12:50');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (29, '20821575-4.jpg', '/assets/uploads/177296fd7081c183e8610ad747afd489f2695f2b6bc6454ef8d8ea9e67aa5243.jpg', 9, '2020-10-09 19:12:50', '2020-10-09 19:12:50');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (30, '21571512-1-pinkclay.jpg', '/assets/uploads/8f57e72499c179c9715d05d6bc2e8f735f9207df55dda48a781b38c775918aef.jpg', 10, '2020-10-09 19:13:39', '2020-10-09 19:13:39');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (31, '21571512-2.jpg', '/assets/uploads/7a4c719685fd0335ff9b015a310f358fa6caeaf16199a515a146fda745dac999.jpg', 10, '2020-10-09 19:13:39', '2020-10-09 19:13:39');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (32, '21571512-3.jpg', '/assets/uploads/dc9aa6471924a111e39a955af02a15493b70594098d9e97f4dfd2fd2362f47bb.jpg', 10, '2020-10-09 19:13:39', '2020-10-09 19:13:39');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (33, '21571512-4.jpg', '/assets/uploads/8a96ab99fff40d0b3e13699630f3d8f28b096d1793045c4d3538f606ebb5d001.jpg', 10, '2020-10-09 19:13:39', '2020-10-09 19:13:39');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (34, '21346198-1-red.jpg', '/assets/uploads/5a8c6462f0c26be118c5b65ec6f10a445b4d96ab74c6a5d765e389a36bc97c49.jpg', 11, '2020-10-09 19:13:58', '2020-10-09 19:13:58');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (35, '21346198-2.jpg', '/assets/uploads/a6f2e855d41bad1968f1fad81ce9327e1b9c0d73e1d0a19ed96c5a80af60801d.jpg', 11, '2020-10-09 19:13:58', '2020-10-09 19:13:58');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (36, '21346198-3.jpg', '/assets/uploads/6f21ba884e79e283f3570d1e34887c9cce66e5ab677f74cb09a2ab6f5b2c9cc5.jpg', 11, '2020-10-09 19:13:58', '2020-10-09 19:13:58');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (37, '21346198-4.jpg', '/assets/uploads/687e3a46a9b081b827bbb518b58e957b70860d3bb9925ebb7a5be1382f2bd207.jpg', 11, '2020-10-09 19:13:58', '2020-10-09 19:13:58');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (38, '12992309-1-black.jpg', '/assets/uploads/0cd6e7cdc96cae16c35353e0523b986e2660dc28849206687eeda2b33e675fbc.jpg', 12, '2020-10-09 19:14:40', '2020-10-09 19:14:40');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (39, '12992309-2.jpg', '/assets/uploads/64c703b00279b9f89eb158398ee6d9a77ca956c4a68ce12a28179885d28c8af1.jpg', 12, '2020-10-09 19:14:40', '2020-10-09 19:14:40');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (40, '12992309-3.jpg', '/assets/uploads/3f1ed70f377f62f56fa48015b7b7ad9a53ffda8659b751cba700b659c2bff6c1.jpg', 12, '2020-10-09 19:14:40', '2020-10-09 19:14:40');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (41, '12992309-4.jpg', '/assets/uploads/64526a8e07de6c4eff4f712fb1b8ae176dcad58a62d54cd1566a4c3f424fe64e.jpg', 12, '2020-10-09 19:14:40', '2020-10-09 19:14:40');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (42, '14665124-1-green.jpg', '/assets/uploads/ca837912f3dbf82dc74dd603b1cd5b7bad3a70de17fa64ca1c66db8a792cd109.jpg', 13, '2020-10-09 19:14:59', '2020-10-09 19:14:59');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (43, '14665124-2.jpg', '/assets/uploads/804979965a01100f70c1a906f48c67d7da4fdba65ce98d1464c1b86f1e969e43.jpg', 13, '2020-10-09 19:14:59', '2020-10-09 19:14:59');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (44, '14665124-3.jpg', '/assets/uploads/8c2d59c41e59b469a85dd5b6de2f22af36d2c498aafede6079f347734c28dfd5.jpg', 13, '2020-10-09 19:14:59', '2020-10-09 19:14:59');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (45, '14665124-4.jpg', '/assets/uploads/55ccd63622f61069a03c619f85ccb53cfd560f62407b1c7469fa05cf6fb78caf.jpg', 13, '2020-10-09 19:14:59', '2020-10-09 19:14:59');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (50, '12347161-1-midwash.jpg', '/assets/uploads/3f718aa0172a34bc2650b2ae5806e4a370ff929b83ba35256ab15bf05c05808a.jpg', 15, '2020-10-09 19:16:34', '2020-10-09 19:16:34');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (51, '12347161-2.jpg', '/assets/uploads/18791922b5d092ee4a0af6be326864528daa75f4a0973aec231d9747241f6c66.jpg', 15, '2020-10-09 19:16:34', '2020-10-09 19:16:34');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (52, '12347161-3.jpg', '/assets/uploads/ee99524ab47edb76dfe85a0cca8781909e13cc84da113b1949bd718a8d505a27.jpg', 15, '2020-10-09 19:16:34', '2020-10-09 19:16:34');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (53, '12347161-4.jpg', '/assets/uploads/e93b6fa596208aa1fb60bdbbc9e4533b4253feedc4cc2c084484cdfe49109f1f.jpg', 15, '2020-10-09 19:16:34', '2020-10-09 19:16:34');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (54, '21207330-1-bluejay.jpg', '/assets/uploads/a537fa220f13aae99fa4f63b0c65089fea7298ff796f29d4179a802d3f7624d8.jpg', 14, '2020-10-09 19:17:06', '2020-10-09 19:17:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (55, '21207330-2.jpg', '/assets/uploads/bcfdf3eb2cb3b93014f739fcb6ec4a77421a84c6fedb6116bd9b448a1c66dafe.jpg', 14, '2020-10-09 19:17:06', '2020-10-09 19:17:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (56, '21207330-3.jpg', '/assets/uploads/67c6270e06914e0d458bfdd6dbeda51987af59cbcf0b30f2852f3b46028347df.jpg', 14, '2020-10-09 19:17:06', '2020-10-09 19:17:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (57, '21207330-4.jpg', '/assets/uploads/6cb52c636594c5d479b87dcdfa17070c1859d983fa1706ec84f6d6fee9dddedc.jpg', 14, '2020-10-09 19:17:06', '2020-10-09 19:17:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (58, '20248978-1-ab076icnmidblue.jpg', '/assets/uploads/f7a532d836360822045bf4fe07827db61ac7a89daefb7c1fcdac695aa1fd4e21.jpg', 16, '2020-10-09 19:17:31', '2020-10-09 19:17:31');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (59, '20248978-2.jpg', '/assets/uploads/b42000b27379d7a4664720fdf529a7de66d19a0472e713b811479d07b9623fe0.jpg', 16, '2020-10-09 19:17:31', '2020-10-09 19:17:31');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (60, '20248978-3.jpg', '/assets/uploads/cb9ec3da16ca061abde1bde148a2ba086565d917bc43f10f9146a814d9ad4a8c.jpg', 16, '2020-10-09 19:17:31', '2020-10-09 19:17:31');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (61, '20248978-4.jpg', '/assets/uploads/6aae56e94616766bcd04631b36f5d256673900310350f3f5074bda84bbf9334a.jpg', 16, '2020-10-09 19:17:31', '2020-10-09 19:17:31');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (62, '21896425-1-grey.jpg', '/assets/uploads/b16763f4363b01f406aa456666ac415135a368250386ccf7ebc3c1b7a58a3350.jpg', 17, '2020-10-09 19:18:06', '2020-10-09 19:18:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (64, '21896425-3.jpg', '/assets/uploads/4fe0d4e922a9866cd84d6edee51cb88bbdc7957d3be6ad47f237ba789f6f70ad.jpg', 17, '2020-10-09 19:18:06', '2020-10-09 19:18:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (66, '20569957-1-pink.jpg', '/assets/uploads/68e7c2a31959f15670718c1a57baf1de97665046e21fb0c378c60992d5185992.jpg', 18, '2020-10-09 19:18:30', '2020-10-09 19:18:30');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (68, '20569957-3.jpg', '/assets/uploads/b4e885f7164d31fcdf287f9d9583697b9a8c449ececfb6754e53a7ac51efd2f1.jpg', 18, '2020-10-09 19:18:30', '2020-10-09 19:18:30');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (70, '14488823-1-greymel.jpg', '/assets/uploads/8bf0e27bfb3530963c5afbaee523f9a5d43f5b0412921d871d58d8e19a69395f.jpg', 19, '2020-10-09 19:18:58', '2020-10-09 19:18:58');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (72, '14488823-3.jpg', '/assets/uploads/c40f0e82d314cf16e104798937ffd87f1723591765d19d59c5a9be7f03ca6cb8.jpg', 19, '2020-10-09 19:18:58', '2020-10-09 19:18:58');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (74, '21800790-1-softbrown.jpg', '/assets/uploads/54d0e27aeaa3c63b85cff2826200ee17f6656dcec1673ca2b04481223da8d636.jpg', 20, '2020-10-09 19:19:19', '2020-10-09 19:19:19');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (76, '21800790-3.jpg', '/assets/uploads/b976fdc463f959a544eed5707e9084c5a5f7d4de65b7c649a7cf2392beb1ce5c.jpg', 20, '2020-10-09 19:19:19', '2020-10-09 19:19:19');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (78, '20620831-1-white.jpg', '/assets/uploads/c83f024c711c93ae1f410e0c78c0b9c314ebc102bb988d634bb0528ada299dcd.jpg', 21, '2020-10-09 19:19:40', '2020-10-09 19:19:40');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (80, '20620831-3.jpg', '/assets/uploads/2cd87aeea8eab30cd5d39800761818678e4bf88bd31b03740495ce333b62ca87.jpg', 21, '2020-10-09 19:19:40', '2020-10-09 19:19:40');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (82, '20610285-1-black.jpg', '/assets/uploads/bbe37299bbd29b4d75caef9f0ee12aaa1afd68429053b7798f58a83489120539.jpg', 22, '2020-10-09 19:20:06', '2020-10-09 19:20:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (84, '20610285-3.jpg', '/assets/uploads/094cafd512b7c8201496cd7bf5595945dd50bddfce8e9c784c12cbb08c33943e.jpg', 22, '2020-10-09 19:20:06', '2020-10-09 19:20:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (86, '21009333-1-beige.jpg', '/assets/uploads/e3cfda7d385530ccfd4cc66339477f5578667884d2211e56f8eaee8cbac91d28.jpg', 23, '2020-10-09 19:20:38', '2020-10-09 19:20:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (88, '21009333-3.jpg', '/assets/uploads/6bf69fcb9d67e445d4761e45dc72b324a0f5a602cbc939cf34927987035196f4.jpg', 23, '2020-10-09 19:20:38', '2020-10-09 19:20:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (90, '20597160-1-brown.jpg', '/assets/uploads/d7ef80bfa95e99e3cdbe2e220b738963c42a516deea915dd1053e0a040d048ca.jpg', 24, '2020-10-09 19:22:04', '2020-10-09 19:22:04');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (92, '20597160-3.jpg', '/assets/uploads/d592a60373e4ca6e33d8e601c8db27e6bd2b688870e4b4e8bd9a9ed13ca39bdc.jpg', 24, '2020-10-09 19:22:04', '2020-10-09 19:22:04');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (94, '14078261-1-black.jpg', '/assets/uploads/86189be73b136d5b61b3138d3ef124620872d11e71d5203e5228830fff97f9c5.jpg', 25, '2020-10-09 19:22:43', '2020-10-09 19:22:43');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (96, '14078261-3.jpg', '/assets/uploads/a5d02f8c6c4cd258ea1717a4598f89e59cb1c822e4fd71de1212ed674cd809ef.jpg', 25, '2020-10-09 19:22:43', '2020-10-09 19:22:43');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (98, '11099693-1-oyster.jpg', '/assets/uploads/16a35254491140e7ad91000e0313fcf00f2fab07d0fe781785d2de7213e48271.jpg', 26, '2020-10-09 19:23:10', '2020-10-09 19:23:10');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (100, '11099693-3.jpg', '/assets/uploads/aa6c0fff0756780f6278184df136854e4469ed6ded93a80e640c3e197a89b711.jpg', 26, '2020-10-09 19:23:10', '2020-10-09 19:23:10');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (102, '21135052-1-blackchileredblea.jpg', '/assets/uploads/7a067e3f12534b195a3347f523d2f31733ef303765d2ad4b144fcd152486fb4d.jpg', 27, '2020-10-09 19:23:41', '2020-10-09 19:23:41');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (104, '21135052-3.jpg', '/assets/uploads/baf35f9802e0c37fa2ac30c41c65057a7e0d4fb0196ac3ca897958b67c14855a.jpg', 27, '2020-10-09 19:23:41', '2020-10-09 19:23:41');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (106, '21095805-1-grey.jpg', '/assets/uploads/a4f8ad2c6bb3f0bf8bdcbddfe7dc2e20622268d8491d6baabf74e70d63ef17ac.jpg', 28, '2020-10-09 19:24:17', '2020-10-09 19:24:17');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (108, '21095805-3.jpg', '/assets/uploads/3ba19ec82c8ad7ff65554068956646b8fa156a1a915574855cd39d60cdbeed91.jpg', 28, '2020-10-09 19:24:17', '2020-10-09 19:24:17');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (110, '20542825-1-black.jpg', '/assets/uploads/ce74f3ecff454167a5edf8ec04f1ba6a541a4015af303aef154aa26d0faf8b3a.jpg', 29, '2020-10-09 19:24:49', '2020-10-09 19:24:49');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (112, '20542825-3.jpg', '/assets/uploads/b96832a5cbfa684975a82692470b6c342810ec82768f5668cb38e2601fb73c4a.jpg', 29, '2020-10-09 19:24:49', '2020-10-09 19:24:49');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (63, '21896425-2.jpg', '/assets/uploads/87a44b830e05bf0d58845a669a380a5556dbf89f035eae330bc96da9694093ed.jpg', 17, '2020-10-09 19:18:06', '2020-10-09 19:18:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (65, '21896425-4.jpg', '/assets/uploads/835820fb69561918e5f99766a2b447ab6b3dcfaa70d513cea50cd2d9670fe112.jpg', 17, '2020-10-09 19:18:06', '2020-10-09 19:18:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (67, '20569957-2.jpg', '/assets/uploads/def1fe9d0b7b572a601e7c567fb05575df11c0f4d1de01b658018384a784839a.jpg', 18, '2020-10-09 19:18:30', '2020-10-09 19:18:30');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (69, '20569957-4.jpg', '/assets/uploads/cf692a02532d5c04a92dacc10235326531f32bcdc4fcd84b799627be0d0fa073.jpg', 18, '2020-10-09 19:18:30', '2020-10-09 19:18:30');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (71, '14488823-2.jpg', '/assets/uploads/d8922a05fa6455641e76b394b0bb1fc091d06b67c086d7ed251b1b3e14d273ce.jpg', 19, '2020-10-09 19:18:58', '2020-10-09 19:18:58');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (73, '14488823-4.jpg', '/assets/uploads/fbfec9dbc7fd07e9cf33995966840712904485983564fa134ed85be539e95a0f.jpg', 19, '2020-10-09 19:18:58', '2020-10-09 19:18:58');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (75, '21800790-2.jpg', '/assets/uploads/8ad8e58f6f0d40cac88c7bbd45c73d6eeb432d912eb6cc288709cd56d4175242.jpg', 20, '2020-10-09 19:19:19', '2020-10-09 19:19:19');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (77, '21800790-4.jpg', '/assets/uploads/78c94ae3bd06764290f4f652aae8e1dc07f10daf878df85f605975676d99d4f7.jpg', 20, '2020-10-09 19:19:19', '2020-10-09 19:19:19');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (79, '20620831-2.jpg', '/assets/uploads/36493f4e1bbb5c02fbbb47d45565940ee9bea54f6c09ff24431248c68c520eeb.jpg', 21, '2020-10-09 19:19:40', '2020-10-09 19:19:40');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (81, '20620831-4.jpg', '/assets/uploads/8343ccc7340ebc0a128136b2f569d4057f0d4dbe9fe75e5ca57e41706fc526c2.jpg', 21, '2020-10-09 19:19:40', '2020-10-09 19:19:40');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (83, '20610285-2.jpg', '/assets/uploads/deea1c3ddddf813f5dc26ec29d0f3a22b2fae1dd23cb862a5ac8c13c3fd6916f.jpg', 22, '2020-10-09 19:20:06', '2020-10-09 19:20:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (85, '20610285-4.jpg', '/assets/uploads/04a028d1215807dd3e0c7856890df4e72e02adcbd75e3b4cf41b98dd079bb5b9.jpg', 22, '2020-10-09 19:20:06', '2020-10-09 19:20:06');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (87, '21009333-2.jpg', '/assets/uploads/ab3aff90af19c3c0379ee96e453038aeb47a39f2a5910b3585602f3eb85a2b43.jpg', 23, '2020-10-09 19:20:38', '2020-10-09 19:20:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (89, '21009333-4.jpg', '/assets/uploads/097ffa2e7f3934e7987748c6e86028a4cd051e33dd169059a910b9985f2ff826.jpg', 23, '2020-10-09 19:20:38', '2020-10-09 19:20:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (91, '20597160-2.jpg', '/assets/uploads/1cbca20b04e34481b1e5e1a74936e733efb6d254b7e0e66076209f26cece4bcb.jpg', 24, '2020-10-09 19:22:04', '2020-10-09 19:22:04');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (93, '20620831-4.jpg', '/assets/uploads/2c762ec89697a9a275f1f4abe42d749ab4b3d764b04720cda003210e212b27d8.jpg', 24, '2020-10-09 19:22:04', '2020-10-09 19:22:04');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (95, '14078261-2.jpg', '/assets/uploads/f34978c447860b7b86761b2c274c35eb368cbf5d84709c6d266fa4b91ebd54c6.jpg', 25, '2020-10-09 19:22:43', '2020-10-09 19:22:43');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (97, '21009333-4.jpg', '/assets/uploads/0a98b621ad5e42818490b673b7b55f3a35984b419a6c848c034a24870b046679.jpg', 25, '2020-10-09 19:22:43', '2020-10-09 19:22:43');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (99, '11099693-2.jpg', '/assets/uploads/91e5bcaa12af80b721280c5ecb36d17455cf9529ad74b8d6a8ccd792c2e1654d.jpg', 26, '2020-10-09 19:23:10', '2020-10-09 19:23:10');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (101, '11099693-4.jpg', '/assets/uploads/967a8a985fd7f90dfc055b40debb7c9f2cec399bd11cdf1e79a9c2385bc27c86.jpg', 26, '2020-10-09 19:23:10', '2020-10-09 19:23:10');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (103, '21135052-2.jpg', '/assets/uploads/5e5e02ff6f8ece4039163aa9e6ccf3513af03bd1a659934f30007f92e9dbc7e5.jpg', 27, '2020-10-09 19:23:41', '2020-10-09 19:23:41');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (105, '21135052-4.jpg', '/assets/uploads/3c85b53019101135218f296357e566d47c78424ca12fc14d2b26f0dd8d61d1ef.jpg', 27, '2020-10-09 19:23:41', '2020-10-09 19:23:41');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (107, '21095805-2.jpg', '/assets/uploads/f082efcbaa9fa11fcd1bad12c0ac68c19ae6fc51b933565d353176edf3126f60.jpg', 28, '2020-10-09 19:24:17', '2020-10-09 19:24:17');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (109, '21095805-4.jpg', '/assets/uploads/680b35872b8ea0ab1d161acecde6a6db596b7adc02e90f5087473eed3030f0d0.jpg', 28, '2020-10-09 19:24:17', '2020-10-09 19:24:17');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (111, '20542825-2.jpg', '/assets/uploads/7ac3edcc9b37ad7479f383a5a03f526c70928e83e647a7ef6cb3e0e3fa195d7e.jpg', 29, '2020-10-09 19:24:49', '2020-10-09 19:24:49');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (113, '20542825-4.jpg', '/assets/uploads/ff3c7369794a3dbe141a6877043fa030626fabff1befcdd810acb1bfda878be9.jpg', 29, '2020-10-09 19:24:49', '2020-10-09 19:24:49');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (114, '21822552-1-brown.jpg', '/assets/uploads/254e918cbae6443e75c158c3357ec425283f6f8adb35ec6536e7ed429840b3ad.jpg', 30, '2020-10-09 19:25:20', '2020-10-09 19:25:20');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (115, '21822552-2.jpg', '/assets/uploads/a17049d5757761b84390ea5257035841d0675411614e1cabdc24a89dd0b050f5.jpg', 30, '2020-10-09 19:25:20', '2020-10-09 19:25:20');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (116, '21822552-3.jpg', '/assets/uploads/54407dbd32a535b0793147a0e23db1bd7983962c4e4a5325eb18626c82407ba1.jpg', 30, '2020-10-09 19:25:20', '2020-10-09 19:25:20');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (117, '21822552-4.jpg', '/assets/uploads/af255cb7d651b07cb7a337f937d81422b499443cbb80c08a7adbca9aa2e0a36a.jpg', 30, '2020-10-09 19:25:20', '2020-10-09 19:25:20');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (118, '22132062-1-black.jpg', '/assets/uploads/936de87ff2d9727123acc549caf24b13891a3b682aa67e1131e0ec329b0595c9.jpg', 31, '2020-10-09 19:25:38', '2020-10-09 19:25:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (119, '22132062-2.jpg', '/assets/uploads/57a6904219bb2d2f013f8bd6f9ff0ac9e1351f26cc7275b5441fb1069a9fbc99.jpg', 31, '2020-10-09 19:25:38', '2020-10-09 19:25:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (120, '22132062-3.jpg', '/assets/uploads/98492915d8207f520156434b457f0dab7ab33fb08181fc88576332e47620b98d.jpg', 31, '2020-10-09 19:25:38', '2020-10-09 19:25:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (121, '22132062-4.jpg', '/assets/uploads/d841cee7aa72bde5af1c68a12f8edfc11fdb57eb38389d185919c7aa9f9303be.jpg', 31, '2020-10-09 19:25:38', '2020-10-09 19:25:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (122, '8445758-1-black.jpg', '/assets/uploads/93ab58255125fb72e7687bbc50442a4f18fc10d18d9d592c696c7f9e7644bf91.jpg', 32, '2020-10-09 19:26:09', '2020-10-09 19:26:09');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (123, '8445758-2.jpg', '/assets/uploads/2e2d529d1f8ced92215c2b7a34121876138b56548669e804ce0052505634d495.jpg', 32, '2020-10-09 19:26:09', '2020-10-09 19:26:09');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (124, '8445758-3.jpg', '/assets/uploads/0a6015076b25e3b4fb086f3f43cce9e7dcc488d69aaacd169f94d11376a945d6.jpg', 32, '2020-10-09 19:26:09', '2020-10-09 19:26:09');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (125, '8445758-4.jpg', '/assets/uploads/e55ad9d322a3786de07c91eb0d9b8c30c597dd3450e4ec6de7ab2f529795ebe0.jpg', 32, '2020-10-09 19:26:09', '2020-10-09 19:26:09');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (126, '8445758-4.jpg', '/assets/uploads/5d59d6736be572d0160073643fa5c881054e8599d92117027c7e9c1116cc346f.jpg', 33, '2020-10-09 19:26:28', '2020-10-09 19:26:28');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (127, '14628457-1-grey.jpg', '/assets/uploads/a79a83ab1dfa18170b2b4310cfe4bacb0807e2d64deaff702d7310558296e2be.jpg', 33, '2020-10-09 19:26:28', '2020-10-09 19:26:28');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (128, '14628457-2.jpg', '/assets/uploads/224ba7ab521c204b0bc499efca453ea117a998b76724afbfb59c9703b5da9c54.jpg', 33, '2020-10-09 19:26:28', '2020-10-09 19:26:28');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (129, '14628457-4.jpg', '/assets/uploads/24e773f4169278e2fd1658585521cc734a894f6f03662a3730397c127cead0a3.jpg', 33, '2020-10-09 19:26:28', '2020-10-09 19:26:28');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (130, '21175454-1-black.jpg', '/assets/uploads/d166a3ad89978d7006fc15308db451a07e6ad2b03bab18c79f9bb0d91cd9876c.jpg', 34, '2020-10-09 19:26:48', '2020-10-09 19:26:48');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (131, '21175454-2.jpg', '/assets/uploads/9bd70b4e3bd91a13dc8790c81f38515ba67e6490acb81cf5e5ee4657c078a9dd.jpg', 34, '2020-10-09 19:26:48', '2020-10-09 19:26:48');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (132, '21175454-3.jpg', '/assets/uploads/a81d3e33d9ce97f098fd99ce777712228e864ba1dbb3c1f4b993f8b232c13bb8.jpg', 34, '2020-10-09 19:26:48', '2020-10-09 19:26:48');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (133, '21175454-4.jpg', '/assets/uploads/177af0b61c0c5c1cac6bb85b5188c728013d08a75135b2da3b9fb5b36f708dc7.jpg', 34, '2020-10-09 19:26:48', '2020-10-09 19:26:48');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (134, '21652391-1-nocolour.jpg', '/assets/uploads/b08af4ab507f941d52fe0702d95f84e008e19add79c3abcd9b41b498cb6ac267.jpg', 35, '2020-10-09 19:27:21', '2020-10-09 19:27:21');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (135, '21652391-2.jpg', '/assets/uploads/d81bd5677f283664a4784b1afc30e7d9287d003fcfd7a3ae238e0a9a82788658.jpg', 35, '2020-10-09 19:27:21', '2020-10-09 19:27:21');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (136, '21652391-3.jpg', '/assets/uploads/fecd94f43e35b3b60490f4f0a3ae34a4eae718c1f84f5eccb6f21e4df3510d9b.jpg', 35, '2020-10-09 19:27:21', '2020-10-09 19:27:21');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (137, '21652391-4.jpg', '/assets/uploads/256cfe751f90439306b5c2d1efb1f5b0dc519eef12ff51dcdce9f765ac8179d3.jpg', 35, '2020-10-09 19:27:21', '2020-10-09 19:27:21');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (138, '9560010-1-purityfacemist.jpg', '/assets/uploads/28f9a194380f67871e1c8fb5be660c502a602134aa42c4c1adb28ee5a81f585d.jpg', 36, '2020-10-09 19:27:44', '2020-10-09 19:27:44');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (139, '9560010-2.jpg', '/assets/uploads/b7fffd47d5709c469b14bce8427e7d401fb2f14fb794eb1cb29dcd759e528d29.jpg', 36, '2020-10-09 19:27:44', '2020-10-09 19:27:44');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (140, '9560010-3.jpg', '/assets/uploads/db6420dd47505a2e75edee513d6ef9b4926eb8088b543e8c6a7909caf260be8e.jpg', 36, '2020-10-09 19:27:44', '2020-10-09 19:27:44');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (141, '9560010-4.jpg', '/assets/uploads/15d17aec55ed10de29951873400ddc6cf278ba87d457f4cb3a63b2ae5c66bd6f.jpg', 36, '2020-10-09 19:27:44', '2020-10-09 19:27:44');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (142, '10366800-1-playingkoi.jpg', '/assets/uploads/cd7a7e14b4c2b32700b559080788805801f8750fff7c87e01f6f0fc8bc1ca545.jpg', 37, '2020-10-09 19:28:09', '2020-10-09 19:28:09');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (144, '10366800-3.jpg', '/assets/uploads/f8706102a11031c591859bdc6998c5a8833975ad30cc5743c6ca4a469bc95651.jpg', 37, '2020-10-09 19:28:09', '2020-10-09 19:28:09');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (143, '10366800-2.jpg', '/assets/uploads/476f489ccbbe6c8a8c39c243e2740a9accd8f607ffc3d5915c212aaa8d832672.jpg', 37, '2020-10-09 19:28:09', '2020-10-09 19:28:09');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (145, '10366800-4.jpg', '/assets/uploads/dce40c8ae479832009d0f472d0363d5e59b65b545552f8a894cf0b7ab61761c0.jpg', 37, '2020-10-09 19:28:09', '2020-10-09 19:28:09');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (146, 'ayoral-navy-blue-milano-jersey-dress-349254-c00f3c1e9d5d40c6aac54f00b0995743d3676d9b.jpg', '/assets/uploads/2129efae99c9f054eaa86db8f6189a07cb606a4a0f09ba1ba6ee4aa183eacde6.jpg', 44, '2020-10-17 15:18:01', '2020-10-17 15:18:01');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (147, 'mayoral-navy-blue-milano-jersey-dress-349254-94f3a7166eff72deabaf66203a15b5fdf1a12378.jpg', '/assets/uploads/fc649f1f4953bf88a6d557b8dcf8cc5683b8c53ace89b6b01584cb55f00e0b1b.jpg', 44, '2020-10-17 15:18:01', '2020-10-17 15:18:01');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (148, 'mayoral-boys-brown-tweed-blazer-348438-1c5be97191c023b0dfce8342a6d9188c2423eeec.jpg', '/assets/uploads/6f59c117a7d07179311e1111565bc1d3621ca30ffb27177caab2e3ca6d684982.jpg', 45, '2020-10-17 15:20:38', '2020-10-17 15:20:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (149, 'mayoral-boys-brown-tweed-blazer-348438-facf9f97e884013446515aedc7a6af7c9973035f.jpg', '/assets/uploads/d6ae37882b925d8ad3f72c7c083af9582cb72431fd99f710021ba5cede188921.jpg', 45, '2020-10-17 15:20:38', '2020-10-17 15:20:38');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (150, 'dolce-gabbana-white-leather-logo-trainers-344560-0a814e373be7f7024ff0e3cde9a549c4965c8231.jpg', '/assets/uploads/2c420c1ec2ced723eb4007ae9e3b385d73c2225a7ec798060cbd116c0bfad743.jpg', 43, '2020-10-17 15:24:55', '2020-10-17 15:24:55');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (151, 'dolce-gabbana-white-leather-logo-trainers-344560-a1c6d96f39e1314d4c30c953297f011882b731e9.jpg', '/assets/uploads/09a21a91eaf96f67072f32bf6662a1c80755281c409755e0553e41092bbc52ca.jpg', 43, '2020-10-17 15:24:55', '2020-10-17 15:24:55');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (152, 'satila-of-sweden-fuchsia-pink-giant-pom-pom-hat-141894-2dc1600aa6b8adcfa8e0dfbc1183379e75a1b562.jpg', '/assets/uploads/9b52db04b8e612b5aa90bef24d69116454dd5304890d432778f6baf46b975d0f.jpg', 42, '2020-10-17 15:27:31', '2020-10-17 15:27:31');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (153, 'satila-of-sweden-fuchsia-pink-giant-pom-pom-hat-141894-73e7c1f51200405afe542a2ef0a9cf995d4a0817 (1).jpg', '/assets/uploads/909ace1a4c3da8fd224d120fa3408545554e8785535fd2b006fa76f2cc7ef02d.jpg', 42, '2020-10-17 15:27:31', '2020-10-17 15:27:31');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (154, 'satila-of-sweden-fuchsia-pink-giant-pom-pom-hat-141894-73e7c1f51200405afe542a2ef0a9cf995d4a0817.jpg', '/assets/uploads/909ace1a4c3da8fd224d120fa3408545554e8785535fd2b006fa76f2cc7ef02d.jpg', 42, '2020-10-17 15:27:31', '2020-10-17 15:27:31');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (155, 'satila-of-sweden-fuchsia-pink-giant-pom-pom-hat-141894-bd40b622b5be880a8989ca2cd5481942aba10c98.jpg', '/assets/uploads/a0109ea8a7c8a8076cbd5ceca63b5d369d7d0a5d326431e5e8646accd4dda3ed.jpg', 42, '2020-10-17 15:27:31', '2020-10-17 15:27:31');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (156, 'patachou-pink-suede-lace-trousers-325711-83bbfdf5766451fa727dd0b495634990696d6380.jpg', '/assets/uploads/a06632a553329645e6e9abe3aabf7dac8dcc144afd8cd91ff96d7e878a3120f5.jpg', 41, '2020-10-17 15:30:34', '2020-10-17 15:30:34');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (157, 'patachou-pink-suede-lace-trousers-325711-74937623afbc2831cc0cfe965f076122a028fdce.jpg', '/assets/uploads/5d34c57ba027d03e1a7c3f3c6209516f2884cb30010c72f4e4dbbdfd35516a19.jpg', 41, '2020-10-17 15:30:34', '2020-10-17 15:30:34');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (158, 'ayoral-beige-cotton-jeans-348497-906c1ee6bcb9143a14d9b25225ac8d9e880b88b0.jpg', '/assets/uploads/508bcaae7acab79e6a7fa58282f99ea0b54ad9b1eceeff7c064e2dfb77b46f68.jpg', 40, '2020-10-17 15:32:34', '2020-10-17 15:32:34');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (159, 'mayoral-beige-cotton-jeans-348497-34a4e961ff823cffef2bb0207e1335ccd5996c96.jpg', '/assets/uploads/5060781a956e6fbebb8be59e26f6fc23aa6938cd1e8299d2d2b7bf150210ef57.jpg', 40, '2020-10-17 15:32:34', '2020-10-17 15:32:34');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (160, 'beau-kid-boys-navy-blue-3-piece-suit-342922-2b069b6bd2990192eeae43835791a1d793f8cad1.jpg', '/assets/uploads/f7d502bd680e5f3350996d43f1818f33655aec5b86176240ebb34688bf7f84dd.jpg', 39, '2020-10-17 15:35:35', '2020-10-17 15:35:35');
INSERT INTO public.files (id, name, path, product_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (161, 'beau-kid-boys-navy-blue-3-piece-suit-342922-6568fb041c53f9fe98aa31f4d566abadf11a989e.jpg', '/assets/uploads/06c7cdf0a62e1c97dcf15bf239dc6643879a3ef05774d10dd481984d6cc6c691.jpg', 39, '2020-10-17 15:35:35', '2020-10-17 15:35:35');


--
-- TOC entry 2960 (class 0 OID 32965)
-- Dependencies: 227
-- Data for Name: post_categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.post_categories (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (5, 'Beauty', '2020-10-14 18:38:12', '2020-10-14 18:38:12');
INSERT INTO public.post_categories (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (4, 'Lifestyle', '2020-10-14 18:37:57', '2020-10-14 18:38:18');
INSERT INTO public.post_categories (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (2, 'Fashion', '2020-10-14 18:37:37', '2020-10-16 18:49:48');
INSERT INTO public.post_categories (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (3, 'Celebrity style', '2020-10-14 18:37:48', '2020-10-16 18:51:47');


--
-- TOC entry 2964 (class 0 OID 32979)
-- Dependencies: 231
-- Data for Name: post_tag_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.post_tag_items (id, post_id, tag_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (19, 1, 2, '2020-10-16 18:48:34', '2020-10-16 18:48:34');
INSERT INTO public.post_tag_items (id, post_id, tag_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (21, 4, 5, '2020-10-16 18:57:50', '2020-10-16 18:57:50');
INSERT INTO public.post_tag_items (id, post_id, tag_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (24, 7, 2, '2020-10-16 19:06:02', '2020-10-16 19:06:02');
INSERT INTO public.post_tag_items (id, post_id, tag_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (25, 7, 3, '2020-10-16 19:06:02', '2020-10-16 19:06:02');
INSERT INTO public.post_tag_items (id, post_id, tag_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (26, 3, 5, '2020-10-16 19:06:32', '2020-10-16 19:06:32');
INSERT INTO public.post_tag_items (id, post_id, tag_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (27, 5, 4, '2020-10-16 19:06:52', '2020-10-16 19:06:52');
INSERT INTO public.post_tag_items (id, post_id, tag_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (28, 6, 4, '2020-10-16 19:07:34', '2020-10-16 19:07:34');


--
-- TOC entry 2962 (class 0 OID 32972)
-- Dependencies: 229
-- Data for Name: post_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.post_tags (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (2, 'Fashion', '2020-10-14 14:16:14', '2020-10-14 14:16:14');
INSERT INTO public.post_tags (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (3, 'Street style', '2020-10-14 18:39:47', '2020-10-14 18:39:47');
INSERT INTO public.post_tags (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (4, 'Diversity', '2020-10-14 18:39:57', '2020-10-14 18:39:57');
INSERT INTO public.post_tags (id, name, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (5, 'Beauty', '2020-10-14 18:40:07', '2020-10-14 18:40:07');


--
-- TOC entry 2958 (class 0 OID 32955)
-- Dependencies: 225
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.posts (id, title, body, image_path, category_id, created_at, updated_at, image_name, author) OVERRIDING SYSTEM VALUE VALUES (1, 'Kelly Rowland Wants You To Feel Fabulous While Working from Home ', 'The last few months have been eventful for Kelly Rowland. As a singer, television personality, and one-third of the iconic girl group Destiny’s Child, Rowland is used to tackling multiple projects. Still, this week she’s delighted fans with three big announcements. On the heels of a newly released single, “Crazy,” which hit Spotify this morning, and the reveal of her second pregnancy with husband Tim Weatherspoon, Rowland dropped the news that she’s launching a capsule collection with JustFab. Though she stepped into the athleisure space last year with a fitness-focused range for Fabletics, Rowland’s latest designs tread new territory. Conceived as a modern wardrobe that offered style and affordability, Rowland set out to make women feel their best. “I think that everybody wants to be fabulous, but they want to be able to do that in a way that doesn’t break the bank,” she shared via Zoom. “Especially right now in these times that we’re living in.”

The pieces reflect Rowland’s essentials and a pragmatic take on getting ready. Though she’s apt to wear Iris van Herpen couture on the red carpet or Alexander McQueen while judging The Voice Australia, Rowland appreciates the ease of staples like oversized sweaters, skinny jeans, and flat riding boots when she’s at home. “I’m not as big on trends as I was when I was younger,” she explains. “I love classic pieces and understanding the value of having things in your closet that work for you long term.” Usually, investment pieces are synonymous with high costs, but during her conversations with JustFab CMO Daria Burke, Rowland realized she wanted to bring that concept to an accessible price point. “When we started getting into the studio and collaborating, we leaned on classic styles that made sense for women in quarantine,” says Rowland. “[Our woman] still feels stylish, but she’s comfortable; she isn’t going to hurt her ankles trying to get into a boot that is too high. If she’s doing the most, she’s doing it in her space and her way.” With a lineup that includes knits, leather pants, fitted blazers, boucle coats, and croc-stamped heels the collection provides plenty of opportunities to dress things up.', '/assets/uploads/c3716b2da02decefb10b7d56610f029288e7642444b3e517b7abc9c513d831d2.jpg', 2, '2020-10-14 18:46:59', '2020-10-16 18:48:34', 'KW-header.jpg', 'JANELLE OKWODU');
INSERT INTO public.posts (id, title, body, image_path, category_id, created_at, updated_at, image_name, author) OVERRIDING SYSTEM VALUE VALUES (5, 'The Best Street Style at Tokyo Fashion Week Spring 2021', 'The spring 2021 season continues apace in Asia. Shanghai Fashion Week is wrapping up this weekend, and Tokyo Fashion Week is just getting started. Kira is on the ground looking for the city’s best street style, and so far it’s a bold, diverse mix. Some locals are going for easy oversized suits, while others are going all out with clashing prints, girly dresses, and experimental layers. One common theme? Practical, walkable shoes, from cushy sneakers to platform boots. Another popular accessory: a mask, naturally. Scroll through our latest coverage below, and come back for Kira’s daily updates.', '/assets/uploads/6973caf30698b4dc7147181b48f81f42d7216d658c9413d5c53bf1e2a295a488.jpg', 2, '2020-10-15 12:35:12', '2020-10-16 19:06:52', '00-story.jpg', 'Ema Timahe');
INSERT INTO public.posts (id, title, body, image_path, category_id, created_at, updated_at, image_name, author) OVERRIDING SYSTEM VALUE VALUES (4, 'Irina Shayk Wears the Most Satisfying Skirt of Fall ', 'Does any garment embody the deliciously crisp fall season like a thigh-skimming, plaid, pleated skirt? I think not. Today, Irina Shayk stepped out in New York City wearing a kicky number by Burberry, in the label’s signature plaid. Shayk, both a supermodel and a glistening muse to Burberry designer Ricardo Tisci, was wearing other items from the brand, including a matching plaid purse and a black hoodie with a “B” emblazoned on the chest. For that final crisp, autumnal moment, she opted for a pair of black tights and black oxfords. 

There is something so satisfying about the skirt with its super sharp pleats. It fits in seamlessly with leaves crunching underneath the feet, or a steaming cup of apple cider that ever-so gently touches the lips. Maybe that is because it reminds me of a uniform-wearing student buttoned up in some cardigan and carrying a stack of books. (Shayk, for her part, held an orange drink and not books.) But Shayk manages to make the parochial look more casual and loosened up with a skirt-and-hoodie combo. Bless her.  ', '/assets/uploads/28a61f8e562ffe97effc0f118c6ce77b46d07bfa998e5cc544ef09ff2f52f72a.jpg', 3, '2020-10-15 12:31:21', '2020-10-16 18:57:50', 'SPL5192511_002.jpg', 'LIANA SATENSTEIN');
INSERT INTO public.posts (id, title, body, image_path, category_id, created_at, updated_at, image_name, author) OVERRIDING SYSTEM VALUE VALUES (7, 'Levi’s Launches Its Own Recommerce Site and Buyback Program, Levi’s Secondhand', 'Finding the perfect pair of vintage Levi’s used to require hours at a thrift store, endless eBay searches, and often a few visits to the tailor. For some of us, it was among the most noble of fashion pursuits; for others, it was just too much work. Today, Levi’s is making it a little easier with the launch of Levi’s Secondhand, a recommerce site for previously worn Levi’s jeans and denim jackets. Some of it will be handpicked vintage, but most of the garments will be sourced directly from Levi’s customers: Starting now, anyone can turn in any Levi’s denim item—even if it’s damaged—for a gift card towards a future purchase.

It marks a significant turning point, both for Levi’s and the fashion industry as a whole. Levi’s is the first denim brand of its size to create a buyback program like this and effectively take responsibility for the full “life cycle” of its garments. It’s an example of true circularity: You could buy a brand-new pair of Levi’s tomorrow, and you’d know exactly what the “end use” might be, should you tire of them in a few years. For conscious shoppers, that’s often the difference between buying something or… not. How long will I wear this? Is it built to last? What will happen to it when I don’t want it anymore?', '/assets/uploads/da61008f92e0d7cab8d61fe52eea0f5790aedb7c040015a8cbadd6072cdf40d1.jpg', 5, '2020-10-15 13:36:52', '2020-10-16 19:06:02', 'Levis_17.jpg', 'EMILY FARRA');
INSERT INTO public.posts (id, title, body, image_path, category_id, created_at, updated_at, image_name, author) OVERRIDING SYSTEM VALUE VALUES (3, 'Meghan Markle Is Spotlighting Ethical Fashion In Her Second Act', 'Meghan Markle is in the middle of a reinvention. After stepping away from royal duties in January, the Duke and Duchess of Sussex have been busy in Los Angeles, building their lives as private citizens. The couple made headlines in September with the announcement of their production company and multiyear deal with Netflix, and as philanthropists, the pair have been taking on speaking engagements, interviews, and public appearances connected to causes they hold dear. Much of their recent work has been via Zoom or teleconferencing, but the Duchess still sends messages through her outfit choices, even over webcam. Markle’s wardrobe has always reflected her socially-conscious viewpoint, and lately, she’s doubled down on ethical dressing by exclusively wearing brands that focus on the common good in her interviews.', '/assets/uploads/c1a77e11a11f76e8980929c6934c504b5f7c4ef44d9d5ebac224338174eaa951.jpg', 3, '2020-10-14 20:47:16', '2020-10-16 19:06:32', 'mmm-post-horizontal.jpg', 'JANELLE OKWODU');
INSERT INTO public.posts (id, title, body, image_path, category_id, created_at, updated_at, image_name, author) OVERRIDING SYSTEM VALUE VALUES (6, 'Kate Moss’s Daughter Lila Grace Makes Her Fashion Week Debut at Miu Miu', 'As far as fashion months go, it’s been a major few weeks for runway star-in-the-making Lila Grace Moss. Just days after celebrating her 18th birthday, the daughter of Kate Moss and Dazed magazine cofounder Jefferson Hack made a blockbuster debut on Miu Miu’s SS21 catwalk. The show, which was broadcast from Milan on October 6, is the exclamation point to end an extraordinary season that’s seen the virtual FROW become something of a new normal.

From the lifestyle changes she’s making for the health of the planet to the TV show that got her and and her supermodel mom through lockdown—this is your one-minute meet up with the fashion world’s brightest new star.

Image may contain Clothing Apparel Skirt Human Person and Miniskirt
Lila Moss, Miu Miu SS21Courtesy of Miu Miu
Vogue Logo
Get 1 Year of Vogue
plus an exclusive gift

SUBSCRIBE NOW
Woman
Vogue: Hi, Lila, congratulations on your first runway show! First up, the details: What should we be zooming in on?

Lila Grace Moss: The makeup was very natural, but with extra attention to detail—look out for the eyebrow slit.

You’ve already worked with luminaries including David Bailey, Tim Walker, and David Sims. Which fashion creatives are you excited to collaborate with next?', '/assets/uploads/d3e0854d5e140f4fea61cb01e4bbfdda4caa2fb413f809731b932af9c7e2ed27.jpg', 5, '2020-10-15 12:39:01', '2020-10-16 19:07:34', 'LOOK-53_DSC5714.jpg', 'JULIA HOBBS');


--
-- TOC entry 2950 (class 0 OID 24784)
-- Dependencies: 217
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (4, 'Leather jacket in black', 29900, 0, 1, 0, 0, 0, 0, 1, 2, 5, 34, '2020-09-22 19:20:27', '2020-10-09 19:05:38', 'Description');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (5, 'Tall bonded plush overcoat in brown', 7500, 0, 1, 0, 1, 0, 0, 2, 4, 5, 35, '2020-09-24 22:21:10', '2020-10-09 19:09:08', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (6, 'Jacket in black', 3000, 0, 0, 0, 0, 0, 0, 1, 2, 21, 0, '2020-09-25 13:26:05', '2020-10-09 19:11:13', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (7, 'Puffer jacket', 6000, 0, 0, 0, 1, 0, 0, 2, 22, 5, 34, '2020-09-25 14:40:05', '2020-10-09 19:11:26', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (8, 'Embellished drape maxi dress', 14500, 0, 1, 0, 1, 0, 0, 4, 11, 6, 34, '2020-09-25 20:09:56', '2020-10-09 19:12:01', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (9, 'Oversized smock back sweat dress', 2000, 0, 1, 0, 0, 0, 0, 2, 8, 6, 35, '2020-09-25 21:08:04', '2020-10-09 19:12:50', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (22, 'Whistles Cardi suede heel', 14000, 0, 1, 0, 0, 0, 0, 1, 2, 17, 13, '2020-10-05 11:16:44', '2020-10-09 19:20:06', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (10, 'Zip fleece in pink', 5000, 0, 1, 0, 1, 0, 0, 5, 13, 7, 33, '2020-10-02 19:44:05', '2020-10-09 19:13:39', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (11, 'Mix & match oversized sweat', 2000, 0, 1, 0, 0, 0, 0, 2, 15, 7, 34, '2020-10-02 19:46:03', '2020-10-09 19:13:58', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (12, 'Embroidered flag logo hoodie', 9000, 0, 0, 0, 1, 0, 0, 7, 2, 7, 35, '2020-10-02 19:51:37', '2020-10-09 19:14:40', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (13, 'Calvin Klein Pieced lounge sweatshirt', 6000, 0, 0, 0, 1, 0, 0, 6, 22, 7, 34, '2020-10-02 19:55:04', '2020-10-09 19:14:59', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (15, 'High rise ''slouchy'' mom jeans', 3200, 0, 1, 0, 0, 0, 0, 2, 3, 9, 35, '2020-10-02 20:05:36', '2020-10-09 19:16:34', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (14, 'High rise mom jeans', 5500, 0, 1, 0, 1, 0, 0, 8, 3, 9, 34, '2020-10-02 20:03:55', '2020-10-09 19:17:06', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (16, 'Calvin Klein dad fit jeans', 8500, 0, 0, 0, 0, 0, 0, 6, 3, 9, 2, '2020-10-02 20:10:48', '2020-10-09 19:17:31', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (17, 'Topman knitted jumper', 2999, 0, 0, 0, 1, 0, 0, 10, 9, 10, 35, '2020-10-02 20:21:15', '2020-10-09 19:18:06', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (18, 'JDY roll neck jumper ', 2000, 0, 1, 0, 0, 0, 0, 9, 13, 10, 34, '2020-10-02 20:22:33', '2020-10-09 19:18:30', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (19, 'Polo shirt', 1125, 2500, 0, 0, 0, 1, 55, 11, 9, 13, 34, '2020-10-05 10:44:48', '2020-10-09 19:18:58', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (20, 'Pleated satin midi skirt', 5500, 0, 1, 0, 0, 0, 0, 12, 4, 12, 34, '2020-10-05 10:47:47', '2020-10-09 19:19:19', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (21, 'Heel ankle boots', 4200, 0, 1, 0, 0, 0, 0, 13, 19, 16, 14, '2020-10-05 11:10:34', '2020-10-09 19:19:40', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (23, 'Embellished heeled sandals', 3299, 0, 1, 0, 0, 0, 0, 14, 1, 18, 14, '2020-10-05 11:21:26', '2020-10-09 19:20:38', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (35, 'Larkstaden vegan hand wash', 800, 0, 1, 0, 0, 0, 0, 20, 1, 31, 0, '2020-10-05 12:32:50', '2020-10-09 19:27:21', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (24, 'Boots in brown suede', 8000, 0, 0, 0, 0, 0, 0, 15, 4, 16, 18, '2020-10-05 11:29:26', '2020-10-09 19:22:04', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (25, 'Sandals in black leather', 3500, 0, 0, 0, 0, 0, 0, 2, 2, 18, 17, '2020-10-05 11:30:15', '2020-10-09 19:22:43', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (26, 'UGG Cozette slide slippers', 7000, 0, 1, 0, 0, 0, 0, 16, 23, 19, 14, '2020-10-05 11:35:14', '2020-10-09 19:23:10', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (27, 'Nike Air Max 2090 trainers', 12995, 0, 1, 0, 0, 0, 0, 17, 10, 20, 13, '2020-10-05 11:42:48', '2020-10-09 19:23:41', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (28, 'Nike SB Charge suede trainers', 5295, 0, 0, 0, 1, 0, 0, 17, 9, 20, 18, '2020-10-05 11:44:00', '2020-10-09 19:24:17', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (29, 'Underarm logo detail bag', 3000, 0, 1, 0, 0, 0, 0, 18, 2, 21, 0, '2020-10-05 11:49:50', '2020-10-09 19:24:49', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (30, 'River Island monogram logo belt', 1400, 0, 1, 0, 0, 0, 0, 18, 4, 22, 0, '2020-10-05 12:10:22', '2020-10-09 19:25:20', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (31, 'River Island monogram suedette belt', 1600, 0, 0, 0, 0, 0, 0, 18, 2, 22, 0, '2020-10-05 12:13:02', '2020-10-09 19:25:38', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (32, 'ASOS DESIGN felt fedora hat', 2000, 0, 1, 0, 0, 0, 0, 2, 2, 23, 0, '2020-10-05 12:16:51', '2020-10-09 19:26:09', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (33, 'ASOS DESIGN trucker cap', 1000, 0, 0, 0, 0, 0, 0, 2, 9, 23, 0, '2020-10-05 12:21:19', '2020-10-09 19:26:28', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (34, 'Leather billfold wallet', 8900, 0, 0, 0, 0, 0, 0, 19, 2, 30, 0, '2020-10-05 12:28:36', '2020-10-09 19:26:48', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (36, 'Self Tan Purity Face Mist', 2200, 0, 0, 0, 0, 0, 0, 21, 1, 31, 0, '2020-10-05 12:38:17', '2020-10-09 19:27:44', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (37, 'Essie Nail Polish', 799, 0, 1, 0, 0, 0, 0, 22, 4, 34, 0, '2020-10-05 12:43:59', '2020-10-09 19:28:09', 'Text');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (44, 'Navy Blue Milano Jersey Dress', 1000, 0, 1, 1, 1, 0, 0, 13, 11, 6, 32, '2020-10-09 12:47:24', '2020-10-17 15:18:01', 'description');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (45, 'Boys Brown Tweed Blazer', 1000, 0, 0, 1, 1, 0, 0, 13, 4, 5, 32, '2020-10-09 12:48:12', '2020-10-17 15:20:38', 'Description');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (43, 'White Leather Logo Trainers', 1200, 0, 1, 1, 0, 0, 0, 2, 1, 20, 4, '2020-10-09 12:45:08', '2020-10-17 15:24:55', 'Description');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (42, 'Fuchsia Pink Giant Pom-Pom Hat', 1200, 0, 1, 1, 0, 0, 0, 20, 13, 23, 0, '2020-10-09 09:59:20', '2020-10-17 15:27:31', 'Description');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (40, 'Beige Cotton Jeans', 1200, 0, 0, 1, 0, 0, 0, 20, 1, 9, 32, '2020-10-08 23:18:42', '2020-10-17 15:32:34', 'Description');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (39, 'Boys Navy Blue 3 Piece Suit', 2500, 0, 0, 1, 0, 0, 0, 20, 3, 10, 32, '2020-10-08 22:54:19', '2020-10-17 15:35:35', 'Description');
INSERT INTO public.products (id, title, price, old_price, gender, is_kids, is_new, is_discount, dsc_percent, brand_id, color_id, category_id, size_id, created_at, updated_at, description) OVERRIDING SYSTEM VALUE VALUES (41, 'Pink Suede Lace Trousers', 1200, 0, 1, 1, 0, 0, 0, 20, 13, 9, 32, '2020-10-08 23:31:20', '2020-10-17 16:00:56', 'Test');


--
-- TOC entry 2936 (class 0 OID 16498)
-- Dependencies: 203
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.sessions (id, session_id, user_id, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (118, '20866737-f811-4101-b0f7-4e82afda20b1', 1, '2020-10-20 21:42:05', '2020-10-20 21:42:05');


--
-- TOC entry 2948 (class 0 OID 24766)
-- Dependencies: 215
-- Data for Name: shoe_sizes; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.shoe_sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (1, 'XXS', '19:44:11', '2020-09-18 19:44:21', NULL);
INSERT INTO public.shoe_sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (2, 'XS', '19:44:38', '2020-09-18 19:44:38', NULL);
INSERT INTO public.shoe_sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (3, 'S', '19:44:54', '2020-09-18 19:44:54', NULL);
INSERT INTO public.shoe_sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (4, 'M', '19:45:03', '2020-09-18 19:45:03', NULL);
INSERT INTO public.shoe_sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (5, 'L', '19:45:10', '2020-09-18 19:45:10', NULL);
INSERT INTO public.shoe_sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (6, 'XL', '19:45:18', '2020-09-18 19:45:18', NULL);
INSERT INTO public.shoe_sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (7, '2XL', '19:45:28', '2020-09-18 19:45:28', NULL);
INSERT INTO public.shoe_sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (8, '3XL', '19:45:40', '2020-09-18 19:45:40', NULL);
INSERT INTO public.shoe_sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (9, '4XL', '19:45:53', '2020-09-18 19:45:53', NULL);
INSERT INTO public.shoe_sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (10, '5XL', '19:46:04', '2020-09-18 19:46:04', NULL);


--
-- TOC entry 2946 (class 0 OID 24759)
-- Dependencies: 213
-- Data for Name: sizes; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (31, '50', '2020-09-18 19:09:13', '2020-09-18 19:09:13', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (33, 'XS', '2020-09-19 14:37:29', '2020-09-19 14:37:29', 0);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (34, 'S', '2020-09-19 14:37:44', '2020-09-19 14:37:44', 0);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (35, 'M', '2020-09-19 14:37:48', '2020-09-19 14:37:48', 0);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (36, 'L', '2020-09-19 14:37:58', '2020-09-19 14:37:58', 0);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (37, 'XL', '2020-09-19 14:38:41', '2020-09-19 14:38:41', 0);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (38, '2XL', '2020-09-19 14:38:54', '2020-09-19 14:38:54', 0);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (39, '3XL', '2020-09-19 14:39:01', '2020-09-19 14:39:01', 0);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (40, '4XL', '2020-09-19 14:39:09', '2020-09-19 14:39:09', 0);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (41, '5XL', '2020-09-19 14:39:15', '2020-09-19 14:39:15', 0);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (32, 'XXS', '2020-09-19 13:58:44', '2020-09-19 13:58:44', 0);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (2, '32', '2020-09-18 19:01:17', '2020-09-19 13:58:21', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (3, '32.5', '2020-09-18 19:01:33', '2020-09-18 19:01:33', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (4, '33', '2020-09-18 19:01:40', '2020-09-18 19:01:40', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (5, '33.5', '2020-09-18 19:01:58', '2020-09-18 19:01:58', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (6, '34', '2020-09-18 19:02:06', '2020-09-18 19:02:06', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (7, '34.5', '2020-09-18 19:02:18', '2020-09-18 19:02:18', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (8, '35', '2020-09-18 19:02:42', '2020-09-18 19:02:42', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (9, '35.5', '2020-09-18 19:02:58', '2020-09-18 19:02:58', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (10, '36.5', '2020-09-18 19:03:07', '2020-09-18 19:03:07', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (11, '37', '2020-09-18 19:03:24', '2020-09-18 19:03:24', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (12, '37.5', '2020-09-18 19:03:31', '2020-09-18 19:03:31', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (13, '38', '2020-09-18 19:03:36', '2020-09-18 19:03:36', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (14, '39', '2020-09-18 19:04:03', '2020-09-18 19:04:03', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (15, '39.5', '2020-09-18 19:04:17', '2020-09-18 19:04:17', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (16, '40', '2020-09-18 19:04:46', '2020-09-18 19:04:46', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (17, '40.5', '2020-09-18 19:04:50', '2020-09-18 19:05:06', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (18, '41.5', '2020-09-18 19:06:19', '2020-09-18 19:06:19', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (19, '42.5', '2020-09-18 19:06:29', '2020-09-18 19:06:29', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (20, '43.5', '2020-09-18 19:06:46', '2020-09-18 19:06:46', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (21, '44', '2020-09-18 19:06:50', '2020-09-18 19:06:50', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (22, '44.5', '2020-09-18 19:07:03', '2020-09-18 19:07:03', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (23, '45.5', '2020-09-18 19:07:08', '2020-09-18 19:07:08', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (24, '46', '2020-09-18 19:07:34', '2020-09-18 19:07:34', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (25, '46.5', '2020-09-18 19:07:43', '2020-09-18 19:07:43', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (26, '47', '2020-09-18 19:07:55', '2020-09-18 19:08:04', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (27, '47.5', '2020-09-18 19:08:10', '2020-09-18 19:08:10', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (28, '48.5', '2020-09-18 19:08:59', '2020-09-18 19:08:59', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (29, '49', '2020-09-18 19:09:04', '2020-09-18 19:09:04', 1);
INSERT INTO public.sizes (id, size, created_at, updated_at, type) OVERRIDING SYSTEM VALUE VALUES (30, '49.5', '2020-09-18 19:09:09', '2020-09-18 19:09:09', 1);


--
-- TOC entry 2935 (class 0 OID 16394)
-- Dependencies: 202
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, first_name, last_name, email, password, created_at, updated_at, role) OVERRIDING SYSTEM VALUE VALUES (1, 'Admin', 'Admin', 'admin@gmail.com', '$2a$04$bcls3zmmrHEp5rjYdwdxfOU9O11y5LXctEn1A4IpAPfSqQwQnDbRK', '2020-09-08 13:21:41', '2020-09-08 13:21:41', 1);
INSERT INTO public.users (id, first_name, last_name, email, password, created_at, updated_at, role) OVERRIDING SYSTEM VALUE VALUES (3, 'John', 'Doe', 'johndoe@gmail.com', '$2a$04$bcls3zmmrHEp5rjYdwdxfOU9O11y5LXctEn1A4IpAPfSqQwQnDbRK', '2020-09-09 16:55:32', '2020-09-09 16:55:32', 0);


--
-- TOC entry 2974 (class 0 OID 0)
-- Dependencies: 208
-- Name: brands_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.brands_id_seq', 22, true);


--
-- TOC entry 2975 (class 0 OID 0)
-- Dependencies: 221
-- Name: cart_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cart_id_seq', 8, true);


--
-- TOC entry 2976 (class 0 OID 0)
-- Dependencies: 222
-- Name: carts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.carts_id_seq', 2, true);


--
-- TOC entry 2977 (class 0 OID 0)
-- Dependencies: 207
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 37, true);


--
-- TOC entry 2978 (class 0 OID 0)
-- Dependencies: 214
-- Name: clothes_sizes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.clothes_sizes_id_seq', 11, true);


--
-- TOC entry 2979 (class 0 OID 0)
-- Dependencies: 210
-- Name: colors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.colors_id_seq', 23, true);


--
-- TOC entry 2980 (class 0 OID 0)
-- Dependencies: 218
-- Name: files_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.files_id_seq', 161, true);


--
-- TOC entry 2981 (class 0 OID 0)
-- Dependencies: 224
-- Name: post_categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.post_categories_id_seq', 7, true);


--
-- TOC entry 2982 (class 0 OID 0)
-- Dependencies: 226
-- Name: post_categories_id_seq1; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.post_categories_id_seq1', 5, true);


--
-- TOC entry 2983 (class 0 OID 0)
-- Dependencies: 230
-- Name: post_tag_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.post_tag_items_id_seq', 28, true);


--
-- TOC entry 2984 (class 0 OID 0)
-- Dependencies: 228
-- Name: post_tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.post_tags_id_seq', 5, true);


--
-- TOC entry 2985 (class 0 OID 0)
-- Dependencies: 216
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_id_seq', 45, true);


--
-- TOC entry 2986 (class 0 OID 0)
-- Dependencies: 204
-- Name: sessions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sessions_id_seq', 118, true);


--
-- TOC entry 2987 (class 0 OID 0)
-- Dependencies: 212
-- Name: shoe_sizes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.shoe_sizes_id_seq', 42, true);


--
-- TOC entry 2988 (class 0 OID 0)
-- Dependencies: 205
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 9, true);


--
-- TOC entry 2786 (class 2606 OID 24791)
-- Name: brands brands_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brands
    ADD CONSTRAINT brands_pkey PRIMARY KEY (id);


--
-- TOC entry 2798 (class 2606 OID 32943)
-- Name: cart_items cart_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_pkey PRIMARY KEY (id);


--
-- TOC entry 2800 (class 2606 OID 32952)
-- Name: carts carts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_pkey PRIMARY KEY (id);


--
-- TOC entry 2784 (class 2606 OID 24798)
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- TOC entry 2792 (class 2606 OID 24835)
-- Name: shoe_sizes clothes_sizes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shoe_sizes
    ADD CONSTRAINT clothes_sizes_pkey PRIMARY KEY (id);


--
-- TOC entry 2788 (class 2606 OID 24810)
-- Name: colors colors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.colors
    ADD CONSTRAINT colors_pkey PRIMARY KEY (id);


--
-- TOC entry 2796 (class 2606 OID 32936)
-- Name: files files_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_pkey PRIMARY KEY (id);


--
-- TOC entry 2802 (class 2606 OID 32962)
-- Name: posts post_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT post_categories_pkey PRIMARY KEY (id);


--
-- TOC entry 2804 (class 2606 OID 32969)
-- Name: post_categories post_categories_pkey1; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_categories
    ADD CONSTRAINT post_categories_pkey1 PRIMARY KEY (id);


--
-- TOC entry 2808 (class 2606 OID 32983)
-- Name: post_tag_items post_tag_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_tag_items
    ADD CONSTRAINT post_tag_items_pkey PRIMARY KEY (id);


--
-- TOC entry 2806 (class 2606 OID 32976)
-- Name: post_tags post_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_tags
    ADD CONSTRAINT post_tags_pkey PRIMARY KEY (id);


--
-- TOC entry 2794 (class 2606 OID 24821)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 2782 (class 2606 OID 16508)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 2790 (class 2606 OID 24828)
-- Name: sizes shoe_sizes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sizes
    ADD CONSTRAINT shoe_sizes_pkey PRIMARY KEY (id);


--
-- TOC entry 2779 (class 2606 OID 16398)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2780 (class 1259 OID 16506)
-- Name: fki_sessions_user_id_fkey; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_sessions_user_id_fkey ON public.sessions USING btree (user_id);


-- Completed on 2020-10-20 22:04:21

--
-- PostgreSQL database dump complete
--

