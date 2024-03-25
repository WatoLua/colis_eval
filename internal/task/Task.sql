/*
##Commands to create and switch to a new database##

drop database if exists eval;
create database eval;

\c eval
*/

drop table if exists task;

create table Task (
    id serial not null primary key,
    title char(50),
    description text,
    status smallint
);