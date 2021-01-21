CREATE TABLE `tests` (
  `id` bigint(20) unsigned NOT NULL,
  `test_name` varchar(127) NOT NULL DEFAULT '',
  `test_description` varchar(255),
  `status` int(10) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
);
