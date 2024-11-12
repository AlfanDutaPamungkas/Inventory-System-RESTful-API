CREATE TABLE `products` (
  `SKU` varchar(100) NOT NULL,
  `product_name` varchar(200) NOT NULL,
  `product_brand` varchar(100) NOT NULL,
  `category` varchar(100) NOT NULL,
  `price` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `image_url` text,
  PRIMARY KEY (`SKU`)
) ENGINE=InnoDB;