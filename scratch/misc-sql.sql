DROP TYPE IF EXISTS WEIGHTING CASCADE;
DROP TYPE IF EXISTS DAY_OF_WEEK CASCADE;
DROP TYPE IF EXISTS TEMPLATE CASCADE;
DROP TYPE IF EXISTS HOURS CASCADE;
DROP TYPE IF EXISTS MINUTES CASCADE;
DROP TYPE IF EXISTS TASKTYPE CASCADE;


CREATE TYPE WEIGHTING AS ENUM ('fixed', 'critical', 'slippable', 'moveable');
CREATE TYPE DAY_OF_WEEK AS ENUM ('0', '1', '2', '3', '4', '5', '6');
CREATE TYPE TEMPLATE AS ENUM ('0', '1', '2', '3', '4', '5');
CREATE TYPE HOURS AS ENUM('00', '01', '02', '03', '04', '05', '06', '07', '08', '09', '10', '11', '12', '13', '14', '15', '16', '17', '18', '19', '20', '21', '22', '23');
CREATE TYPE MINUTES AS ENUM('00', '01', '02', '03', '04', '05', '06', '07', '08', '09', '10', '11', '12', '13', '14', '15', '16', '17', '18', '19', '20', '21', '22', '23', '24', '25', '26', '27', '28', '29', '30', '31', '32', '33', '34', '35', '36', '37', '38', '39', '64', '41', '42', '43', '44', '45', '46', '47', '48', '49', '50', '51', '52', '53', '54', '55', '56', '57', '58', '59');

CREATE TYPE TASKTYPE AS ENUM('Volume Related', 'Customer Related', 'Customer Related Multiple', 'Fixed Length Filler', 'Variable Length Filler');

-- for rotasort to create tables
DROP TABLE IF EXISTS public.tasks;
DROP TABLE IF EXISTS public.staff;
DROP TABLE IF EXISTS public.skills_staff_link;
DROP TABLE IF EXISTS public.skills_task_link;
DROP TABLE IF EXISTS public.calendars;
DROP TABLE IF EXISTS public.break_values;
DROP TABLE IF EXISTS public.companies;
DROP TABLE IF EXISTS public.stores;
DROP TABLE IF EXISTS public.departments;
DROP TABLE IF EXISTS public.operations;
DROP TABLE IF EXISTS public.skills;
DROP TABLE IF EXISTS public.templates;
DROP TABLE IF EXISTS public.requirements;
DROP TABLE IF EXISTS public.registrations;

DROP TABLE IF EXISTS public.tasks;
CREATE TABLE IF NOT EXISTS tasks (
id serial primary key,
task_name VARCHAR(64) not null,
task_type TASKTYPE DEFAULT 'Volume Related',
weighting WEIGHTING DEFAULT 'fixed',
time_when TIME DEFAULT '00:00:00',
time_offset INTEGER DEFAULT 0,
time_float BOOLEAN  DEFAULT false,
time_length INTEGER DEFAULT 0,
no_of_jobs INTEGER DEFAULT 1,
time_min INTEGER DEFAULT 0,   
time_max INTEGER DEFAULT 0,
day_of_week DAY_OF_WEEK DEFAULT '0',
active BOOLEAN DEFAULT true, 
companies_id INTEGER not null DEFAULT 0,
stores_id INTEGER not null DEFAULT 0,
dept_id INTEGER not null DEFAULT 0,
updated_on timestamp without time zone DEFAULT now(),
created_on timestamp without time zone DEFAULT now()
);                    
COMMENT ON TABLE public.tasks IS 'tasks to be performed';

CREATE TABLE IF NOT EXISTS staff(
id serial primary key,
email VARCHAR(255) DEFAULT,
username VARCHAR(60) DEFAULT,
password VARCHAR(64) DEFAULT,
phone VARCHAR(64) DEFAULT,
address TEXT DEFAULT,
postcode VARCHAR(64) DEFAULT,
firstname VARCHAR(64) DEFAULT,
lastname VARCHAR(64) DEFAULT,
companies_id INTEGER not null,
stores_id INTEGER not null,
dept_id INTEGER not null,
siteowner BOOLEAN DEFAULT false,
active BOOLEAN DEFAULT true,
breaks_id INTEGER DEFAULT,
created_on timestamp without time zone DEFAULT now(),
updated_on timestamp without time zone DEFAULT now()
);
COMMENT ON TABLE public.staff IS 'staff in the store';

-- wrong should be staff to skills 
-- and a tasks to skills or something better

