CREATE DATABASE IF NOT EXISTS casbin;

USE casbin;

CREATE TABLE IF NOT EXISTS casbin_rule (
	p_type VARCHAR(100) NOT NULL,
	v0 VARCHAR(100),
	v1 VARCHAR(100),
	v2 VARCHAR(100),
	v3 VARCHAR(100),
	v4 VARCHAR(100),
	v5 VARCHAR(100)
);

INSERT INTO casbin_rule VALUES
('p', 'dajun', 'data1', 'read', '', '', ''),
('p', 'lizi', 'data2', 'write', '', '', '');