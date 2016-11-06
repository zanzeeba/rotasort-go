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

