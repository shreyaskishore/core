package migration

const create_tokens_table string = `
CREATE TABLE tokens (
	username VARCHAR(32) NOT NULL,
	token VARCHAR(256) NOT NULL,
	expiration BIGINT NOT NULL
);
`
