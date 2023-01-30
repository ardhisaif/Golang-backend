CREATE TABLE vehicles (
    vehicle_id serial NOT NULL,
    type varchar(355) DEFAULT '',
    name varchar(355) DEFAULT '',
    location varchar(100) DEFAULT '',
    capacity int DEFAULT 1,
    price int default 0,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp,
    PRIMARY KEY (vehicle_id)
); 