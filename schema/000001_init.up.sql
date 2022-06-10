CREATE TABLE users (
    id_user         serial not null unique,
    mail            varchar(255) not null unique,
    name            varchar(255) not null, 
    username        varchar(255) not null unique,
    password_hash   varchar(255) not null,
    url_avatar      varchar(255),
    phone_number    varchar(255),
    rank            boolean not null
);

CREATE TABLE items (
    id_item         serial not null unique,
    title           varchar(128) not null,
    url_media       varchar(255),
    description     varchar(255) not null,
    status          varchar(16) not null,
    count           int,
    price           int
);

CREATE TABLE followers (
    id_user                 int references users (id_user) on delete cascade not null,
    id_user_follower        int references users (id_user) on delete cascade not null
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
    description     varchar(255) not null,
    url_media       varchar(255),
    creation_time   timestamptz not null,
    id_item         int references items (id_item),
    price           int
);

CREATE TABLE users_posts (
    id_user     int references users (id_user) on delete cascade not null,
    id_post     int references posts (id_post) on delete cascade not null
);

CREATE TABLE notifications (
    id_notice              serial not null unique,
    content                varchar(255) not null,
    status                 varchar(16)      not null
);

CREATE TABLE users_notifications (
    id_user             int references users (id_user) on delete cascade not null,
    id_notice           int references notifications (id_notice) on delete cascade not null
);

CREATE TABLE orders (
    id_order            serial not null unique,
    id_user_owner       int references users (id_user) on delete cascade not null,
    id_item             int references items (id_item) on delete cascade not null,
    status              varchar(16) not null,
    description         varchar(255) not null
);

CREATE TABLE users_orders (
    id_row              serial not null unique,
    id_user             int references users (id_user) on delete cascade not null,
    id_order            int references orders (id_order) on delete cascade not null
);

CREATE TABLE chats (
    id_chat             serial not null unique,
    status              varchar(16) not null,
    type                varchar(16) not null,
    id_order            int references orders (id_order) on delete cascade not null
);

CREATE TABLE users_chats (
    id_user             int references users (id_user) on delete cascade not null,
    id_chat             int references chats (id_chat) on delete cascade not null
);


CREATE TABLE messages (
    id_message              serial not null unique,
    id_chat                 int references chats (id_chat) on delete cascade not null,
    id_user                 int references users (id_user) on delete cascade not null,
    content                 varchar(255) not null,
    creation_time           timestamptz not null
);

ALTER TABLE favorite  
ADD CONSTRAINT flag_favorite UNIQUE (id_user, id_item);

