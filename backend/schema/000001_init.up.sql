CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255),
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    email         varchar(255) not null unique,
    pfp_url       varchar(255)          default null,
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
    id            serial                                                                 not null unique,
    title         varchar(255)                                                           not null,
    description   text                                                                   not null default '',
    picture_url   varchar(255),
    start         timestamp                                                              not null default NOW(),
    author_id     int references users (id) on delete restrict on update restrict        NOT NULL,
    state_id      int references event_states (id) on delete restrict on update restrict not null,
    creation_date timestamp                                                              not null default NOW()
);

CREATE TABLE suggestions
(
    id          serial                                                          not null unique,
    author_id   int references users (id) on delete restrict on update restrict not null,
    post_date   timestamp                                                       not null default now(),
    sprint_id   int                                                             not null,
    votes_count int                                                             not null default 0
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
    id                   serial                                                                  not null unique,
    state_id             int references sprint_states (id) on delete restrict on update restrict not null,
    suggestion_winner_id int references suggestions (id) on delete restrict on update restrict unique,
    event_id             int references events (id) on delete cascade on update cascade          not null,
    start                timestamp default NOW()
);