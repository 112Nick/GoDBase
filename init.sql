CREATE EXTENSION IF NOT EXISTS citext;


-- DROP TABLE IF EXISTS forum;
-- DROP TABLE IF EXISTS users;
-- DROP TABLE IF EXISTS post;
-- DROP TABLE IF EXISTS vote;
-- DROP TABLE IF EXISTS thread;
-- DROP TABLE IF EXISTS forum_users;


CREATE TABLE IF NOT EXISTS forum (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  "user" TEXT NOT NULL,
  slug CITEXT UNIQUE,
  posts INTEGER,
  threads INTEGER
);

CREATE TABLE IF NOT EXISTS thread (
  id SERIAL PRIMARY KEY,
  votes INTEGER,
  slug CITEXT UNIQUE ,
  forum TEXT NOT NULL ,
  forumID INTEGER,
  author TEXT,
  title TEXT,
  message TEXT,
  created TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  fullname text,
  nickname  CITEXT COLLATE ucs_basic NOT NULL UNIQUE,
  email CITEXT NOT NULL UNIQUE,
  about text
);

CREATE TABLE IF NOT EXISTS post (
  id SERIAL PRIMARY KEY ,
  parent INTEGER DEFAULT 0,
  author TEXT,
  message TEXT,
  isedited BOOLEAN,
  forum TEXT,
  created TIMESTAMP WITH TIME ZONE ,

  thread INTEGER ,
  path INTEGER[]
);

CREATE TABLE IF NOT EXISTS vote (
  id SERIAL PRIMARY KEY,
  userId INTEGER,
  nickname TEXT,
  threadID INTEGER,
  voice int,
  forum CITEXT
);

CREATE TABLE IF NOT EXISTS forum_users (
  id SERIAL PRIMARY KEY,
  --userId INTEGER,
  fullname text,
  nickname  CITEXT COLLATE ucs_basic NOT NULL,
  email CITEXT,
  about text,
  forumID INTEGER,
  UNIQUE (forumID, nickname)
);