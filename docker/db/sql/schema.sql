CREATE DATABASE IF NOT EXISTS `geonames` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;


CREATE TABLE IF NOT EXISTS `geonames`.`feature` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `code` VARCHAR(64),
    `name` VARCHAR(255),
    `description` TEXT,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`code`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

create if not exists table `location` (
    geoname_id varchar(32) not null,
    PRIMARY KEY (`geoname_id`)
);