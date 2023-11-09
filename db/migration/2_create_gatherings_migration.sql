-- gathering_app.gatherings definition

CREATE TABLE `gatherings` (
  `id` bigint NOT NULL,
  `creator` bigint NOT NULL,
  `type` int NOT NULL,
  `scheduled_at` datetime DEFAULT NULL,
  `name` varchar(100) NOT NULL,
  `location` varchar(100) NOT NULL,
  `created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
	`deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`)
);