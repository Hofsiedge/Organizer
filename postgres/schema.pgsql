CREATE SCHEMA IF NOT EXISTS common;

CREATE TABLE IF NOT EXISTS common.user
(
    id BIGSERIAL
        PRIMARY KEY
);

CREATE SCHEMA IF NOT EXISTS tasktracker;

-- @formatter:off
CREATE TYPE tasktracker.task_status AS ENUM (
    'created',
    'in process',
    'failed',
    'done',
    'planned',
    'archived'
);
-- @formatter:on

CREATE TABLE IF NOT EXISTS tasktracker.task
(
    id          BIGSERIAL
        PRIMARY KEY,
    name        TEXT,
    description TEXT,
    status      tasktracker.task_status DEFAULT 'created',
    owner_id    INTEGER
        REFERENCES common.user (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tasktracker.resource
(
    id           BIGSERIAL
        PRIMARY KEY,
    name         TEXT,
    description  TEXT,
    preview_link TEXT,
    link         TEXT
);

CREATE TABLE IF NOT EXISTS tasktracker.task_dependency
(
    id         BIGSERIAL
        PRIMARY KEY,
    task_id    BIGINT
        REFERENCES tasktracker.task (id) ON DELETE CASCADE,
    depends_on BIGINT
        REFERENCES tasktracker.task (id) ON DELETE CASCADE
);

CREATE SCHEMA IF NOT EXISTS environment;

CREATE TABLE IF NOT EXISTS environment.user_event
(
    id          BIGSERIAL
        PRIMARY KEY,
    name        TEXT,
    description TEXT,
    start_time  TIMESTAMP,
    end_time    TIMESTAMP,
    owner_id    INTEGER
        REFERENCES common.user (id) ON DELETE CASCADE
);
