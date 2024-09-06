
-- +migrate Up

CREATE TABLE books(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid TEXT,
	isbn TEXT,
	title TEXT,
	pages TEXT,
	current_page TEXT,
	author TEXT,
	year TEXT,
	status TEXT
);

-- +migrate Down

DROP TABLE books;
