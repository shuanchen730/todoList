    CREATE TABLE `cards` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `location` int NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


    CREATE TABLE `tasks` (
    `id` int NOT NULL AUTO_INCREMENT,
    `status` tinyint(1) NOT NULL,
    `content` varchar(255) NOT NULL,
    `card_id` int NOT NULL,
    `location` int NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;