CREATE TABLE IF NOT EXISTS skills_staff_link(
id serial primary key,
staff_id INTEGER not null,
skills_id INTEGER not null,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.skills_staff_link IS 'the linkage of staff with skills';

CREATE TABLE IF NOT EXISTS skills_task_link(
id serial primary key,
task_id INTEGER not null,
skills_id INTEGER not null,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.skills_task_link IS 'the linkage of tasks with skills';

CREATE TABLE IF NOT EXISTS calendars(
id serial primary key,
companies_id INTEGER not null,
stores_id INTEGER not null,
day DAY_OF_WEEK,
date DATE,
opo TIME,
opc TIME,
puo TIME,
puc TIME,
open BOOLEAN,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.calendars IS 'ability to build a calendar to use in the system';

CREATE TABLE IF NOT EXISTS break_values(
id serial primary key,
break_start TIME,
break_end TIME,
companies_id INTEGER not null,
stores_id INTEGER not null,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.break_values IS 'the times allowed for breaks';

CREATE TABLE IF NOT EXISTS companies(
id serial primary key,
company_name VARCHAR(64) not null,
company_email VARCHAR(255) not null,
address VARCHAR(255) not null,
postcode VARCHAR(255) not null,
phone_number VARCHAR(255) not null,
company_url VARCHAR(255) not null,
company_notes TEXT,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.companies IS 'company information';

CREATE TABLE IF NOT EXISTS stores(
id serial primary key,
companies_id INTEGER not null,
store_name VARCHAR(255) not null,
email VARCHAR(255) not null,
address VARCHAR(255) not null,
postcode VARCHAR(255) not null,
phone_number VARCHAR(255) not null,
stores_url VARCHAR(255) not null,
stores_notes TEXT,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.stores IS 'store information';

CREATE TABLE IF NOT EXISTS departments(
id serial primary key,
department_name VARCHAR(255) not null,
companies_id INTEGER not null,
stores_id INTEGER not null,
department_notes TEXT,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.departments IS 'department information';

CREATE TABLE IF NOT EXISTS operations(
id serial primary key,
companies_id INTEGER not null,
stores_id INTEGER not null,
day_of_week DAY_OF_WEEK,
open_closed BOOLEAN,
opening_time TIME,
closing_time TIME,
operating_hours_start TIME,
operating_hours_end TIME,
shift_lengths INTERVAL,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.operations IS 'operationg times, opening times etc';

CREATE TABLE IF NOT EXISTS skills(
id serial primary key,
skill VARCHAR(64) not null,
description TEXT,
companies_id INTEGER not null,
stores_id INTEGER not null,
skill_weighting WEIGHTING,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.skills IS 'skills require to do a task';


CREATE TABLE IF NOT EXISTS templates(
id serial primary key,
name VARCHAR(255) not null,
companies_id INTEGER not null,
stores_id INTEGER not null,
template_type TEMPLATE,
description TEXT,
day DAY_OF_WEEK,
inherit_template_id INTEGER,
date DATE,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.templates IS 'i thingk this was for the more complex shifts model';


CREATE TABLE IF NOT EXISTS requirements(
id serial primary key,
name VARCHAR(255) not null,
companies_id INTEGER not null,
stores_id INTEGER not null,
start_hour HOURS,
start_minute MINUTES,
end_hour HOURS,
end_minute MINUTES,
skill_id INTEGER,
template_id INTEGER,
staff_id INTEGER,
percentage_cover INTEGER,
start_offset INTERVAL,
end_offset INTERVAL,
date DATE,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.requirements IS 'not sure what i created this for';


CREATE TABLE IF NOT EXISTS registrations(
id serial primary key,
email VARCHAR(255) not null,
companies_id INTEGER not null,
stores_id INTEGER not null,
name VARCHAR(255) not null,
comments TEXT,
classification  VARCHAR(255) not null,
subject VARCHAR(255) not null,
message VARCHAR(255) not null,
created_on DATE,
updated_on DATE
);
COMMENT ON TABLE public.registrations IS 'not sure what i created this for';



-- ....................................
DROP TABLE public.rainbow;

CREATE TABLE public.rainbow
(
"abstime" abstime,
"smallint" smallint,
"integer" integer,
"bigint" bigint,
"decimal" decimal,
"numeric" numeric,
"real" real,
"double_precision" double precision,
"smallserial" smallserial,
"serial" serial,
"bigserial" bigserial,

"money" money,   

"character_varying" character varying(4), 
"varchar" varchar(4),

"character" character(6), 
"char" char(6),
"text" text,    

"bytea" bytea,   

"timestamp_wtz" timestamp,
"timestamp_no_tz" timestamp,
"date" date,
"time_notz"  time,
"time_wtz" time,
"interval" interval,

"boolean" boolean, 

"current_mood" mood,

"point" point,  
"line" line,   
"lseg" lseg,   
"box" box,    
"path_closed" path,   
"path_open" path,   
"polygon" polygon,
"circle" circle, 


"cidr"    cidr,    
"inet"    inet,    
"macaddr" macaddr 




)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.rainbow
    OWNER to jmhowitt;
COMMENT ON TABLE public.rainbow
    IS 'so i can see how all the data types are written';


COMMENT ON COLUMN public.rainbow.smallint IS '2 bytes small-range integer -32768 to +32767';
COMMENT ON COLUMN public.rainbow.integer IS '4 bytes typical choice for integer  -2147483648 to +2147483647';
COMMENT ON COLUMN public.rainbow.bigint  IS '8 bytes large-range integer -9223372036854775808 to 9223372036854775807';
COMMENT ON COLUMN public.rainbow.decimal IS 'variable    user-specified precision, exact up to 131072 digits before the decimal point; up to 16383 digits after the decimal point';
COMMENT ON COLUMN public.rainbow.numeric IS 'variable    user-specified precision, exact up to 131072 digits before the decimal point; up to 16383 digits after the decimal point';
COMMENT ON COLUMN public.rainbow.real IS '4 bytes variable-precision, inexact 6 decimal digits precision';
COMMENT ON COLUMN public.rainbow.double_precision IS '8 bytes variable-precision, inexact 15 decimal digits precision';
COMMENT ON COLUMN public.rainbow.smallserial IS '2 bytes small autoincrementing integer  1 to 32767';
COMMENT ON COLUMN public.rainbow.serial  IS '4 bytes autoincrementing integer    1 to 2147483647';
COMMENT ON COLUMN public.rainbow.bigserial IS '8 bytes   large autoincrementing integer  1 to 9223372036854775807';

COMMENT ON COLUMN public.rainbow.money IS '8 bytes currency amount -92233720368547758.08 to +92233720368547758.07';

COMMENT ON COLUMN public.rainbow.character_varying IS ' variable-length with limit';
COMMENT ON COLUMN public.rainbow.varchar IS ' variable-length with limit';

  
COMMENT ON COLUMN public.rainbow.character IS 'fixed-length, blank padded';
COMMENT ON COLUMN public.rainbow.char IS 'fixed-length, blank padded';

COMMENT ON COLUMN public.rainbow.text IS 'variable unlimited length';


COMMENT ON COLUMN public.rainbow.bytea IS   '1 or 4 bytes plus the actual binary string  variable-length binary string';

COMMENT ON COLUMN public.rainbow.timestamp_wtz IS '[ (p) ] [ without time zone ] 8 bytes both date and time (no time zone
)   4713 BC 294276 AD   1 microsecond / 14 digits';
COMMENT ON COLUMN public.rainbow.timestamp_no_tz IS '[ (p) ] with time zone    8 bytes both date and time, with time zone  4713 BC 294276 AD   1 microsecond / 14 digits';
COMMENT ON COLUMN public.rainbow.date IS '4 bytes date (no time of day)   4713 BC 5874897 AD  1 day';
COMMENT ON COLUMN public.rainbow.time_notz IS '[ (p) ] [ without time zone ]  8 bytes time of day (no date
)   00:00:00    24:00:00    1 microsecond / 14 digits';
COMMENT ON COLUMN public.rainbow.time_wtz IS '[ (p) ] with time zone 12 bytes    times of day only, with time zone   00:00:00+1459   24:00:00-1459   1 microsecond / 14 digits';
COMMENT ON COLUMN public.rainbow.interval IS '[ fields ] [ (p) ] 16 bytes    time interval   -178000000 years    178000000 years 1 microsecond / 14 digits';

COMMENT ON COLUMN public.rainbow.boolean IS '1 byte  state of true or false';

COMMENT ON COLUMN public.rainbow.current_mood IS 'enumerated Enum types are created using the CREATE TYPE command, for example';

COMMENT ON COLUMN public.rainbow.point IS '16 bytes    Point on a plane    (x,y)';
COMMENT ON COLUMN public.rainbow.line IS '32 bytes    Infinite line (not fully implemented)   ((x1,y1),(x2,y2))';
COMMENT ON COLUMN public.rainbow.lseg IS '32 bytes    Finite line segment ((x1,y1),(x2,y2))';
COMMENT ON COLUMN public.rainbow.box IS '32 bytes    Rectangular box ((x1,y1),(x2,y2))';
COMMENT ON COLUMN public.rainbow.path_closed IS '16+16n bytes    Closed path (similar to polygon)    ((x1,y1),...)';
COMMENT ON COLUMN public.rainbow.path_open IS '16+16n bytes    Open path   [(x1,y1),...]';
COMMENT ON COLUMN public.rainbow.polygon IS '64+16n bytes    Polygon (similar to closed path)    ((x1,y1),...)';
COMMENT ON COLUMN public.rainbow.circle IS '24 bytes    Circle  <(x,y),r> (center point and radius)';

COMMENT ON COLUMN public.rainbow.cidr  IS '7 or 19 bytes   IPv4 and IPv6 networks';
COMMENT ON COLUMN public.rainbow.inet IS '7 or 19 bytes   IPv4 and IPv6 hosts and networks';
COMMENT ON COLUMN public.rainbow.macaddr IS '6 bytes MAC addresses';






