CREATE TABLE IF NOT EXISTS `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) UNIQUE,
  `full_name` varchar(255),
  `password` varchar(100),
  `type` ENUM ('admin', 'merchants', 'users'),
  `created_at` datetime,
  `updated_at` timestamp
);

CREATE TABLE IF NOT EXISTS `address` (
  `id` int,
  `postal` varchar(6),
  `floor` varchar(3),
  `unit` varchar(9),
  FOREIGN KEY (`id`) REFERENCES `users` (`id`)
);

CREATE TABLE IF NOT EXISTS `merchant_details` (
  `id` int,
  `open_time` varchar(255),
  `close_time` varchar(255),
  FOREIGN KEY (`id`) REFERENCES `users` (`id`)
);

CREATE TABLE IF NOT EXISTS `health` (
  `id` int,
  `height` float,
  `weight` float,
  `age` int,
  `bmi` float,
  `active` ENUM (`low`,`moderate`,`high`),
  `target_weight` float,
  `target_bmi` float,
  `target_cal` float,
  `reset` varchar(4),
  FOREIGN KEY (`id`) REFERENCES `users` (`id`)
);

CREATE TABLE IF NOT EXISTS `items` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `merchant_id` int,
  `price` float,
  `status` ENUM ('soldout', 'available'),
  `calories` float,
  `created_at` datetime,
  `updated_at` timestamp,
  FOREIGN KEY (`merchant_id`) REFERENCES `users` (`id`)
);

CREATE TABLE IF NOT EXISTS `cart` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `created_at` datetime,
  `updated_at` timestamp,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE IF NOT EXISTS `cart_items` (
  `cart_id` int,
  `item_id` int,
  `quantity` int,
  `request` varchar(255),
  FOREIGN KEY (`cart_id`) REFERENCES `cart` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `items` (`id`)
);

CREATE TABLE IF NOT EXISTS `order` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `status` ENUM ('pending', 'ready', 'done'),
  `created_at` datetime,
  `updated_at` timestamp,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE IF NOT EXISTS `order_items` (
  `order_id` int,
  `item_id` int,
  `quantity` int,
  `request` varchar(255),
  FOREIGN KEY (`order_id`) REFERENCES `order` (`id`),
  FOREIGN KEY (`item_id`) REFERENCES `items` (`id`)
);
