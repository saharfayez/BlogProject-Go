create table "posts"
(
    "id"      SERIAL PRIMARY KEY,
    "title"   varchar(50),
    "content" varchar(50),
    "user_id" INTEGER
);
