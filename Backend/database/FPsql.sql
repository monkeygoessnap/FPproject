CREATE DATABASE IF NOT EXISTS `FP`;

USE `FP`;

CREATE TABLE IF NOT EXISTS `user` (
  `id` varchar(36) PRIMARY KEY,
  `username` varchar(255) UNIQUE,
  `name` varchar(255),
  `password` varchar(255),
  `type` ENUM ('admin', 'merchant', 'customer'),
  `created` datetime,
  `updated` datetime
);

CREATE TABLE IF NOT EXISTS `address` (
  `id` varchar(36) PRIMARY KEY,
  `postal` varchar(6),
  `floor` varchar(3),
  `unit` varchar(9),
  `created` datetime,
  `updated` datetime,
  FOREIGN KEY (`id`) REFERENCES `user` (`id`)
);

CREATE TABLE IF NOT EXISTS `userhealth` (
  `id` varchar(36) PRIMARY KEY,
  `gender` varchar(10), 
  `height` float,
  `weight` float,
  `dob` varchar(6),
  `active` varchar(10), 
  `target` varchar(10),
  `created` datetime,
  `updated` datetime,
  FOREIGN KEY (`id`) REFERENCES `user` (`id`)
);

CREATE TABLE IF NOT EXISTS `food` (
  `id` varchar(36) PRIMARY KEY,
  `merchant_id` varchar(36),
  `name` varchar(255),
  `price` float,
  `status` varchar(10),
  `description` varchar(255),
  `imglink` varchar(255),
  `created` datetime,
  `updated` datetime,
  FOREIGN KEY (`merchant_id`) REFERENCES `user` (`id`)
);

CREATE TABLE IF NOT EXISTS `cart_item` (
  `item_id` varchar(36) PRIMARY KEY,
  `user_id` varchar(36),
  `qty` int,
  `remarks` varchar(255),
  `created` datetime,
  `updated` datetime,
  FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `food` (`id`)
);

CREATE TABLE IF NOT EXISTS `sgfood` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `serving` varchar(20),
  `calories` int
);

-- CREATE TABLE IF NOT EXISTS `order` (
--   `id` varchar(36) PRIMARY KEY,
--   `user_id` varchar(36),
--   `status` ENUM ('pending', 'completed'),
--   `created` datetime,
--   `updated` datetime,
--   FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
-- );

-- CREATE TABLE IF NOT EXISTS `order_item` (
--   `food_id` varchar(36) PRIMARY KEY,
--   `merchant_id` varchar(36),
--   `name` varchar(255),
--   `price` float,
--   `description` varchar(255),
--   `imglink` varchar(255),
--   `order_id` varchar(36),
--   `qty` int,
--   `remarks` varchar(255),
--   `created` datetime,
--   `updated` datetime,
--   FOREIGN KEY (`merchant_id`) REFERENCES `user` (`id`),
--   FOREIGN KEY (`food_id`) REFERENCES `food` (`id`),
--   FOREIGN KEY (`order_id`) REFERENCES `order` (`id`)
-- );


-- LOAD CSV FILE INTO DB

-- CLI COMMAND TO LOAD FILE PREVENTED BY --secure-file-priv
-- option so it cannot execute this statement

-- CREATE TEMPORARY TABLE `temp_table` LIKE `sgfood`;
-- SHOW INDEX FROM `temp_table`;
-- DROP INDEX `PRIMARY` ON `temp_table`;

-- LOAD DATA INFILE 'sgfood.csv'
-- INTO TABLE `temp_table`
-- FIELDS TERMINATED BY ','
-- OPTIONALLY ENCLOSED BY '"'
-- LINES TERMINATED BY '/n'
-- IGNORE 1 ROWS;

-- SHOW COLUMNS FROM `sgfood`;
-- INSERT INTO `sgfood`
-- SELECT * FROM `temp_table`
-- ON DUPLICATE KEY UPDATE `name`=VALUES(`name`), `serving`=VALUES(`serving`), `calories`=VALUES(`calories`);

-- DROP TEMPORARY TABLE `temp_table`;