--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4 (Debian 17.4-1.pgdg120+2)
-- Dumped by pg_dump version 17.4

-- Started on 2025-06-23 17:08:09 UTC

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

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--
--
-- TOC entry 3516 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 16385)
-- Name: accident_character; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.accident_character (
    id integer NOT NULL,
    name character varying NOT NULL,
    id_specialization integer NOT NULL
);


ALTER TABLE public.accident_character OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16390)
-- Name: accident_character_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.accident_character_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.accident_character_id_seq OWNER TO postgres;

--
-- TOC entry 3517 (class 0 OID 0)
-- Dependencies: 218
-- Name: accident_character_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.accident_character_id_seq OWNED BY public.accident_character.id;


--
-- TOC entry 219 (class 1259 OID 16391)
-- Name: accident_content; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.accident_content (
    id integer NOT NULL,
    name character varying NOT NULL,
    id_character integer NOT NULL,
    recommended_time integer
);


ALTER TABLE public.accident_content OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16396)
-- Name: accident_content_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.accident_content_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.accident_content_id_seq OWNER TO postgres;

--
-- TOC entry 3518 (class 0 OID 0)
-- Dependencies: 220
-- Name: accident_content_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.accident_content_id_seq OWNED BY public.accident_content.id;


--
-- TOC entry 221 (class 1259 OID 16397)
-- Name: appeal; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.appeal (
    id integer NOT NULL,
    description text,
    create_date timestamp without time zone,
    id_accident integer NOT NULL,
    applicant_name character varying,
    applicant_number character varying,
    id_application integer,
    address text,
    additional_number character varying,
    is_active boolean DEFAULT true
);


ALTER TABLE public.appeal OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 16403)
-- Name: appeal_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.appeal_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.appeal_id_seq OWNER TO postgres;

--
-- TOC entry 3519 (class 0 OID 0)
-- Dependencies: 222
-- Name: appeal_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.appeal_id_seq OWNED BY public.appeal.id;


--
-- TOC entry 223 (class 1259 OID 16404)
-- Name: application; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.application (
    id bigint NOT NULL,
    description text,
    create_date timestamp without time zone,
    id_accident integer NOT NULL,
    id_importance integer NOT NULL,
    id_status integer NOT NULL,
    address text,
    accident_cause text,
    damage_point text,
    id_material integer,
    id_damage integer,
    id_brigade integer,
    id_operator integer NOT NULL,
    start_date timestamp without time zone,
    finish_date timestamp without time zone,
    CONSTRAINT application_check CHECK ((((id_status = ANY (ARRAY[2, 3, 4])) AND (id_brigade IS NOT NULL)) OR (id_status = 1)))
);


ALTER TABLE public.application OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 16410)
-- Name: application_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.application_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.application_id_seq OWNER TO postgres;

--
-- TOC entry 3520 (class 0 OID 0)
-- Dependencies: 224
-- Name: application_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.application_id_seq OWNED BY public.application.id;


--
-- TOC entry 225 (class 1259 OID 16411)
-- Name: brigade; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.brigade (
    id integer NOT NULL,
    id_brigadir integer,
    people_count integer
);


ALTER TABLE public.brigade OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 16414)
-- Name: brigade_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.brigade_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.brigade_id_seq OWNER TO postgres;

--
-- TOC entry 3521 (class 0 OID 0)
-- Dependencies: 226
-- Name: brigade_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.brigade_id_seq OWNED BY public.brigade.id;


--
-- TOC entry 227 (class 1259 OID 16415)
-- Name: damage_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.damage_type (
    id integer NOT NULL,
    name character varying
);


ALTER TABLE public.damage_type OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 16420)
-- Name: damage_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.damage_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.damage_type_id_seq OWNER TO postgres;

--
-- TOC entry 3522 (class 0 OID 0)
-- Dependencies: 228
-- Name: damage_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.damage_type_id_seq OWNED BY public.damage_type.id;


--
-- TOC entry 229 (class 1259 OID 16421)
-- Name: importance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.importance (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.importance OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 16426)
-- Name: importance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.importance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.importance_id_seq OWNER TO postgres;

--
-- TOC entry 3523 (class 0 OID 0)
-- Dependencies: 230
-- Name: importance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.importance_id_seq OWNED BY public.importance.id;


