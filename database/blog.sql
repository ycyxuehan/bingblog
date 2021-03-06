create database if not exists blog default character set UTF8MB4;
use blog;

drop table if exists category;
create table if not exists category (
	id int auto_increment primary key,
    name varchar(64) not null,
    description varchar(256) default ''
);

drop table if exists user;
create table if not exists user (
	id int auto_increment primary key,
    name varchar(64) not null,
    password varchar(64) default '123456'
);

drop table if exists article;
create table if exists article(
    id int auto_increment primary key,
    title varchar(128) not null,
    category int default '-1',
    outline varchar(256) default '',
    create_date bigint default 0,
    reads int default 0,
    deleted tinyint default 0
);

drop table if exists article_content;
create table if not exists article_content(
    id int auto_increment primary key,
    article int not null,
    content text default ''
);
