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

CREATE TABLE IF NOT EXISTS "star"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "name"       VARCHAR  NOT NULL,
    "alias_name" INTEGER,
    "cover"      TEXT,
    "created_at" DATETIME NOT NULL default CURRENT_TIMESTAMP,
    "updated_at" DATETIME,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "star_index_0"
    ON "star" ("id");
CREATE INDEX IF NOT EXISTS "star_index_1"
    ON "star" ("name");

CREATE TABLE IF NOT EXISTS "movie_star"
(
    "id"         INTEGER  NOT NULL UNIQUE,
    "movie_id"   INTEGER  NOT NULL,
    "star_id"    INTEGER,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "movie_star_index_0"
    ON "movie_star" ("movie_id");
CREATE INDEX IF NOT EXISTS "movie_star_index_1"
    ON "movie_star" ("star_id");


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
