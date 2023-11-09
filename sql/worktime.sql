-- Drops all tables
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS deleted_messages;
DROP TABLE IF EXISTS edited_messages;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS channels;
DROP TABLE IF EXISTS guilds;
--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

-- Creates all tables

CREATE TABLE channels (
guild_id    BIGINT NOT NULL,
channel_id  BIGINT NOT NULL,
channel_id  VARCHAR (2000),
PRIMARY KEY (channel_id),
FOREIGN KEY (guild_id) REFERENCES guilds(guild_id)
);

CREATE TABLE messages_Deleted (
message_id BIGINT NOT NULL,
channel_id BIGINT NOT NULL,
author_id  BIGINT NOT NULL,
guild_id   BIGINT NOT NULL,
message_content VARCHAR(2000) NOT NULL,
date_deleted DATE NOT NULL,
time_deleted TIME NOT NULL,
FOREIGN KEY (chanel_id) REFERENCES  channels     (channel_id),
FOREIGN KEY (guild_id) REFERENCES   guilds      (guild_id),
FOREIGN KEY (author_id) REFERENCES  users        (author_id)
);

CREATE TABLE messages_Edited (
message_id  BIGINT NOT NULL,
channel_id  BIGINT NOT NULL,
author_id   BIGINT NOT NULL,
guild_id    BIGINT NOT NULL,
before_edited_content   VARCHAR(2000) NOT NULL,
after_edited_content    VARCHAR(2000) NOT NULL,
date_edited DATE NOT NULL,
time_edited TIME NOT NULL,
PRIMARY KEY (messag_id),
FOREIGN KEY (channel_id) REFERENCES channels    (chanel_id),
FOREIGN KEY (guild_id) REFERENCES   guilds      (guild_id),
FOREIGN KEY (author_id) REFERENCES  users       (author_id)
);

CREATE TABLE guilds (
guild_id BIGINT NOT NULL,
guild_name VARCHAR(2000),
PRIMARY KEY (guild_id)
);

CREATE TABLE messages (
message_id  BIGINT NOT NULL,
guild_id    BIGINT NOT NULL,
channel_id  BIGINT NOT NULL,
author_id   BIGINT NOT NULL,
message_content VARCHAR(2000),
date_sent   DATE NOT NULL,
time_sent   TIME NOT NULL,
PRIMARY KEY (message_id),
FOREIGN KEY (channels)  REFERENCES channels(channel_id),
FOREIGN KEY (guilds)    REFERENCES guilds(guild_id),
FOREIGN KEY (users)     REFERENCES users(author_id)
);

CREATE TABLE users (
author_id   BIGINT NOT NULL,
author_tag  BIGINT NOT NULL,
PRIMARY KEY (author_id)
);