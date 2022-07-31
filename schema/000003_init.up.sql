CREATE TABLE usersInfo
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    image varchar(255),
    hobbies text,
    lore text,
    sobriquet varchar(255)
);

CREATE TABLE userIntegrations (
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    discord_id bigint
);