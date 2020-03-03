package migrations

const create_users_table = `
CREATE TABLE users (
	username VARCHAR(32) NOT NULL,
	first_name VARCHAR(64) NOT NULL,
	last_name VARCHAR(64) NOT NULL,
	graduation_year INTEGER NOT NULL,
	major VARCHAR(32) NOT NULL,
	mark VARCHAR(32) NOT NULL
);
`
