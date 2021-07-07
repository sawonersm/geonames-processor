CREATE DATABASE IF NOT EXISTS `geonames` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;


CREATE TABLE IF NOT EXISTS `geonames`.`feature` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `class` VARCHAR(8),
    `code` VARCHAR(64),
    `name` VARCHAR(255),
    `description` TEXT,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`code`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `geonames`.`location` (
    `geoname_id` VARCHAR(32) NOT NULL,
    `name` VARCHAR(32) NOT NULL,
    `latitude` DECIMAL(10,7) NOT NULL,
    `longitude` DECIMAL(11,7) NOT NULL,
    `feature` INT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME,
    PRIMARY KEY (`geoname_id`),
    FOREIGN KEY (`feature`) REFERENCES `geonames`.`feature`(`id`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;