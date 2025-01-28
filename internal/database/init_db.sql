--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2 (Debian 17.2-1.pgdg120+1)
-- Dumped by pg_dump version 17.2

-- Started on 2025-01-10 15:34:10 UTC

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



ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3501 (class 0 OID 0)
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
-- TOC entry 3502 (class 0 OID 0)
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
    id_character integer NOT NULL
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
-- TOC entry 3503 (class 0 OID 0)
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
    is_active boolean
);


ALTER TABLE public.appeal OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 16402)
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
-- TOC entry 3504 (class 0 OID 0)
-- Dependencies: 222
-- Name: appeal_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.appeal_id_seq OWNED BY public.appeal.id;


--
-- TOC entry 223 (class 1259 OID 16403)
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
    id_brigade integer
);


ALTER TABLE public.application OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 16408)
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
-- TOC entry 3505 (class 0 OID 0)
-- Dependencies: 224
-- Name: application_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.application_id_seq OWNED BY public.application.id;


--
-- TOC entry 242 (class 1259 OID 16561)
-- Name: brigade; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.brigade (
    id integer NOT NULL,
    id_brigadir integer,
    people_count integer
);


ALTER TABLE public.brigade OWNER TO postgres;

--
-- TOC entry 241 (class 1259 OID 16560)
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
-- TOC entry 3506 (class 0 OID 0)
-- Dependencies: 241
-- Name: brigade_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.brigade_id_seq OWNED BY public.brigade.id;


--
-- TOC entry 225 (class 1259 OID 16409)
-- Name: damage_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.damage_type (
    id integer NOT NULL,
    name character varying
);


ALTER TABLE public.damage_type OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 16414)
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
-- TOC entry 3507 (class 0 OID 0)
-- Dependencies: 226
-- Name: damage_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.damage_type_id_seq OWNED BY public.damage_type.id;


--
-- TOC entry 227 (class 1259 OID 16415)
-- Name: importance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.importance (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.importance OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 16420)
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
-- TOC entry 3508 (class 0 OID 0)
-- Dependencies: 228
-- Name: importance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.importance_id_seq OWNED BY public.importance.id;


--
-- TOC entry 229 (class 1259 OID 16421)
-- Name: material_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.material_type (
    id integer NOT NULL,
    name character varying
);


ALTER TABLE public.material_type OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 16426)
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
-- TOC entry 3509 (class 0 OID 0)
-- Dependencies: 230
-- Name: material_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.material_type_id_seq OWNED BY public.material_type.id;


--
-- TOC entry 231 (class 1259 OID 16427)
-- Name: role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.role OWNER TO postgres;

--
-- TOC entry 232 (class 1259 OID 16432)
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
-- TOC entry 3510 (class 0 OID 0)
-- Dependencies: 232
-- Name: role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.role_id_seq OWNED BY public.role.id;


--
-- TOC entry 233 (class 1259 OID 16433)
-- Name: shutdown; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shutdown (
    id integer NOT NULL,
    address text,
    id_accident integer,
    date timestamp without time zone,
    is_active boolean,
    day_count integer,
    id_application integer,
    description text
);


ALTER TABLE public.shutdown OWNER TO postgres;

--
-- TOC entry 234 (class 1259 OID 16438)
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
-- TOC entry 3511 (class 0 OID 0)
-- Dependencies: 234
-- Name: shutdown_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.shutdown_id_seq OWNED BY public.shutdown.id;


