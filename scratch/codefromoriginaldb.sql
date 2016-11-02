-- Table: public.cities

-- DROP TABLE public.cities;

CREATE TABLE public.cities
(
    city character varying(80) COLLATE pg_catalog."default" NOT NULL,
    location point,
    CONSTRAINT cities_pkey PRIMARY KEY (city)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.cities
    OWNER to qrqgaffpqojjtm;



-- Table: public.tasks

-- DROP TABLE public.tasks;

CREATE TABLE public.tasks
(
    id integer NOT NULL DEFAULT nextval('tasks_id_seq'::regclass),
    task character varying(225) COLLATE pg_catalog."default",
    companies_id character varying(225) COLLATE pg_catalog."default",
    stores_id character varying(225) COLLATE pg_catalog."default",
    weighting character varying(225) COLLATE pg_catalog."default",
    task_types character varying(225) COLLATE pg_catalog."default",
    frequency character varying(225) COLLATE pg_catalog."default",
    time_when character varying(225) COLLATE pg_catalog."default",
    time_offset character varying(225) COLLATE pg_catalog."default",
    time_float character varying(225) COLLATE pg_catalog."default",
    time_length character varying(225) COLLATE pg_catalog."default",
    person_per_flow character varying(225) COLLATE pg_catalog."default",
    person_per_volume character varying(225) COLLATE pg_catalog."default",
    time_max character varying(225) COLLATE pg_catalog."default",
    time_min character varying(225) COLLATE pg_catalog."default",
    day_of_week character varying(225) COLLATE pg_catalog."default",
    active character varying(225) COLLATE pg_catalog."default",
    updated_at character varying(225) COLLATE pg_catalog."default",
    created_at character varying(225) COLLATE pg_catalog."default",
    CONSTRAINT id PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.tasks
    OWNER to qrqgaffpqojjtm;




-- Table: public.taskstafflink

-- DROP TABLE public.taskstafflink;

CREATE TABLE public.taskstafflink
(
    id integer NOT NULL DEFAULT nextval('taskstafflink_id_seq'::regclass),
    staffid integer,
    taskid integer,
    stores_id integer,
    companies_id integer,
    created_on timestamp without time zone DEFAULT now(),
    updated_on timestamp without time zone DEFAULT now(),
    CONSTRAINT taskstafflink_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.taskstafflink
    OWNER to qrqgaffpqojjtm;



