create database NIXdb;

CREATE TABLE posts
(
    UserId int                   not null unique,
    id     serial auto_increment not null unique,
    title  varchar(255)          not null,
    body   varchar(255)          not null
);


CREATE TABLE comments
(
    id    serial       not null unique,
    name  varchar(255) not null,
    email varchar(255) not null,
    body  varchar(255) not null
);

CREATE TABLE posts_comments
(
    id          serial                                         not null unique,
    post_id     int references posts (id) on DELETE cascade   ,
    comments_id int references comments (id) on DELETE cascade
);


ALTER TABLE comments ADD COLUMN post_id VARCHAR(20);