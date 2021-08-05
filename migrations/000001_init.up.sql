CREATE TABLE `links`
(
    `id`         INT(11) NOT NULL AUTO_INCREMENT,
    `link`       LONGTEXT     NOT NULL,
    `crop_link`  VARCHAR(255) NOT NULL,
    `created_at` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;