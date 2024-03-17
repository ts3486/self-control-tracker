CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "paid_user" bool NOT NULL DEFAULT (false),
  "favorite_topic_1" varchar NOT NULL DEFAULT (''),
  "favorite_topic_2" varchar NOT NULL DEFAULT (''),
  "favorite_topic_3" varchar NOT NULL DEFAULT (''),
  "favorite_topic_4" varchar NOT NULL DEFAULT (''),
  "favorite_topic_5" varchar NOT NULL DEFAULT (''),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);