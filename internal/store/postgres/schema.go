package postgres

var migrate = []string{
	`
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		email TEXT UNIQUE,
		"password" TEXT,
		api_key UUID,
		reset_password_token UUID,
		reset_password_expires TIMESTAMP,
		verification_expires TIMESTAMP,
		verification_token UUID,
		verified BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	);

	CREATE UNIQUE INDEX ON users(email);
`,

	`
	CREATE TABLE IF NOT EXISTS urls(
		id UUID PRIMARY KEY,
		owner UUID,
		original_url TEXT UNIQUE,
		short_url TEXT,
		visit_count INTEGER,
		created_at TIMESTAMP,
		updated_at TIMESTAMP,
		FOREIGN KEY (owner)  REFERENCES users(id) ON DELETE CASCADE
	);

	CREATE UNIQUE INDEX ON urls(original_url);
	CREATE UNIQUE INDEX ON urls(short_url);
	`,
}

var drop = []string{
	`DROP TABLE IF EXISTS users CASCADE`,
	`DROP TABLE IF EXISTS urls CASCADE`,
}
