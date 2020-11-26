package db

const CreateUsersAccaunt = `create table  if not exists users(
id integer Primary Key autoincrement,
name text not null,
surname text not null,
age integer not null,
gender text not null,
login text unique,
password text not null,
remove boolean not null default false
)`

