create table sample
(
    id   varchar(100) not null,
    name varchar(100) not null,
    primary key (id)
);

select *
from sample;

create table users
(
    id         VARCHAR(100) NOT NULL,
    password   VARCHAR(100) NOT NULL,
    name       VARCHAR(100) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
select *
from users;

alter table users
    rename column name to first_name;
alter table users
    add column middle_name varchar(200) null;

alter table users
    add column last_name varchar(100) null;