--
-- TOC entry 231 (class 1259 OID 16427)
-- Name: material_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.material_type (
    id integer NOT NULL,
    name character varying
);


ALTER TABLE public.material_type OWNER TO postgres;

--
-- TOC entry 232 (class 1259 OID 16432)
-- Name: material_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.material_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.material_type_id_seq OWNER TO postgres;

--
-- TOC entry 3524 (class 0 OID 0)
-- Dependencies: 232
-- Name: material_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.material_type_id_seq OWNED BY public.material_type.id;


--
-- TOC entry 233 (class 1259 OID 16433)
-- Name: role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.role OWNER TO postgres;

--
-- TOC entry 234 (class 1259 OID 16438)
-- Name: role_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.role_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.role_id_seq OWNER TO postgres;

--
-- TOC entry 3525 (class 0 OID 0)
-- Dependencies: 234
-- Name: role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.role_id_seq OWNED BY public.role.id;


--
-- TOC entry 235 (class 1259 OID 16439)
-- Name: shutdown; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shutdown (
    id integer NOT NULL,
    address text,
    id_accident integer,
    date timestamp without time zone,
    is_active boolean DEFAULT true,
    day_count integer,
    id_application integer,
    description text,
    id_type integer,
    hour_count integer
);


ALTER TABLE public.shutdown OWNER TO postgres;

--
-- TOC entry 236 (class 1259 OID 16445)
-- Name: shutdown_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.shutdown_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.shutdown_id_seq OWNER TO postgres;

--
-- TOC entry 3526 (class 0 OID 0)
-- Dependencies: 236
-- Name: shutdown_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.shutdown_id_seq OWNED BY public.shutdown.id;


