-- -------------------------------------------------------------
-- TablePlus 5.1.0(468)
--
-- https://tableplus.com/
--
-- Database: bds
-- Generation Time: 2022-11-01 11:43:52.7300
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

DROP TABLE IF EXISTS `real_estate_price_histories`;
CREATE TABLE `real_estate_price_histories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `re_id` int NOT NULL,
  `price` float NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `real_estate_amenities`;
CREATE TABLE `real_estate_amenities` (
  `re_id` int NOT NULL AUTO_INCREMENT,
  `amenity_id` int NOT NULL,
  PRIMARY KEY (`re_id`,`amenity_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `real_estate_images`;
CREATE TABLE `real_estate_images` (
                                      `re_id` int NOT NULL AUTO_INCREMENT,
                                      `image_id` int NOT NULL,
                                      PRIMARY KEY (`re_id`,`image_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `duration` bigint NOT NULL,
  `auto_renew` tinyint(1) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DROP TABLE IF EXISTS `real_estate_types`;
CREATE TABLE `real_estate_types` (
                                     `type_id` int NOT NULL,
                                     `re_id` int NOT NULL,
                                     `is_rent` tinyint(1) NOT NULL,
                                     PRIMARY KEY (`re_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



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
  `built_at` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `documents` text NOT NULL,
  `status` int DEFAULT '1',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `real_estate_post_types` (`id`, `name`, `price_per_day`) VALUES
(1, 'premium', 20000),
(2, 'normal', 10000);


DROP TABLE IF EXISTS `images`;
CREATE TABLE `images` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `url` text NOT NULL,
                          `width` int DEFAULT NULL,
                          `height` int DEFAULT NULL,
                          `cloud_name` varchar(50) DEFAULT NULL,
                          `extension` varchar(50) DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;