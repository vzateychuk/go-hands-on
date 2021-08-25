DROP TABLE IF EXISTS items;

CREATE TABLE items (
                       id SERIAL,
                       title VARCHAR(255) NOT NULL,
                       description text NOT NULL,
                       updated varchar(255) DEFAULT NULL,
                       PRIMARY KEY (id)
);

INSERT INTO items (id, title, description, updated) VALUES
(1,	'database/sql',	'Рассказать про базы данных',	'vzateychuk'),
(2,	'memcache',	'Рассказать про мемкеш с примером использования',	NULL);