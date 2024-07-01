CREATE TABLE users(
    id serial primary key not null,
    username varchar(255) not null,
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp
);

CREATE TABLE photos(
    id serial primary key not null,
    title varchar(255) not null,
    caption text,
    photo_url text not null,
    user_id int not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    constraint fk_photo_user_id foreign key (user_id) references users(id)
);