CREATE TABLE users (
    id_user serial not null unique,
    mail    varchar(255) not null,
    name    varchar(255) not null,
    username    varchar(255) not null unique,
    password_hash   varchar(255) not null,
    phone_number    varchar(255),
    rank    boolean not null
);

CREATE TABLE items (
    id_item         serial not null unique,
    title           varchar(128) not null,
    discription     varchar(255) not null,
    status          varchar(16) not null,
    count           int,
    price           int
);

CREATE TABLE users_items (
    id          serial not null unique,
    id_user     int references users (id_user) on delete cascade not null,
    id_item     int references items (id_item) on delete cascade not null
);

CREATE TABLE favorite (
    id_user     int references users (id_user) on delete cascade not null,
    id_item     int references items (id_item) on delete cascade not null
);

CREATE TABLE posts (
    id_post         serial not null unique,
    discription     varchar(255) not null,
    creation_time   timestamptz not null,
    id_item         int references items (id_item),
    price           int
);

CREATE TABLE users_posts (
    id_user     int references users (id_user) on delete cascade not null,
    id_post     int references posts (id_post) on delete cascade not null
);

CREATE TABLE notifications (
    id          serial not null unique,
    content     varchar(255) not null,
    creation_time   timestamptz not null
);

CREATE TABLE users_notifications (
    id_user             int references users (id_user) on delete cascade not null,
    id_notification     int references notifications (id) on delete cascade not null
);

CREATE TABLE notifications_status (
    id_notification     int references notifications (id) on delete cascade not null,
    status              boolean      not null default false
);