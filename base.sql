CREATE TABLE users (
    id PRIMARY key NOT NULL,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    name TEXT,
);

CREATE TABLE exercises
(
  id         SERIAL            NOT NULL
    CONSTRAINT exercises_pkey
    PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE,
  rus        TEXT              NOT NULL,
  eng        TEXT              NOT NULL,
  rank       NUMERIC DEFAULT 0 NOT NULL
);

CREATE INDEX idx_exercises_deleted_at
  ON exercises (deleted_at);


CREATE TABLE user2exercise
(
    user_id INTEGER NOT NULL,
    exercise_id INTEGER NOT NULL,
    rank numeric DEFAULT 0 NOT NULL,
    passed INTEGER DEFAULT 0 NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY(user_id, exercise_id),
)