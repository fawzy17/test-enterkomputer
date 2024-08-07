CREATE TABLE IF NOT EXISTS products (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `category` ENUM('Makanan', 'Minuman', 'Promo') NOT NULL,
    `variant` VARCHAR(255),
    `price` INT(255) NOT NULL,
    `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id)
);