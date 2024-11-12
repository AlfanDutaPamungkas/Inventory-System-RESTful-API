CREATE TABLE `log_activity` (
  `id` int NOT NULL AUTO_INCREMENT,
  `message` text,
  `admin_id` varchar(200) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;