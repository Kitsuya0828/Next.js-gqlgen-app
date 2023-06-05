CREATE TABLE todos (
	id VARCHAR(20) NOT NULL PRIMARY KEY,
	content VARCHAR(255) NOT NULL,
	done BOOLEAN NOT NULL,
    user_id VARCHAR(10)
) ENGINE InnoDB, CHARSET utf8mb4, COLLATE utf8mb4_0900_ai_ci;

INSERT INTO todos (id, content, done, user_id) VALUES ('1', 'test', false, 'admin');

CREATE TABLE users (
	id VARCHAR(20) NOT NULL PRIMARY KEY,
	name VARCHAR(50) NOT NULL
) ENGINE InnoDB, CHARSET utf8mb4, COLLATE utf8mb4_0900_ai_ci;

INSERT INTO users (id, name) VALUES ('1', 'test 1');