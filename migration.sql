create table IF NOT EXISTS users(
									id int NOT NULL AUTO_INCREMENT,
									name varchar(255),
	occupation varchar(255),
	email varchar(255),
	password varchar(255),
	avatar varchar(255),
	role varchar(255),
	created_at datetime,
	updated_at datetime,
	primary key (id)
	);


create table IF NOT EXISTS campaign(
	id int NOT NULL AUTO_INCREMENT,
	user_id int,
	name varchar(255),
	short_description varchar(255),
	description text,
	perk text,
	backer_count int,
	goal_amount int,
	current_amount int,
	slug varchar(255),
	created_at datetime,
	updated_at datetime,
	primary key (id),
	foreign key(user_id) references users(id) on update cascade on delete set null
);

create table if not EXISTS transactions(
    id int NOT NULL AUTO_INCREMENT,
    campaign_id int,
    user_id int,
    amount int,
    status varchar(255),
    code varchar(255),
	payment_url varchar(255),
	created_at datetime,
	updated_at datetime,
    primary key (id),
	foreign key(user_id) references users(id) on update cascade on delete set null,
	foreign key(campaign_id) references campaign(id) on update cascade on delete set null
);

create table if not EXISTS campaign_image(
id int NOT NULL AUTO_INCREMENT,
campaign_id int,
file_name varchar(255),
is_primary boolean,
	created_at datetime,
	updated_at datetime,
primary key (id),
foreign key(campaign_id) references campaign(id) on update cascade on delete set null
);