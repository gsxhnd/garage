CREATE TABLE IF NOT EXISTS "movie"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "code"       TEXT,
    "title"       TEXT,
    "updated_at" DATETIME,
    "created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "movie_index_0"
    ON "movie" ("id");


CREATE TABLE IF NOT EXISTS "star"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "name"       VARCHAR  NOT NULL,
    "alias_name"        INTEGER,
    "created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP,
    "updated_at" DATETIME,
    PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS "star_index_0"
    ON "star" ("id");



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

