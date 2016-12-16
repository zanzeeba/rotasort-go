-- staff
INSERT INTO staff (email, username, password, phone, address, postcode, firstname, lastname, companies_id, stores_id, dept_id, siteowner, active, breaks_id, created_on, updated_on) VALUES 
('fb@sophiadex.com', 'fb', 'password', 'phone', 'address', 'postcode', 'fred', 'blogs', 1, 1, 1, true, true, 0, '2016-11-05','2016-11-05'),
('aib@sophiadex.com', 'aib', 'password', 'phone', 'address', 'postcode', 'ada', 'blogs', 1, 1, 1, true, true,  0, '2016-11-05','2016-11-05'),
('jb@sophiadex.com', 'jb', 'password','phone', 'address', 'postcode', 'josephine', 'blogs', 1, 1, 1, true, true,  0, '2016-11-05','2016-11-05'),
('sb@sophiadex.com', 'sb', 'password','phone', 'address', 'postcode', 'sydney', 'blogs', 1, 1, 1, true, true,  0, '2016-11-05','2016-11-05'),
('jar@sophiadex.com', 'jar', 'password','phone', 'address', 'postcode', 'jason', 'rankin', 1, 1, 1, true, true,  0, '2016-11-05','2016-11-05'),
('jr@sophiadex.com', 'jr', 'password','phone', 'address', 'postcode', 'josephine', 'rankin', 1, 1, 1, true, true,  0, '2016-11-05','2016-11-05'),
('sr@sophiadex.com', 'sr', 'passwrord', 'phone', 'address', 'postcode','sydney', 'rankin', 1, 1, 1, true, true,  0, '2016-11-05','2016-11-05'),
('jar@sophiadex.com', 'sr', 'password','phone', 'address', 'postcode', 'sarah', 'rankin', 1, 1, 1, true, true,  0, '2016-11-05','2016-11-05');

INSERT INTO staff (email, username, password, phone, address, postcode, firstname, lastname, companies_id, stores_id, dept_id, siteowner, active, breaks_id, created_on, updated_on) VALUES 
('fb@sophiadex.com', 'xxxx', 'password', 'phone', 'address', 'postcode', 'fred', 'blogs', 1, 1, 1, true, true, 0, '2016-11-05','2016-11-05'),
('aib@sophiadex.com', 'xxxxx', 'password', 'phone', 'address', 'postcode', 'ada', 'blogs', 1, 1, 1, true, true,  0, '2016-11-05','2016-11-05'),
('jb@sophiadex.com', 'xxxxxx', 'password','phone', 'address', 'postcode', 'josephine', 'blogs', 1, 1, 1, true, true,  0, '2016-11-05','2016-11-05');


-- tasks

INSERT INTO tasks (task_name) VALUES ('stack shelves'),
('rotate stock'),
('person tills'),
('person self service'),
('open store'),
('rewcieve deliveries');

INSERT INTO tasks (task_name) VALUES ('stack shelves'),
('delete rotate stock'),
('delete person tills'),
('delete person self service'),
('delete open store'),
('delete rewcieve deliveries');

-- breaks
INSERT INTO breaks (break_name, break_length) VALUES ('ten', 10),('fifteen', 15),('thirty', 30),('fourty five', 45),('one hour', 60);

-- staff_task_link
INSERT INTO staff_task_link (task_id, staff_id) VALUES (1,1),(1,2),(1,3);

-- holidays
INSERT INTO holidays (holiday_name, holiday_start, holiday_end) VALUES ('one', '2016-12-01', '2016-12-14'),('two', '2016-12-03', '2016-12-17'),('three', '2017-12-01', '2017-12-05');
INSERT INTO holidays (holiday_name, holiday_start, holiday_end) VALUES ('delete one', '2016-12-01', '2016-12-14'),('delete two', '2016-12-03', '2016-12-17'),('delete three', '2017-12-01', '2017-12-05');


-- Table: public.shifts

-- DROP TABLE public.shifts;

CREATE TABLE public.shifts
(
    id integer NOT NULL DEFAULT nextval('shifts_id_seq'::regclass),
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    companies_id integer NOT NULL,
    stores_id integer NOT NULL,
    mon_start_time time without time zone DEFAULT '00:00:00'::time without time zone,
    mon_end_time time without time zone DEFAULT '00:00:00'::time without time zone,
    tue_start_time time without time zone DEFAULT '00:00:00'::time without time zone,
    tue_end_time time without time zone DEFAULT '00:00:00'::time without time zone,
    wed_start_time time without time zone DEFAULT '00:00:00'::time without time zone,
    wed_end_time time without time zone DEFAULT '00:00:00'::time without time zone,
    thu_start_time time without time zone DEFAULT '00:00:00'::time without time zone,
    thu_end_time time without time zone DEFAULT '00:00:00'::time without time zone,
    fri_start_time time without time zone DEFAULT '00:00:00'::time without time zone,
    fri_end_time time without time zone DEFAULT '00:00:00'::time without time zone,
    sat_start_time time without time zone DEFAULT '00:00:00'::time without time zone,
    sat_end_time time without time zone DEFAULT '00:00:00'::time without time zone,
    sun_start_time time without time zone DEFAULT '00:00:00'::time without time zone,
    sun_end_time time without time zone DEFAULT '00:00:00'::time without time zone,
    created_on timestamp without time zone DEFAULT now(),
    updated_on timestamp without time zone DEFAULT now(),
    CONSTRAINT shifts_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.shifts
    OWNER to rlnzxnlqoxmncu;
COMMENT ON TABLE public.shifts
    IS 'basic week times for the shift';

INSERT INTO shifts (sp_name,user_id,sp_start_date,sp_end_date,week_1,week_2,week_3,week_4,companies_id,stores_id,roll_over) VALUES
('Bernie',1,'2016-12-11','2016-12-17',1,2,',',1,1,True);


-- Table: public.shift_patterns

-- DROP TABLE public.shift_patterns;

CREATE TABLE public.shift_patterns
(
    id integer NOT NULL DEFAULT nextval('shift_patterns_id_seq'::regclass),
    sp_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    user_id integer NOT NULL,
    sp_start_date date,
    sp_end_date date,
    week_1 integer,
    week_2 integer,
    week_3 integer,
    week_4 integer,
    companies_id integer NOT NULL,
    stores_id integer NOT NULL,
    roll_over boolean DEFAULT true,
    created_on timestamp without time zone DEFAULT now(),
    updated_on timestamp without time zone DEFAULT now(),
    CONSTRAINT shift_patterns_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.shift_patterns
    OWNER to rlnzxnlqoxmncu;

COMMENT ON COLUMN public.shift_patterns.week_1
    IS 'a shift id w1 to w4';


INSERT INTO shift_patterns (name,companies_id,stores_id,mon_start_time,mon_end_time,tue_start_time,tue_end_time,wed_start_time,wed_end_time,thu_start_time,thu_end_time,fri_start_time,fri_end_time,sat_start_time,sat_end_time,sun_start_time,sun_end_time) VALUES 
('test',1,1,'09:00:00','17:00:00','09:00:00','17:00:00','09:00:00','17:00:00','09:00:00','17:00:00','09:00:00','17:00:00','09:00:00','17:00:00','09:00:00','17:00:00');








