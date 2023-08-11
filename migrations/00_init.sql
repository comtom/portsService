CREATE TABLE IF NOT EXISTS ports (
	id SERIAL PRIMARY KEY,
	code_id TEXT NOT NULL UNIQUE,
	name TEXT,
	city TEXT,
	province TEXT,
	country TEXT,
	coordinates POINT,
	-- ideally should be geography type, this is just for avoiding installing PostGis
	timezone TEXT,
	unlocs TEXT [],
	code TEXT
);
CREATE INDEX idx_ports_code_id on ports using btree(code_id);