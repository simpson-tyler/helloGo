CREATE TABLE apples
(
	id BIGINT PRIMARY KEY,
	color VARCHAR(50),
	variety VARCHAR(50),
	date_of_pollination TIMESTAMP NOT NULL,
	tree_id BIGINT NOT NULL
);

CREATE TABLE trees
(
	id BIGINT PRIMARY KEY,
	date_of_planting TIMESTAMP NOT NULL
	name VARCHAR(50) NOT NULL
);

ALTER TABLE trees ADD FOREIGN KEY (tree_id) REFERENCES trees (id);
