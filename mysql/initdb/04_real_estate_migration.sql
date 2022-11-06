-- -------------------------------------------------------------
-- TablePlus 5.1.0(468)
--
-- https://tableplus.com/
--
-- Database: bds
-- Generation Time: 2022-11-06 16:55:12.7840
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


USE bds;

DROP TABLE IF EXISTS `amenities`;
CREATE TABLE `amenities` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `images`;
CREATE TABLE `images` (
  `id` int NOT NULL AUTO_INCREMENT,
  `url` text NOT NULL,
  `width` int DEFAULT NULL,
  `height` int DEFAULT NULL,
  `cloud_name` varchar(50) DEFAULT NULL,
  `extension` varchar(50) DEFAULT NULL,
  `owner_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `real_estate_amenities`;
CREATE TABLE `real_estate_amenities` (
  `re_id` int NOT NULL AUTO_INCREMENT,
  `amenity_id` int NOT NULL,
  PRIMARY KEY (`re_id`,`amenity_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `real_estate_images`;
CREATE TABLE `real_estate_images` (
  `re_id` int NOT NULL AUTO_INCREMENT,
  `image_id` int NOT NULL,
  PRIMARY KEY (`re_id`,`image_id`),
  KEY `image_id` (`image_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `real_estate_post_types`;
CREATE TABLE `real_estate_post_types` (
  `id` int NOT NULL,
  `name` varchar(200) NOT NULL,
  `price_per_day` float NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `real_estate_posts`;
CREATE TABLE `real_estate_posts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `re_id` int NOT NULL,
  `post_type_id` int NOT NULL,
  `start_date` timestamp NOT NULL,
  `duration` bigint NOT NULL,
  `auto_renew` tinyint(1) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` int DEFAULT '1',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `real_estate_price_histories`;
CREATE TABLE `real_estate_price_histories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `re_id` int NOT NULL,
  `price` float NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `real_estate_types`;
CREATE TABLE `real_estate_types` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `is_rent` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `real_estates`;
CREATE TABLE `real_estates` (
  `id` int NOT NULL AUTO_INCREMENT,
  `province_id` int NOT NULL,
  `district_id` int NOT NULL,
  `ward_id` int NOT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `latitude` float NOT NULL,
  `longitude` float NOT NULL,
  `price` float NOT NULL,
  `owner_id` int NOT NULL,
  `floors` int NOT NULL,
  `area` float NOT NULL,
  `no_bedrooms` int NOT NULL,
  `no_wc` int NOT NULL,
  `house_facing` enum('east','north','west','south','south_east','south_west','north_east','north_west') DEFAULT NULL,
  `balcony_facing` enum('east','north','west','south','south_east','south_west','north_east','north_west') DEFAULT NULL,
  `reason` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `built_at` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `interiors` text,
  `documents` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `re_type_id` int NOT NULL,
  `status` int DEFAULT '1',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `phone` varchar(15) NOT NULL,
  `fb_id` varchar(50) DEFAULT NULL,
  `gg_id` varchar(50) DEFAULT NULL,
  `password` varchar(100) NOT NULL,
  `salt` varchar(100) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `role` enum('buyer','seller','agency','admin') NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `avatar_url` varchar(255) DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO `amenities` (`id`, `name`) VALUES
(1, 'pool'),
(2, 'spa'),
(3, 'bar'),
(4, 'view'),
(5, 'elevator'),
(6, 'garage'),
(7, 'air_conditioning');

INSERT INTO `images` (`id`, `url`, `width`, `height`, `cloud_name`, `extension`, `owner_id`) VALUES
(4, 'd2asms0h14kd1r.cloudfront.net/img/521018297.jpg', 398, 398, NULL, NULL, 0),
(5, 'd2asms0h14kd1r.cloudfront.net/img/120323715.jpg', 398, 398, NULL, NULL, 0);

INSERT INTO `real_estate_post_types` (`id`, `name`, `price_per_day`) VALUES
(1, 'premium', 20000),
(2, 'normal', 10000);

INSERT INTO `real_estate_types` (`id`, `name`, `is_rent`) VALUES
(1, 'apartment', 0),
(2, 'house', 0),
(3, 'villa', 0),
(4, 'townhouse', 0),
(5, 'ground', 0),
(6, 'land', 0),
(7, 'resort', 0),
(8, 'condotel', 0),
(9, 'warehouse', 0),
(10, 'other', 0),
(11, 'apartment', 1),
(12, 'house', 1),
(13, 'villa', 1),
(14, 'townhouse', 1),
(15, 'ground', 1),
(16, 'land', 1),
(17, 'resort', 1),
(18, 'condotel', 1),
(19, 'warehouse', 1);



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;