--
-- TOC entry 237 (class 1259 OID 16446)
-- Name: shutdown_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shutdown_type (
    id integer NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.shutdown_type OWNER TO postgres;

--
-- TOC entry 238 (class 1259 OID 16449)
-- Name: specialization; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.specialization (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.specialization OWNER TO postgres;

--
-- TOC entry 239 (class 1259 OID 16454)
-- Name: specialization_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.specialization_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.specialization_id_seq OWNER TO postgres;

--
-- TOC entry 3527 (class 0 OID 0)
-- Dependencies: 239
-- Name: specialization_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.specialization_id_seq OWNED BY public.specialization.id;


--
-- TOC entry 240 (class 1259 OID 16455)
-- Name: status; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.status (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.status OWNER TO postgres;

--
-- TOC entry 241 (class 1259 OID 16460)
-- Name: status_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.status_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.status_id_seq OWNER TO postgres;

--
-- TOC entry 3528 (class 0 OID 0)
-- Dependencies: 241
-- Name: status_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.status_id_seq OWNED BY public.status.id;


--
-- TOC entry 242 (class 1259 OID 16461)
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    login character varying NOT NULL,
    password character varying NOT NULL,
    last_name character varying NOT NULL,
    first_name character varying NOT NULL,
    patronymic character varying,
    id_role integer NOT NULL
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- TOC entry 243 (class 1259 OID 16466)
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_id_seq OWNER TO postgres;

--
-- TOC entry 3529 (class 0 OID 0)
-- Dependencies: 243
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- TOC entry 244 (class 1259 OID 16467)
-- Name: user_id_seq1; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public."user" ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.user_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 3275 (class 2604 OID 16468)
-- Name: accident_character id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_character ALTER COLUMN id SET DEFAULT nextval('public.accident_character_id_seq'::regclass);


--
-- TOC entry 3276 (class 2604 OID 16469)
-- Name: accident_content id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_content ALTER COLUMN id SET DEFAULT nextval('public.accident_content_id_seq'::regclass);


--
-- TOC entry 3277 (class 2604 OID 16470)
-- Name: appeal id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appeal ALTER COLUMN id SET DEFAULT nextval('public.appeal_id_seq'::regclass);


--
-- TOC entry 3279 (class 2604 OID 16471)
-- Name: application id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application ALTER COLUMN id SET DEFAULT nextval('public.application_id_seq'::regclass);


--
-- TOC entry 3280 (class 2604 OID 16472)
-- Name: brigade id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brigade ALTER COLUMN id SET DEFAULT nextval('public.brigade_id_seq'::regclass);


--
-- TOC entry 3281 (class 2604 OID 16473)
-- Name: damage_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.damage_type ALTER COLUMN id SET DEFAULT nextval('public.damage_type_id_seq'::regclass);


--
-- TOC entry 3282 (class 2604 OID 16474)
-- Name: importance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.importance ALTER COLUMN id SET DEFAULT nextval('public.importance_id_seq'::regclass);


--
-- TOC entry 3283 (class 2604 OID 16475)
-- Name: material_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.material_type ALTER COLUMN id SET DEFAULT nextval('public.material_type_id_seq'::regclass);


--
-- TOC entry 3284 (class 2604 OID 16476)
-- Name: role id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role ALTER COLUMN id SET DEFAULT nextval('public.role_id_seq'::regclass);


--
-- TOC entry 3285 (class 2604 OID 16477)
-- Name: shutdown id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shutdown ALTER COLUMN id SET DEFAULT nextval('public.shutdown_id_seq'::regclass);


--
-- TOC entry 3287 (class 2604 OID 16478)
-- Name: specialization id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.specialization ALTER COLUMN id SET DEFAULT nextval('public.specialization_id_seq'::regclass);


--
-- TOC entry 3288 (class 2604 OID 16479)
-- Name: status id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.status ALTER COLUMN id SET DEFAULT nextval('public.status_id_seq'::regclass);


--
-- TOC entry 3483 (class 0 OID 16385)
-- Dependencies: 217
-- Data for Name: accident_character; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.accident_character (id, name, id_specialization) VALUES (1, 'Отсутствие воды', 1);
INSERT INTO public.accident_character (id, name, id_specialization) VALUES (2, 'Утечка воды', 1);
INSERT INTO public.accident_character (id, name, id_specialization) VALUES (3, 'Проблема с давлением', 1);
INSERT INTO public.accident_character (id, name, id_specialization) VALUES (4, 'Засор в системе', 1);
INSERT INTO public.accident_character (id, name, id_specialization) VALUES (5, 'Поломка оборудования', 1);


--
-- TOC entry 3485 (class 0 OID 16391)
-- Dependencies: 219
-- Data for Name: accident_content; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.accident_content (id, name, id_character, recommended_time) VALUES (1, 'Нет воды на улице', 1, 8);
INSERT INTO public.accident_content (id, name, id_character, recommended_time) VALUES (2, 'Нет воды в подъезде', 1, 6);
INSERT INTO public.accident_content (id, name, id_character, recommended_time) VALUES (3, 'Нет воды в квартире', 1, 4);
INSERT INTO public.accident_content (id, name, id_character, recommended_time) VALUES (4, 'Утечка на улице', 2, 12);
INSERT INTO public.accident_content (id, name, id_character, recommended_time) VALUES (5, 'Утечка в подъезде ', 2, 8);
INSERT INTO public.accident_content (id, name, id_character, recommended_time) VALUES (6, 'Низкое давление во всем доме', 3, 24);
INSERT INTO public.accident_content (id, name, id_character, recommended_time) VALUES (7, 'Засор в системе на улице', 4, 10);
INSERT INTO public.accident_content (id, name, id_character, recommended_time) VALUES (8, 'Неисправность насоса в системе', 5, 24);


--
-- TOC entry 3487 (class 0 OID 16397)
-- Dependencies: 221
-- Data for Name: appeal; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3489 (class 0 OID 16404)
-- Dependencies: 223
-- Data for Name: application; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3491 (class 0 OID 16411)
-- Dependencies: 225
-- Data for Name: brigade; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.brigade (id, id_brigadir, people_count) VALUES (1, 3, 10);
INSERT INTO public.brigade (id, id_brigadir, people_count) VALUES (5, 4, 10);
INSERT INTO public.brigade (id, id_brigadir, people_count) VALUES (6, 5, 20);
INSERT INTO public.brigade (id, id_brigadir, people_count) VALUES (7, 7, 30);


--
-- TOC entry 3493 (class 0 OID 16415)
-- Dependencies: 227
-- Data for Name: damage_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.damage_type (id, name) VALUES (1, 'Перелом');
INSERT INTO public.damage_type (id, name) VALUES (2, 'Разрыва');
INSERT INTO public.damage_type (id, name) VALUES (3, 'Свищ');
INSERT INTO public.damage_type (id, name) VALUES (4, 'Трещина');
INSERT INTO public.damage_type (id, name) VALUES (5, 'Коррозия');
INSERT INTO public.damage_type (id, name) VALUES (6, 'Износ');


--
-- TOC entry 3495 (class 0 OID 16421)
-- Dependencies: 229
-- Data for Name: importance; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.importance (id, name) VALUES (1, 'Высокая');
INSERT INTO public.importance (id, name) VALUES (2, 'Средняя');
INSERT INTO public.importance (id, name) VALUES (3, 'Низкая');


--
-- TOC entry 3497 (class 0 OID 16427)
-- Dependencies: 231
-- Data for Name: material_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.material_type (id, name) VALUES (1, 'Сталь');
INSERT INTO public.material_type (id, name) VALUES (2, 'Железобетон');
INSERT INTO public.material_type (id, name) VALUES (3, 'Полиэтилен');
INSERT INTO public.material_type (id, name) VALUES (4, 'ПВХ');


--
-- TOC entry 3499 (class 0 OID 16433)
-- Dependencies: 233
-- Data for Name: role; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.role (id, name) VALUES (1, 'Администратор');
INSERT INTO public.role (id, name) VALUES (3, 'Бригадир');
INSERT INTO public.role (id, name) VALUES (2, 'Оператор');


--
-- TOC entry 3501 (class 0 OID 16439)
-- Dependencies: 235
-- Data for Name: shutdown; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3503 (class 0 OID 16446)
-- Dependencies: 237
-- Data for Name: shutdown_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.shutdown_type (id, name) VALUES (1, 'Горячая');
INSERT INTO public.shutdown_type (id, name) VALUES (2, 'Холодная');


--
-- TOC entry 3504 (class 0 OID 16449)
-- Dependencies: 238
-- Data for Name: specialization; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.specialization (id, name) VALUES (1, 'Водоснабжение');


--
-- TOC entry 3506 (class 0 OID 16455)
-- Dependencies: 240
-- Data for Name: status; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.status (id, name) VALUES (1, 'Новая');
INSERT INTO public.status (id, name) VALUES (2, 'Взята');
INSERT INTO public.status (id, name) VALUES (3, 'В работе');
INSERT INTO public.status (id, name) VALUES (4, 'Завершена');


--
-- TOC entry 3508 (class 0 OID 16461)
-- Dependencies: 242
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) OVERRIDING SYSTEM VALUE VALUES (1, '1', '$2a$12$KrDWOxDt4Yll0BOuPd.U.ODSFlJ8jRHdNCU3nr7wPDrKYORqvMPNy', 'Хурматуллин', 'Булат', 'Саматович', 1);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) OVERRIDING SYSTEM VALUE VALUES (2, '2', '$2a$12$EY.xgUrwMsGxd.zhzlLYve..9j5n/CeKvJM8H/134nVUYzX77Ut1y', 'Хурматуллина', 'Лиана', 'Александровна', 2);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) OVERRIDING SYSTEM VALUE VALUES (3, '3', '$2a$12$XHoDvWwQvoxQVgPB8Xw7S.6h2Q8qLguqS4MyB.FC8F1oYZkLj2O1a', 'Пахомов', 'Александр', 'Александрович', 3);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) OVERRIDING SYSTEM VALUE VALUES (4, '4', '$2a$12$ZVMPZuxRk/rtcF0jpwmQ1.BpjLSVIEj1.C1tZT0AuVuRRMdk69FZW', 'Виктор', 'Волков', 'Павлович', 3);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) OVERRIDING SYSTEM VALUE VALUES (5, '5', '$2a$12$lihkS8vmrfW66n0CSG6g9Og9R05PZ6fjUJTUlN0GN7z3JPQ/ZTele', 'Прохор', 'Тимофеев', 'Геннадьевич', 3);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) OVERRIDING SYSTEM VALUE VALUES (7, '6', '$2a$12$qc71InYiqK3XKKKJbgFBBeYeCy3QJOzkHj11BDXKYEVBLfqjqF09W', 'Владлен', 'Марков', 'Владимирович', 3);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) OVERRIDING SYSTEM VALUE VALUES (8, '7', '$2a$12$IK//pik8IEGN0280LrC40OqjJ6fFIE5kaQYTTCb0XLm/A2kGr4YSW', 'Артур', 'Попов', 'Артемович', 3);


