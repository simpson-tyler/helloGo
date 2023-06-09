BEGIN;

CREATE TABLE IF NOT EXISTS apples
(
	id BIGINT PRIMARY KEY,
	color VARCHAR(50),
	variety VARCHAR(50),
	date_of_pollination TIMESTAMP NOT NULL,
	tree_id BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS trees
(
	id BIGINT PRIMARY KEY,
	date_of_planting TIMESTAMP NOT NULL,
	name VARCHAR(50) NOT NULL
);

ALTER TABLE apples ADD FOREIGN KEY (tree_id) REFERENCES trees (id);

COMMIT;
