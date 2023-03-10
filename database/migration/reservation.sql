CREATE TABLE reservations (
	reservation_id serial not null,
	user_id int not null ,
    vehicle_id int NOT null,
    quantity int default 1,
    start_date timestamp,
    return_date timestamp,
    total_payment int,
    is_booked boolean default false,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp,
    PRIMARY KEY (reservation_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) on delete cascade on update cascade,
    FOREIGN KEY (vehicle_id) REFERENCES vehicles(vehicle_id) on delete cascade on update cascade
); 