--
-- TOC entry 3530 (class 0 OID 0)
-- Dependencies: 218
-- Name: accident_character_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.accident_character_id_seq', 5, true);


--
-- TOC entry 3531 (class 0 OID 0)
-- Dependencies: 220
-- Name: accident_content_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.accident_content_id_seq', 8, true);


--
-- TOC entry 3532 (class 0 OID 0)
-- Dependencies: 222
-- Name: appeal_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.appeal_id_seq', 1, false);


--
-- TOC entry 3533 (class 0 OID 0)
-- Dependencies: 224
-- Name: application_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.application_id_seq', 1, false);


--
-- TOC entry 3534 (class 0 OID 0)
-- Dependencies: 226
-- Name: brigade_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.brigade_id_seq', 7, true);


--
-- TOC entry 3535 (class 0 OID 0)
-- Dependencies: 228
-- Name: damage_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.damage_type_id_seq', 6, true);


--
-- TOC entry 3536 (class 0 OID 0)
-- Dependencies: 230
-- Name: importance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.importance_id_seq', 3, true);


--
-- TOC entry 3537 (class 0 OID 0)
-- Dependencies: 232
-- Name: material_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.material_type_id_seq', 4, true);


--
-- TOC entry 3538 (class 0 OID 0)
-- Dependencies: 234
-- Name: role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.role_id_seq', 3, true);


