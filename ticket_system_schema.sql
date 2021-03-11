create database ticket_system;
-- drop database ticket_system;
use  ticket_system;

-- drop table users;
create table users(
	ID int AUTO_INCREMENT not null primary key,
    name nvarchar(50),
    email varchar(200),
    password varchar(200),
    salt varchar(20),
    birthday date,
    phone varchar(11),
    work_place nvarchar(200),
    role int
);

-- drop table projects;
-- create table projects(
-- 	ID int AUTO_INCREMENT not null primary key,
--     name nvarchar(50),
--     description text,
--     date_created datetime,
--     date_updated datetime
-- );

-- create table assets_in_use(
-- ID int AUTO_INCREMENT not null primary key,
--     user_id int,
--     asset_id int,
--     valid_from datetime,
--     valid_to datetime
-- );

-- drop table tickets;
create table assets(
	ID int AUTO_INCREMENT not null primary key,
    name nvarchar(50),
    serial_number varchar(20),
    type_id int,
    user_id int,
    description text
);
-- drop table ticket_type;
create table asset_type(
	ID int AUTO_INCREMENT not null primary key,
    name nvarchar(50),
    description text
);

alter table users 
	add constraint fk_users_project_id foreign key(project_id) references projects(ID);
alter table assets 
	add constraint fk_assets_type_id foreign key(type_id) references assets(ID),
	add constraint fk_assets_user_id foreign key(user_id) references users(ID);

