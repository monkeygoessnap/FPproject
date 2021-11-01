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
  `calories` int,
  `created` datetime,
  `updated` datetime,
  FOREIGN KEY (`merchant_id`) REFERENCES `user` (`id`)
);

CREATE TABLE IF NOT EXISTS `cart_item` (
  `item_id` varchar(36) UNIQUE,
  `user_id` varchar(36),
  `qty` int,
  `remarks` varchar(255),
  `created` datetime,
  `updated` datetime,
  FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `food` (`id`)
);

-- CREATE TABLE IF NOT EXISTS `sgfood` (
--   `ID` int PRIMARY KEY AUTO_INCREMENT,
--   `name` varchar(255),
--   `serving` varchar(20),
--   `calories` int
-- );