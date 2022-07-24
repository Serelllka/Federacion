ALTER TABLE users
    ADD COLUMN email varchar(255);

CREATE TABLE articles
(
    id serial not null unique,
    name varchar(255),
    current_cost int not null,
    description text
);

CREATE TABLE condemnations
(
    id serial not null unique,
    condemner_id int references users(id) on delete cascade not null,
    convicted_id int references users(id) on delete cascade not null,
    article_id int references articles(id) on delete cascade not null,
    description text,
    cost int not null,
    time time not null
);
