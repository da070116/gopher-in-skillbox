CREATE TABLE IF NOT EXISTS users
(
    "id"   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" VARCHAR(200),
    "age"  INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS friends
(
    "id"        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "owner_id"  INTEGER NOT NULL,
    "friend_id" INTEGER NOT NULL,
    CONSTRAINT "fk_owner"
        foreign key ("owner_id")
            references users (id),
    CONSTRAINT "fk_friend"
        foreign key ("friend_id")
            references users (id)
);