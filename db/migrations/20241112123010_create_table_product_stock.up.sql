CREATE TABLE `product_stock` (
  `id` int NOT NULL AUTO_INCREMENT,
  `stock_amount` int NOT NULL,
  `expired_date` date DEFAULT NULL,
  `SKU` varchar(100) NOT NULL,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;