CREATE TABLE usersInfo
(
    id serial not null unique primary key,
    user_id int references users(id) on delete cascade not null,
    image varchar(255) not null,
    hobbies text not null,
    lore text not null,
    sobriquet varchar(255) not null
);

CREATE TABLE userIntegrations (
    id serial not null unique primary key,
    user_id int references users(id) on delete cascade not null,
    discord_id bigint not null
);

CREATE TABLE achievements (
    id serial not null unique primary key,
    name varchar(255),
    description text,
    image varchar(255)
);

CREATE TABLE usersAchievements (
    id serial not null unique primary key,
    user_id int references users(id) on delete cascade not null,
    achievements int references achievements(id) on delete cascade not null,
    time timestamp not null default current_timestamp
);