--
-- TOC entry 235 (class 1259 OID 16439)
-- Name: specialization; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.specialization (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.specialization OWNER TO postgres;

--
-- TOC entry 236 (class 1259 OID 16444)
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
-- TOC entry 3512 (class 0 OID 0)
-- Dependencies: 236
-- Name: specialization_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.specialization_id_seq OWNED BY public.specialization.id;


--
-- TOC entry 237 (class 1259 OID 16445)
-- Name: status; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.status (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.status OWNER TO postgres;

--
-- TOC entry 238 (class 1259 OID 16450)
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
-- TOC entry 3513 (class 0 OID 0)
-- Dependencies: 238
-- Name: status_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.status_id_seq OWNED BY public.status.id;


--
-- TOC entry 239 (class 1259 OID 16451)
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
-- TOC entry 240 (class 1259 OID 16456)
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
-- TOC entry 3514 (class 0 OID 0)
-- Dependencies: 240
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- TOC entry 3270 (class 2604 OID 16457)
-- Name: accident_character id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_character ALTER COLUMN id SET DEFAULT nextval('public.accident_character_id_seq'::regclass);


--
-- TOC entry 3271 (class 2604 OID 16458)
-- Name: accident_content id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_content ALTER COLUMN id SET DEFAULT nextval('public.accident_content_id_seq'::regclass);


--
-- TOC entry 3272 (class 2604 OID 16459)
-- Name: appeal id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appeal ALTER COLUMN id SET DEFAULT nextval('public.appeal_id_seq'::regclass);


--
-- TOC entry 3273 (class 2604 OID 16460)
-- Name: application id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application ALTER COLUMN id SET DEFAULT nextval('public.application_id_seq'::regclass);


--
-- TOC entry 3282 (class 2604 OID 16564)
-- Name: brigade id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brigade ALTER COLUMN id SET DEFAULT nextval('public.brigade_id_seq'::regclass);


--
-- TOC entry 3274 (class 2604 OID 16461)
-- Name: damage_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.damage_type ALTER COLUMN id SET DEFAULT nextval('public.damage_type_id_seq'::regclass);


--
-- TOC entry 3275 (class 2604 OID 16462)
-- Name: importance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.importance ALTER COLUMN id SET DEFAULT nextval('public.importance_id_seq'::regclass);


--
-- TOC entry 3276 (class 2604 OID 16463)
-- Name: material_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.material_type ALTER COLUMN id SET DEFAULT nextval('public.material_type_id_seq'::regclass);


--
-- TOC entry 3277 (class 2604 OID 16464)
-- Name: role id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role ALTER COLUMN id SET DEFAULT nextval('public.role_id_seq'::regclass);


--
-- TOC entry 3278 (class 2604 OID 16465)
-- Name: shutdown id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shutdown ALTER COLUMN id SET DEFAULT nextval('public.shutdown_id_seq'::regclass);


--
-- TOC entry 3279 (class 2604 OID 16466)
-- Name: specialization id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.specialization ALTER COLUMN id SET DEFAULT nextval('public.specialization_id_seq'::regclass);


--
-- TOC entry 3280 (class 2604 OID 16467)
-- Name: status id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.status ALTER COLUMN id SET DEFAULT nextval('public.status_id_seq'::regclass);


--
-- TOC entry 3281 (class 2604 OID 16468)
-- Name: user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- TOC entry 3470 (class 0 OID 16385)
-- Dependencies: 217
-- Data for Name: accident_character; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.accident_character (id, name, id_specialization) VALUES (1, 'Отсутствие воды', 1);
INSERT INTO public.accident_character (id, name, id_specialization) VALUES (2, 'Утечка воды', 1);
INSERT INTO public.accident_character (id, name, id_specialization) VALUES (3, 'Проблемы с давлением', 1);
INSERT INTO public.accident_character (id, name, id_specialization) VALUES (4, 'Засор в системе', 1);
INSERT INTO public.accident_character (id, name, id_specialization) VALUES (5, 'Поломка оборудования', 1);


--
-- TOC entry 3472 (class 0 OID 16391)
-- Dependencies: 219
-- Data for Name: accident_content; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.accident_content (id, name, id_character) VALUES (1, 'Нет воды на улице', 1);
INSERT INTO public.accident_content (id, name, id_character) VALUES (2, 'Нет воды в подъезде', 1);
INSERT INTO public.accident_content (id, name, id_character) VALUES (3, 'Нет воды в квартире', 1);
INSERT INTO public.accident_content (id, name, id_character) VALUES (4, 'Утечка на улице', 2);
INSERT INTO public.accident_content (id, name, id_character) VALUES (5, 'Утечка в подъезде', 2);
INSERT INTO public.accident_content (id, name, id_character) VALUES (6, 'Низкое давление во всем доме', 3);
INSERT INTO public.accident_content (id, name, id_character) VALUES (7, 'Засор в системе на улице', 4);
INSERT INTO public.accident_content (id, name, id_character) VALUES (8, 'Неисправность насоса в системе', 5);


--
-- TOC entry 3474 (class 0 OID 16397)
-- Dependencies: 221
-- Data for Name: appeal; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (10, '', '2024-11-22 13:40:34.784', 4, 'Шамсутдинов Валерий Малонович', '8 (989) 389-38-93', 13, 'Казань,Ленская улица,10', '8 (394) 398-43-98', true);
INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (11, 'Не работает система', '2024-11-22 13:41:34.422', 7, 'Ленов Сергей Павлович', '+7 (989) 898-98-99', 14, 'Казань,улица Богатырёва,5к2', '+7 (989) 898-90-89', true);
INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (9, '', '2024-11-22 13:39:11.764', 1, 'Валенок Имиль Валерьевич', '8 (589) 595-30-93', NULL, 'Казань,Архангельская улица,9', '8 (598) 985-98-98', true);
INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (7, '', '2024-11-22 13:34:11.775', 1, 'Алонова Елизавета Петровна', '+7 (902) 902-93-02', 10, 'Казань,Полевая улица,30', '+7 (984) 938-04-83', true);
INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (16, 'Утечка воды в магазине', '2024-11-23 07:32:21.687', 4, 'Скрылева Лиана Александровна', '+7 (999) 188-51-23', 24, 'Казань,Беломорская улица,81', '+7 (933) 568-33-11', true);
INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (17, 'пашок', '2024-11-23 09:04:43.032', 4, 'пашок ебаний эжи', '+7 (923) 992-49-24', 26, 'посёлок Кизнер,Кизнерская улица,82', '+55454554552222222', true);
INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (8, 'Насос не работает', '2024-11-22 13:35:39.07', 8, 'Ларинов Эмиль Евгеньевич', '8 (787) 484-78-78', NULL, 'Казань,Октябрьская улица,5А', '+7 (984) 849-84-98', true);
INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (1, 'В 26 квартире нет воды', '2023-11-21 18:26:29.112', 3, 'Хурматуллин Булат Саматович', '+7 (999) 123-14-45', NULL, 'Казань,улица Ульянова-Ленина,50', '', false);
INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (3, '', '2024-11-22 08:44:54.186', 2, 'Логинова Юлия Александровна', '8 (999) 999-99-99', 5, 'Казань,улица Коротченко,2', '', false);
INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (2, '', '2024-11-22 08:44:03.631', 6, 'Кумышбаева София Алексеевна', '8 (999) 999-99-99', 16, 'Казань,улица Батурина,5', '+7 (999) 999-99-99', false);
INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (15, 'В магазине Дом Обоев', '2024-11-23 06:42:44.339', 2, 'Скрылева Лиана Алекандровна', '+7 (993) 849-38-49', 21, 'Казань,улица Химиков,13', '', false);
INSERT INTO public.appeal (id, description, create_date, id_accident, applicant_name, applicant_number, id_application, address, additional_number, is_active) VALUES (14, 'Нет воды на улице', '2024-11-22 19:48:15.721', 7, 'Хурматуллин Булат Саматович', '+7 (999) 999-99-99', 21, 'Казань,улица Бари Галеева,3А', '8 (787) 989-89-89', false);


--
-- TOC entry 3476 (class 0 OID 16403)
-- Dependencies: 223
-- Data for Name: application; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (22, 'В магазине Дом Обоев', '2024-11-23 06:42:56.544', 2, 2, 1, 'Казань,улица Химиков,13', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (23, 'Нет воды в 3 подъезде.', '2024-11-23 06:47:55.058', 2, 3, 1, 'Казань,Поперечно-Авангардная улица,5лит3', 'Прорыв трубы', 'Правый сектор', 1, 2, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (24, 'Утечка воды в магазине', '2024-11-23 07:32:21.669', 4, 3, 1, 'Казань,Беломорская улица,81', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (25, 'Вытекает во втором подъезде', '2024-11-23 07:34:22.728', 5, 1, 1, 'Казань,Авангардная улица,167А', 'Высокое давление', 'Правый сектор', 1, 2, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (26, 'пашок', '2024-11-23 09:05:27.194', 4, 1, 1, 'посёлок Кизнер,Кизнерская улица,82', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (6, 'В магазинах нет подачи воды', '2024-11-22 08:48:15.402', 1, 1, 1, 'Казань,улица Гладилова,32', 'Большое давление', 'Улица', 1, 1, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (7, '', '2024-11-22 09:22:23.062', 5, 1, 3, 'Казань,улица Баумана,1К2', '', '', 2, 2, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (5, '', '2024-11-22 08:44:54.169', 2, 2, 3, 'Казань,улица Коротченко,2', '', '', 1, 1, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (9, '', '2024-11-22 13:30:04.363', 6, 3, 1, 'Казань,проспект Ибрагимова,54', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (10, '', '2024-11-22 13:34:11.759', 1, 2, 1, 'Казань,Полевая улица,30', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (11, 'Насос не работает', '2024-11-22 13:35:39.056', 8, 1, 1, 'Казань,Октябрьская улица,5А', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (12, '', '2024-11-22 13:39:11.747', 1, 3, 1, 'Казань,Архангельская улица,9', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (13, '', '2024-11-22 13:40:34.766', 4, 2, 1, 'Казань,Ленская улица,10', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (14, 'Не работает система', '2024-11-22 13:41:34.405', 7, 3, 1, 'Казань,улица Богатырёва,5к2', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (16, '', '2024-11-22 18:43:31.837', 6, 3, 3, 'Казань,улица Батурина,5', '', '', NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (18, 'Вода вытекает во дворе. Около песочницы', '2024-11-22 19:00:01.718', 4, 1, 1, 'Казань,Авангардная улица,167А', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (19, 'В школе не работает краны', '2024-11-22 19:15:57.129', 3, 1, 2, 'Казань,Дубравная улица,35А', 'Высокое давление в системе', 'Двор', 2, 2, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (20, 'Нет воды на улице', '2024-11-22 19:48:26.167', 7, 3, 1, 'Казань,улица Бари Галеева,3А', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (3, 'Насос не работает', '2024-10-15 17:10:30', 8, 2, 3, 'Казань,Авангардная улица,167', '', '', 3, 1, NULL);
INSERT INTO public.application (id, description, create_date, id_accident, id_importance, id_status, address, accident_cause, damage_point, id_material, id_damage, id_brigade) VALUES (21, 'Нет воды на улице', '2024-11-22 19:50:44.707', 7, 1, 3, 'Казань,улица Бари Галеева,3А', 'Не аккуратное испо', 'Правый изгиб', 2, 5, NULL);


--
-- TOC entry 3495 (class 0 OID 16561)
-- Dependencies: 242
-- Data for Name: brigade; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3478 (class 0 OID 16409)
-- Dependencies: 225
-- Data for Name: damage_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.damage_type (id, name) VALUES (1, 'Перелом');
INSERT INTO public.damage_type (id, name) VALUES (2, 'Разрыв');
INSERT INTO public.damage_type (id, name) VALUES (3, 'Свищ');
INSERT INTO public.damage_type (id, name) VALUES (4, 'Трещина');
INSERT INTO public.damage_type (id, name) VALUES (5, 'Коррозия');
INSERT INTO public.damage_type (id, name) VALUES (6, 'Износ');


--
-- TOC entry 3480 (class 0 OID 16415)
-- Dependencies: 227
-- Data for Name: importance; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.importance (id, name) VALUES (1, 'Высокая');
INSERT INTO public.importance (id, name) VALUES (2, 'Средняя');
INSERT INTO public.importance (id, name) VALUES (3, 'Низкая');


--
-- TOC entry 3482 (class 0 OID 16421)
-- Dependencies: 229
-- Data for Name: material_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.material_type (id, name) VALUES (1, 'Сталь');
INSERT INTO public.material_type (id, name) VALUES (2, 'Железобетон');
INSERT INTO public.material_type (id, name) VALUES (3, 'Полиэтилен');
INSERT INTO public.material_type (id, name) VALUES (4, 'ПВХ');


--
-- TOC entry 3484 (class 0 OID 16427)
-- Dependencies: 231
-- Data for Name: role; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.role (id, name) VALUES (1, 'Администратор');
INSERT INTO public.role (id, name) VALUES (2, 'Оператор');
INSERT INTO public.role (id, name) VALUES (3, 'Бригадир');


--
-- TOC entry 3486 (class 0 OID 16433)
-- Dependencies: 233
-- Data for Name: shutdown; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.shutdown (id, address, id_accident, date, is_active, day_count, id_application, description) VALUES (14, 'Казань,улица Батурина,7А', 1, '2024-11-26 00:00:00', false, 3, NULL, '');
INSERT INTO public.shutdown (id, address, id_accident, date, is_active, day_count, id_application, description) VALUES (6, 'Казань,проспект Ямашева,1', 6, '2024-11-28 00:00:00', false, 1, NULL, '');
INSERT INTO public.shutdown (id, address, id_accident, date, is_active, day_count, id_application, description) VALUES (4, 'Казань,Беломорская улица,81', 1, '2024-11-23 00:00:00', false, 5, 3, '');
INSERT INTO public.shutdown (id, address, id_accident, date, is_active, day_count, id_application, description) VALUES (7, 'Казань,улица Бари Галеева,4А', 3, '2024-11-23 00:00:00', false, 5, NULL, '');
INSERT INTO public.shutdown (id, address, id_accident, date, is_active, day_count, id_application, description) VALUES (15, 'Казань,Беломорская улица,81', 4, '2024-11-23 00:00:00', false, 4, 3, 'Необходим отключить горячую воду');
INSERT INTO public.shutdown (id, address, id_accident, date, is_active, day_count, id_application, description) VALUES (9, 'Казань,улица Алексея Козина,2', 1, '2024-11-22 00:00:00', false, 5, NULL, '');
INSERT INTO public.shutdown (id, address, id_accident, date, is_active, day_count, id_application, description) VALUES (10, 'Казань,проспект Ямашева,93', 3, '2024-11-23 00:00:00', false, 4, NULL, '');
INSERT INTO public.shutdown (id, address, id_accident, date, is_active, day_count, id_application, description) VALUES (11, 'Казань,улица Адоратского,7', 6, '2024-11-28 00:00:00', false, 2, 3, '');
INSERT INTO public.shutdown (id, address, id_accident, date, is_active, day_count, id_application, description) VALUES (12, 'Казань,Коллективная улица,39', 6, '2024-11-26 00:00:00', false, 3, NULL, '');
INSERT INTO public.shutdown (id, address, id_accident, date, is_active, day_count, id_application, description) VALUES (1, 'Казань,проспект Победы,141', 7, '2024-11-12 00:00:00', false, 9, 3, 'Отключение ');
INSERT INTO public.shutdown (id, address, id_accident, date, is_active, day_count, id_application, description) VALUES (2, 'Казань,улица Вишневского,21', 3, '2024-11-07 00:00:00', false, 2, NULL, '');


--
-- TOC entry 3488 (class 0 OID 16439)
-- Dependencies: 235
-- Data for Name: specialization; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.specialization (id, name) VALUES (1, 'Водоснабжение');


--
-- TOC entry 3490 (class 0 OID 16445)
-- Dependencies: 237
-- Data for Name: status; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.status (id, name) VALUES (1, 'Новая');
INSERT INTO public.status (id, name) VALUES (2, 'В работе');
INSERT INTO public.status (id, name) VALUES (3, 'Завершена');


--
-- TOC entry 3492 (class 0 OID 16451)
-- Dependencies: 239
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) VALUES (1, 'adminUrban', '$2a$12$RVI/9df.QlBOz4.4i2ldiemnw3kJRtG8tK3y7dQAnSP0z4186cf/6', 'Скрылева', 'Лиана', 'Александровна', 1);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) VALUES (7, 'kulD', '$2a$12$pvlFlDqP/47JAGd57HCoPu5Wj2Bx.p0JRD.r64owUYkwdMk/TKkM2', 'Куликов', 'Дмитрий', 'Николаевич', 2);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) VALUES (8, 'nikA', '$2a$12$T.l81/HVq1gPStQZ1dO6ruKvaZIIKNjFEiLVChAZzpLebauVcmePe', 'Акифьев', 'Никита', 'Олегович', 2);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) VALUES (9, 'albS', '$2a$12$aldH.CGFPm.5fjw4coAILOF0W0QBJuqT6Yx4LSblsAWB9xBuWeir6', 'Скрылева', 'Альбина', 'Александровна', 2);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) VALUES (10, 'ivan12', '$2a$12$hHGYZpz5vsb5TPJeDkhHPeI.Ib6jBpUBPG5U/FqRDAD9C9wJviNuS', 'Раричев', 'Иван', 'Николаевич', 2);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) VALUES (11, 'evg13', '$2a$12$0v7p0UUmcUKrdFzofoRSiuXFsCNsH4OdmYElrX0s6w..BG49npRkK', 'Мамонов', 'Евгений', 'Васильевич', 2);
INSERT INTO public."user" (id, login, password, last_name, first_name, patronymic, id_role) VALUES (12, 'bulat', '$2a$12$l4xvRVCPzNskmmkbD/vyM.FzHRsPZsIlV.FkTo21MPR2evlcTpqCK', 'Хурматуллин ', 'Булат', 'Саматович', 2);


--
-- TOC entry 3515 (class 0 OID 0)
-- Dependencies: 218
-- Name: accident_character_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.accident_character_id_seq', 1, false);


--
-- TOC entry 3516 (class 0 OID 0)
-- Dependencies: 220
-- Name: accident_content_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.accident_content_id_seq', 1, false);


--
-- TOC entry 3517 (class 0 OID 0)
-- Dependencies: 222
-- Name: appeal_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.appeal_id_seq', 17, true);


--
-- TOC entry 3518 (class 0 OID 0)
-- Dependencies: 224
-- Name: application_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.application_id_seq', 27, true);


--
-- TOC entry 3519 (class 0 OID 0)
-- Dependencies: 241
-- Name: brigade_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.brigade_id_seq', 1, false);


--
-- TOC entry 3520 (class 0 OID 0)
-- Dependencies: 226
-- Name: damage_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.damage_type_id_seq', 6, true);


--
-- TOC entry 3521 (class 0 OID 0)
-- Dependencies: 228
-- Name: importance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.importance_id_seq', 3, true);


--
-- TOC entry 3522 (class 0 OID 0)
-- Dependencies: 230
-- Name: material_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.material_type_id_seq', 4, true);


--
-- TOC entry 3523 (class 0 OID 0)
-- Dependencies: 232
-- Name: role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.role_id_seq', 3, true);


--
-- TOC entry 3524 (class 0 OID 0)
-- Dependencies: 234
-- Name: shutdown_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.shutdown_id_seq', 16, true);


--
-- TOC entry 3525 (class 0 OID 0)
-- Dependencies: 236
-- Name: specialization_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.specialization_id_seq', 1, true);


--
-- TOC entry 3526 (class 0 OID 0)
-- Dependencies: 238
-- Name: status_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.status_id_seq', 3, true);


--
-- TOC entry 3527 (class 0 OID 0)
-- Dependencies: 240
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 12, true);


--
-- TOC entry 3284 (class 2606 OID 16470)
-- Name: accident_character accident_character_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_character
    ADD CONSTRAINT accident_character_pkey PRIMARY KEY (id);


--
-- TOC entry 3286 (class 2606 OID 16472)
-- Name: accident_content accident_content_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_content
    ADD CONSTRAINT accident_content_pkey PRIMARY KEY (id);


--
-- TOC entry 3288 (class 2606 OID 16474)
-- Name: appeal appeal_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appeal
    ADD CONSTRAINT appeal_pkey PRIMARY KEY (id);


--
-- TOC entry 3290 (class 2606 OID 16476)
-- Name: application application_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT application_pkey PRIMARY KEY (id);


--
-- TOC entry 3310 (class 2606 OID 16566)
-- Name: brigade brigade_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brigade
    ADD CONSTRAINT brigade_pkey PRIMARY KEY (id);


--
-- TOC entry 3292 (class 2606 OID 16478)
-- Name: damage_type damage_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.damage_type
    ADD CONSTRAINT damage_type_pkey PRIMARY KEY (id);


--
-- TOC entry 3294 (class 2606 OID 16480)
-- Name: importance importance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.importance
    ADD CONSTRAINT importance_pkey PRIMARY KEY (id);


--
-- TOC entry 3296 (class 2606 OID 16482)
-- Name: material_type material_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.material_type
    ADD CONSTRAINT material_type_pkey PRIMARY KEY (id);


--
-- TOC entry 3298 (class 2606 OID 16484)
-- Name: role role_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT role_pkey PRIMARY KEY (id);


--
-- TOC entry 3300 (class 2606 OID 16486)
-- Name: shutdown shutdown_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shutdown
    ADD CONSTRAINT shutdown_pkey PRIMARY KEY (id);


--
-- TOC entry 3302 (class 2606 OID 16488)
-- Name: specialization specialization_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.specialization
    ADD CONSTRAINT specialization_pkey PRIMARY KEY (id);


--
-- TOC entry 3304 (class 2606 OID 16490)
-- Name: status status_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.status
    ADD CONSTRAINT status_pkey PRIMARY KEY (id);


--
-- TOC entry 3306 (class 2606 OID 16492)
-- Name: user user_login_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_login_key UNIQUE (login);


--
-- TOC entry 3308 (class 2606 OID 16494)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 3311 (class 2606 OID 16495)
-- Name: accident_character fk_accident_character_specialization; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_character
    ADD CONSTRAINT fk_accident_character_specialization FOREIGN KEY (id_specialization) REFERENCES public.specialization(id);


--
-- TOC entry 3312 (class 2606 OID 16500)
-- Name: accident_content fk_accident_content_character; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accident_content
    ADD CONSTRAINT fk_accident_content_character FOREIGN KEY (id_character) REFERENCES public.accident_character(id);


--
-- TOC entry 3313 (class 2606 OID 16505)
-- Name: appeal fk_appeal_accident; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appeal
    ADD CONSTRAINT fk_appeal_accident FOREIGN KEY (id_accident) REFERENCES public.accident_content(id);


--
-- TOC entry 3314 (class 2606 OID 16510)
-- Name: appeal fk_appeal_application; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appeal
    ADD CONSTRAINT fk_appeal_application FOREIGN KEY (id_application) REFERENCES public.application(id) ON DELETE CASCADE;


--
-- TOC entry 3315 (class 2606 OID 16515)
-- Name: application fk_application_accident; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_accident FOREIGN KEY (id_accident) REFERENCES public.accident_content(id);


--
-- TOC entry 3316 (class 2606 OID 16572)
-- Name: application fk_application_brigade; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_brigade FOREIGN KEY (id_brigade) REFERENCES public.brigade(id) NOT VALID;


--
-- TOC entry 3317 (class 2606 OID 16520)
-- Name: application fk_application_damage; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_damage FOREIGN KEY (id_damage) REFERENCES public.damage_type(id);


--
-- TOC entry 3318 (class 2606 OID 16525)
-- Name: application fk_application_importance; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_importance FOREIGN KEY (id_importance) REFERENCES public.importance(id);


--
-- TOC entry 3319 (class 2606 OID 16530)
-- Name: application fk_application_material; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_material FOREIGN KEY (id_material) REFERENCES public.material_type(id);


--
-- TOC entry 3320 (class 2606 OID 16535)
-- Name: application fk_application_status; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.application
    ADD CONSTRAINT fk_application_status FOREIGN KEY (id_status) REFERENCES public.status(id);


--
-- TOC entry 3324 (class 2606 OID 16567)
-- Name: brigade fk_brigade_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brigade
    ADD CONSTRAINT fk_brigade_user FOREIGN KEY (id_brigadir) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3321 (class 2606 OID 16540)
-- Name: shutdown fk_shutdown_accident; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shutdown
    ADD CONSTRAINT fk_shutdown_accident FOREIGN KEY (id_accident) REFERENCES public.accident_content(id);


--
-- TOC entry 3322 (class 2606 OID 16545)
-- Name: shutdown fk_shutdown_application; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shutdown
    ADD CONSTRAINT fk_shutdown_application FOREIGN KEY (id_application) REFERENCES public.application(id) ON DELETE CASCADE;


--
-- TOC entry 3323 (class 2606 OID 16550)
-- Name: user fk_user_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT fk_user_role FOREIGN KEY (id_role) REFERENCES public.role(id);


-- Completed on 2025-01-10 15:34:10 UTC

--
-- PostgreSQL database dump complete
--

