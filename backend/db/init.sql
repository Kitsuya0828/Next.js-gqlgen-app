CREATE TABLE `todos` (
	`id` VARCHAR(20) NOT NULL,
	`text` VARCHAR(255) NOT NULL,
	`done` BOOLEAN NOT NULL,
    `userId` VARCHAR(10)
	PRIMARY KEY (`id`)
) ENGINE InnoDB, CHARSET utf8mb4, COLLATE utf8mb4_0900_ai_ci;

INSERT INTO `todos` (id, text, done, userId) VALUES ('1', 'test', 'false', 'admin');