--
-- TOC entry 3539 (class 0 OID 0)
-- Dependencies: 236
-- Name: shutdown_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.shutdown_id_seq', 1, false);


--
-- TOC entry 3540 (class 0 OID 0)
-- Dependencies: 239
-- Name: specialization_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.specialization_id_seq', 1, false);


--
-- TOC entry 3541 (class 0 OID 0)
-- Dependencies: 241
-- Name: status_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.status_id_seq', 4, true);


--
-- TOC entry 3542 (class 0 OID 0)
-- Dependencies: 243
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 1, false);


--
-- TOC entry 3543 (class 0 OID 0)
-- Dependencies: 244
-- Name: user_id_seq1; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq1', 8, true);


--
-- TOC entry 3291 (class 2606 OID 16481)
-- Name: accident_character accident_character_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_character
    ADD CONSTRAINT accident_character_pkey PRIMARY KEY (id);


--
-- TOC entry 3293 (class 2606 OID 16483)
-- Name: accident_content accident_content_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_content
    ADD CONSTRAINT accident_content_pkey PRIMARY KEY (id);


--
-- TOC entry 3295 (class 2606 OID 16485)
-- Name: appeal appeal_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appeal
    ADD CONSTRAINT appeal_pkey PRIMARY KEY (id);


--
-- TOC entry 3297 (class 2606 OID 16487)
-- Name: application application_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_pkey PRIMARY KEY (id);


--
-- TOC entry 3299 (class 2606 OID 16489)
-- Name: brigade brigade_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brigade
    ADD CONSTRAINT brigade_pkey PRIMARY KEY (id);


--
-- TOC entry 3301 (class 2606 OID 16491)
-- Name: damage_type damage_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.damage_type
    ADD CONSTRAINT damage_type_pkey PRIMARY KEY (id);


--
-- TOC entry 3303 (class 2606 OID 16493)
-- Name: importance importance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.importance
    ADD CONSTRAINT importance_pkey PRIMARY KEY (id);


--
-- TOC entry 3305 (class 2606 OID 16495)
-- Name: material_type material_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.material_type
    ADD CONSTRAINT material_type_pkey PRIMARY KEY (id);


--
-- TOC entry 3307 (class 2606 OID 16497)
-- Name: role role_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT role_pkey PRIMARY KEY (id);


--
-- TOC entry 3309 (class 2606 OID 16499)
-- Name: shutdown shutdown_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shutdown
    ADD CONSTRAINT shutdown_pkey PRIMARY KEY (id);


--
-- TOC entry 3311 (class 2606 OID 16501)
-- Name: shutdown_type shutdown_type_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shutdown_type
    ADD CONSTRAINT shutdown_type_name_key UNIQUE (name);


--
-- TOC entry 3313 (class 2606 OID 16503)
-- Name: shutdown_type shutdown_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shutdown_type
    ADD CONSTRAINT shutdown_type_pkey PRIMARY KEY (id);


