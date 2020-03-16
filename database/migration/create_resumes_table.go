package migration

const create_resumes_table string = `
CREATE TABLE resumes (
	username VARCHAR(32) NOT NULL,
	first_name VARCHAR(64) NOT NULL,
	last_name VARCHAR(64) NOT NULL,
	email VARCHAR(256) NOT NULL,
	graduation_month INT NOT NULL,
	graduation_year INT NOT NULL,
	major VARCHAR(64) NOT NULL,
	degree VARCHAR(64) NOT NULL,
	seeking VARCHAR(64) NOT NULL,
	blob_key VARCHAR(512) NOT NULL,
	approved BOOLEAN NOT NULL,
	UNIQUE KEY unique_username (username)
);
`
