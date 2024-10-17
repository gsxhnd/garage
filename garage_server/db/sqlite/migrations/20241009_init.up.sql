CREATE TABLE IF NOT EXISTS "movie"
(
    "id"              INTEGER  NOT NULL UNIQUE,
    "code"            TEXT,
    "title"           TEXT,
    "cover"           TEXT,
    "publish_date"    DATETIME,
    "director"        TEXT,
    "produce_company" TEXT,
    "publish_company" TEXT,
    "series"          TEXT,
    "created_at"      DATETIME NOT NULL default CURRENT_TIMESTAMP,
    "updated_at"      DATETIME,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "movie_index_0"
    ON "movie" ("id");
CREATE INDEX IF NOT EXISTS "movie_index_1"
    ON "movie" ("code");

CREATE TABLE IF NOT EXISTS "actor"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "name"       TEXT  NOT NULL,
    "alias_name" TEXT,
    "cover"      TEXT,
    "created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP,
    "updated_at" DATETIME,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "actor_index_0"
    ON "actor" ("id");
CREATE INDEX IF NOT EXISTS "actor_index_1"
    ON "actor" ("name");

CREATE TABLE IF NOT EXISTS "movie_actor"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "movie_id"   INTEGER  NOT NULL,
    "actor_id"    INTEGER,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "movie_actor_index_0"
    ON "movie_actor" ("movie_id");
CREATE INDEX IF NOT EXISTS "movie_actor_index_1"
    ON "movie_actor" ("actor_id");


CREATE TABLE IF NOT EXISTS "tag"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "name"       TEXT     NOT NULL,
    "pid"        INTEGER,
    "created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP,
    "updated_at" DATETIME,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "tag_index_0"
    ON "tag" ("id");



CREATE TABLE IF NOT EXISTS "movie_tag"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "movie_id"   INTEGER  NOT NULL,
    "tag_id"     INTEGER,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "movie_tag_index_0"
    ON "tag" ("id");


CREATE TABLE IF NOT EXISTS "anime"
(
    "id"           INTEGER  NOT NULL UNIQUE,
    "title"        TEXT,
    "cover"        TEXT,
    "publish_date" DATETIME,
    "created_at"   DATETIME NOT NULL default CURRENT_TIMESTAMP,
    "updated_at"   DATETIME,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "anime_index_0"
    ON "anime" ("id");
