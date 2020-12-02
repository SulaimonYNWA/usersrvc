package db

const CreateUsersAccaunt = `create table if not exists users(
id integer Primary Key autoincrement,
name text not null,
surname text not null,
age integer not null,
gender text not null,
login text unique,
password text not null,
isAdmin boolean not null default false,
remove boolean not null default false
)`

const CreateCurrencyTable = `create table if not exists currency(
id integer Primary Key autoincrement,
name text not null
)`

const CreateATMsTable = `create table if not exists ATMs(
id integer Primary Key autoincrement,
address text not null,
works boolean not null default true
);`

