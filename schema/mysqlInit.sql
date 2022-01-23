drop table posts;
drop table comments;

CREATE TABLE posts
(
    UserId int          not null,
    id     int          not null unique,
    title  varchar(255) not null,
    body   varchar(255) not null
);


CREATE TABLE comments
(
    id      int       not null unique,
    post_id int references posts (id) on delete cascade,
    name    varchar(255) not null,
    email   varchar(255) not null,
    body    varchar(255) not null
);




