CREATE TABLE IF NOT EXISTS products (
  `id` SERIAL,
  `name` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `image` VARCHAR(255) NOT NULL,
  `price` DECIMAL(10, 2) NOT NULL CHECK (price > 0),
  `quantity` INT NOT NULL CHECK (quantity > 0),
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (`id`)
);