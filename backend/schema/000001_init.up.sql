CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255)          default '',
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    email         varchar(255) not null unique,
    pfp_url       varchar(255)          default '',
    bio           text         not null default ''
);

CREATE TABLE event_states
(
    id   serial      not null unique,
    name varchar(20) not null unique
);

INSERT INTO event_states (name)
VALUES ('scheduled'),
       ('in_progress'),
       ('ended');

CREATE TABLE events
(
    id            serial                                                          not null unique,
    title         varchar(255)                                                    not null,
    description   text                                                            not null default '',
    picture_url   varchar(255)                                                             default '',
    start_date    timestamp                                                       not null default NOW(),
    author_id     int references users (id) on delete restrict on update restrict not null,
    state_id      int references event_states (id)                                not null default 2,
    creation_date timestamp                                                       not null default NOW()
);

CREATE TABLE suggestions
(
    id               serial                                                          not null unique,
    author_id        int references users (id) on delete restrict on update restrict not null,
    post_date        timestamp                                                       not null default now(),
    sprint_id        int references sprints (id) on delete restrict on update restrict       not null,
    votes            int                                                             not null default 0,
    link_to_the_text varchar(255)                                                    not null
);

CREATE TABLE sprint_states
(
    id   serial      not null unique,
    name varchar(20) not null unique
);

INSERT INTO sprint_states (name)
VALUES ('voting'),
       ('implementing'),
       ('ended');

CREATE TABLE sprints
(
    id                   serial                                       not null unique,
    state_id             int references sprint_states (id)            not null default 1,
    suggestion_winner_id int                                                   default 0,
    event_id             int references events (id) on delete cascade not null,
    start                timestamp                                             default NOW(),
    winner_description   varchar(255)                                          default ''
);