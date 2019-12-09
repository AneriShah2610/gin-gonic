-- migrate:up
CREATE TABLE todo (
	id INT8 NOT NULL DEFAULT unique_rowid(),
	title STRING NOT NULL,
	todo_status STRING NOT NULL DEFAULT 'created':::STRING,
	created_by STRING NOT NULL,
	created_at TIMESTAMP NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	FAMILY "primary" (id, title, todo_status, created_by, created_at)
);

-- migrate:down
DROP TABLE IF EXISTS todo;