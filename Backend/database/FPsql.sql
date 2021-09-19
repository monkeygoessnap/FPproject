CREATE DATABASE IF NOT EXISTS `FP`;

USE `FP`;

CREATE TABLE IF NOT EXISTS `user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) UNIQUE,
  `name` varchar(255),
  `password` varchar(100),
  `type` ENUM ('admin', 'merchant', 'customer'),
  `created` datetime,
  `updated` timestamp
);

CREATE TABLE IF NOT EXISTS `address` (
  `id` int,
  `postal` varchar(6),
  `floor` varchar(3),
  `unit` varchar(9),
  `created` datetime,
  `updated` timestamp,
  FOREIGN KEY (`id`) REFERENCES `user` (`id`)
);

CREATE TABLE IF NOT EXISTS `userhealth` (
  `id` int,
  `gender` ENUM ('male', 'female'),
  `height` float,
  `weight` float,
  `dob` varchar(6),
  `active` ENUM ('low', 'moderate', 'high'),
  `target` ENUM ('gain', 'lose', 'maintain'),
  `created` datetime,
  `updated` timestamp,
  FOREIGN KEY (`id`) REFERENCES `user` (`id`)
);

CREATE TABLE IF NOT EXISTS `food` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `merchant_id` int,
  `name` varchar(255),
  `price` float,
  `status` ENUM ('avail', 'soldout'),
  `description` varchar(255),
  `imglink` varchar(255),
  `created` datetime,
  `updated` timestamp,
  FOREIGN KEY (`merchant_id`) REFERENCES `user` (`id`)
);

CREATE TABLE IF NOT EXISTS `cart_items` (
  `item_id` int,
  `user_id` int,
  `qty` int,
  `remarks` varchar(255),
  `created` datetime,
  `updated` timestamp,
  FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `food` (`id`)
);

CREATE TABLE IF NOT EXISTS `order` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `status` ENUM ('pending', 'completed'),
  `created` datetime,
  `updated` timestamp,
  FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
);

CREATE TABLE IF NOT EXISTS `order_items` (
  `food_id` int,
  `merchant_id` int,
  `name` varchar(255),
  `price` float,
  `description` varchar(255),
  `imglink` varchar(255),
  `order_id` int,
  `qty` int,
  `remarks` varchar(255),
  `created` datetime,
  `updated` timestamp,
  FOREIGN KEY (`merchant_id`) REFERENCES `user` (`id`),
  FOREIGN KEY (`food_id`) REFERENCES `food` (`id`),
  FOREIGN KEY (`order_id`) REFERENCES `order` (`id`)
);

CREATE TABLE IF NOT EXISTS `sgfood` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `serving` varchar(20),
  `calories` int
);

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