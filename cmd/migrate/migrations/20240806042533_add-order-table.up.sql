CREATE TABLE IF NOT EXISTS orders (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `orderId` VARCHAR(255) NOT NULL,
    `productId` INT UNSIGNED NOT NULL,
    `quantity` INT(255) NOT NULL,
    `totalPrice` INT(255) NOT NULL,
    `meja` ENUM('1', '2', '3') NOT NULL,
    `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    FOREIGN KEY (`productId`) REFERENCES products(`id`)
);

