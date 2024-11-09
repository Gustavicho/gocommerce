CREATE TABLE IF NOT EXISTS orders_items (
  `id` SERIAL,
  `order_id` BIGINT UNSIGNED NOT NULL,
  `product_id` BIGINT UNSIGNED NOT NULL,
  `quantity` INT NOT NULL CHECK (quantity > 0),
  `price` DECIMAL(10, 2) NOT NULL CHECK (price > 0),
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (`id`),
  FOREIGN KEY (`order_id`) REFERENCES orders(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`product_id`) REFERENCES products(`id`) ON DELETE CASCADE
);