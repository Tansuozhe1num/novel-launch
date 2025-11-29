CREATE TABLE `books` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `book_name` VARCHAR(255) NOT NULL,
  `author` VARCHAR(100) NOT NULL,
  `desc` TEXT,
  `tag` VARCHAR(100),
  `cover` VARCHAR(500),
  `create_time` BIGINT NOT NULL DEFAULT 0,
  `update_time` BIGINT NOT NULL DEFAULT 0,
  `status` BIGINT NOT NULL DEFAULT 0,
  `is_deleted` BIGINT NOT NULL DEFAULT 0,
  `heat` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_tag` (`tag`),
  KEY `idx_status` (`status`),
  KEY `idx_heat` (`heat`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
