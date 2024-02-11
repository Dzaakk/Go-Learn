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

DELETE
FROM users
WHERE id BETWEEN '10' AND '20';

create table user_logs
(
    id         serial,
    user_id    varchar(100) not null,
    action     varchar(100) not null,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
);
select *
from user_logs;

delete
from user_logs
where action = 'Test';

-- Remove default value
ALTER TABLE user_logs
    ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE user_logs
    ALTER COLUMN updated_at TYPE bigint
        USING EXTRACT(EPOCH FROM updated_at)::bigint;

select *
from information_schema.columns
where table_name = 'user_logs';

create table todos
(
    id          serial,
    user_id     varchar(100) not null,
    title       varchar(100) not null,
    description text         null,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at  timestamp    null,
    primary key (id)
);
drop table todos;
select *
from todos;

create table wallets
(
    id         VARCHAR(100) NOT NULL,
    user_id    VARCHAR(100) NOT NULL,
    balance    bigint       NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    foreign key (user_id) references users (id)
);
select *
from wallets;
create table addresses
(
    id         serial,
    user_id    VARCHAR(100) NOT NULL,
    address    VARCHAR(100) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    foreign key (user_id) references users (id)
);
drop table addresses;
select *
from addresses;

create table products
(
    id         varchar(100) not null,
    name       varchar(100) not null,
    price      bigint       not null,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
truncate products;
drop table products;

create table user_like_product
(
    user_id    varchar(100) not null,
    product_id varchar(100) not null,
    PRIMARY KEY (user_id, product_id),
    foreign key (user_id) references users (id),
    foreign key (product_id) references products (id)
)
select * from user_like_product;
truncate user_like_product;