-- +goose Up
create table Users(
    id  int primary key generated by default as identity,
    name varchar(20) not null ,
    surname varchar(20) not null ,
    patronymic varchar(20),
    age int not null ,
    gender varchar(10) not null ,
    country varchar(30) not null
);

-- +goose Down
TRUNCATE TABLE users;