--
-- TOC entry 3315 (class 2606 OID 16505)
-- Name: specialization specialization_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.specialization
    ADD CONSTRAINT specialization_pkey PRIMARY KEY (id);


--
-- TOC entry 3317 (class 2606 OID 16507)
-- Name: status status_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.status
    ADD CONSTRAINT status_pkey PRIMARY KEY (id);


--
-- TOC entry 3319 (class 2606 OID 16509)
-- Name: user user_login_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_login_key UNIQUE (login);


--
-- TOC entry 3321 (class 2606 OID 16511)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 3322 (class 2606 OID 16512)
-- Name: accident_character fk_accident_character_specialization; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_character
    ADD CONSTRAINT fk_accident_character_specialization FOREIGN KEY (id_specialization) REFERENCES public.specialization(id);


--
-- TOC entry 3323 (class 2606 OID 16517)
-- Name: accident_content fk_accident_content_character; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_content
    ADD CONSTRAINT fk_accident_content_character FOREIGN KEY (id_character) REFERENCES public.accident_character(id);


--
-- TOC entry 3324 (class 2606 OID 16522)
-- Name: appeal fk_appeal_accident; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appeal
    ADD CONSTRAINT fk_appeal_accident FOREIGN KEY (id_accident) REFERENCES public.accident_content(id);


--
-- TOC entry 3325 (class 2606 OID 16527)
-- Name: appeal fk_appeal_application; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appeal
    ADD CONSTRAINT fk_appeal_application FOREIGN KEY (id_application) REFERENCES public.application(id) ON DELETE CASCADE;


--
-- TOC entry 3326 (class 2606 OID 16532)
-- Name: application fk_application_accident; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_accident FOREIGN KEY (id_accident) REFERENCES public.accident_content(id);


--
-- TOC entry 3327 (class 2606 OID 16537)
-- Name: application fk_application_brigade; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_brigade FOREIGN KEY (id_brigade) REFERENCES public.brigade(id) NOT VALID;


--
-- TOC entry 3328 (class 2606 OID 16542)
-- Name: application fk_application_damage; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_damage FOREIGN KEY (id_damage) REFERENCES public.damage_type(id);


--
-- TOC entry 3329 (class 2606 OID 16547)
-- Name: application fk_application_importance; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_importance FOREIGN KEY (id_importance) REFERENCES public.importance(id);


--
-- TOC entry 3330 (class 2606 OID 16552)
-- Name: application fk_application_material; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_material FOREIGN KEY (id_material) REFERENCES public.material_type(id);


--
-- TOC entry 3331 (class 2606 OID 16557)
-- Name: application fk_application_operator; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_operator FOREIGN KEY (id_operator) REFERENCES public."user"(id) ON DELETE SET NULL NOT VALID;


--
-- TOC entry 3332 (class 2606 OID 16562)
-- Name: application fk_application_status; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_status FOREIGN KEY (id_status) REFERENCES public.status(id);


--
-- TOC entry 3333 (class 2606 OID 16567)
-- Name: brigade fk_brigade_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brigade
    ADD CONSTRAINT fk_brigade_user FOREIGN KEY (id_brigadir) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3334 (class 2606 OID 16572)
-- Name: shutdown fk_shutdown_accident; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shutdown
    ADD CONSTRAINT fk_shutdown_accident FOREIGN KEY (id_accident) REFERENCES public.accident_content(id);


--
-- TOC entry 3335 (class 2606 OID 16577)
-- Name: shutdown fk_shutdown_application; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shutdown
    ADD CONSTRAINT fk_shutdown_application FOREIGN KEY (id_application) REFERENCES public.application(id) ON DELETE CASCADE;


--
-- TOC entry 3337 (class 2606 OID 16582)
-- Name: user fk_user_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT fk_user_role FOREIGN KEY (id_role) REFERENCES public.role(id);


--
-- TOC entry 3336 (class 2606 OID 16587)
-- Name: shutdown shutdown_id_type_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shutdown
    ADD CONSTRAINT shutdown_id_type_fkey FOREIGN KEY (id_type) REFERENCES public.shutdown_type(id);


-- Completed on 2025-06-23 17:08:10 UTC

--
-- PostgreSQL database dump complete
--

