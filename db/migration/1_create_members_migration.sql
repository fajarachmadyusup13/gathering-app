-- gathering_app.members definition

CREATE TABLE gathering_app.members (
	id BIGINT NOT NULL,
	first_name varchar(100) NOT NULL,
	last_name varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	deleted_at DATETIME NULL,
	PRIMARY KEY (id)
);