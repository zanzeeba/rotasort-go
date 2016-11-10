        id           int64
        email        string
        username     string
        password     string
        phone        string
        address      string
        postcode     string
        firstname    string
        lastname     string
        companies_id int64
        stores_id    int64
        dept_id      int64
            siteowner    bool
        active       bool
            breaks_id    int64
            created_on   int64
            updated_on   int64

phone, address, postcode

breaks_id, created_on, updated_on

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



task_name VARCHAR(64) not null,
task_type TASKTYPE DEFAULT 'Volume Related',
weighting WEIGHTING DEFAULT 'fixed',
time_when TIME DEFAULT '00:00:00',
time_offset INTEGER DEFAULT 0,
time_float BOOLEAN  DEFAULT false,
time_length INTERVAL DEFAULT '0 1:01:01',
no_of_jobs INTEGER DEFAULT 1,
time_min INTERVAL DEFAULT '0 2:02:02',   
time_max INTERVAL DEFAULT '0 3:03:03',
day_of_week DAY_OF_WEEK DEFAULT 0,
active BOOLEAN DEFAULT true, 
companies_id INTEGER not null DEFAULT 0,
stores_id INTEGER not null DEFAULT 0,
dept_id INTEGER not null DEFAULT 0,

INSERT INTO tasks (task_name) VALUES ('stack shelves'),
('rotate stock'),
('person tills'),
('person self service'),
('open store'),
('rewcieve deliveries');


  id             int64
  task_name       string  
  task_type       string  
  weighting      string   
  time_when         time.Time
  time_offset     int64    
  time_float         time.Time 
  time_length      int64   
  no_of_jobs        int64  
  time_min       string  
  time_max        string 
  day_of_week     string    
  active         bool
  companies_id    int64     
  stores_id      int64   
  dept_id       int64  

id,task_name,task_type,weighting,time_when,time_offset,time_float,time_length,no_of_jobs,time_min,time_max,day_of_week,active,companies_id,stores_id,dept_id

&id ,&task_name ,&task_type ,&weighting ,&time_when ,&time_offset ,&time_float ,&time_length ,&no_of_jobs ,&time_min ,&time_max ,&day_of_week ,&active ,&companies_id ,&stores_id ,&dept_id

Id, TaskName, TaskType, Weightin, TimeWhen, Time_Offset TimeFloat, TimeLength, NoOfJobs, TimeMin, TimeMax, DayOfWeek, Active, CompaniesId StoresId, DeptId


tRes.Id          id
tRes.TaskName    task_name
tRes.TaskType    task_type
tRes.Weightin    weighting
tRes.TimeWhen    time_when
tRes.Time_Offset time_offset
tRes.TimeFloat   time_float
tRes.TimeLength  time_length
tRes.NoOfJobs    no_of_jobs
tRes.TimeMin     time_min
tRes.TimeMax     time_max
tRes.DayOfWeek   day_of_week
tRes.Active      active
tRes.CompaniesId companies_id
tRes.StoresId    stores_id
tRes.DeptId      dept_id




