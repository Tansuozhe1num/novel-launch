CREATE TABLE `book_chapter` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `book_id` BIGINT NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `is_vip` TINYINT(1) NOT NULL DEFAULT 0,
  `create_time` BIGINT NOT NULL DEFAULT 0,
  `update_time` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_chapter_book_id` (`book_id`),
  KEY `idx_chapter_is_vip` (`is_vip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
