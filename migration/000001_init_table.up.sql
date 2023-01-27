CREATE TABLE `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '',
    `occupation` varchar(100) DEFAULT '',
    `email` varchar(100) DEFAULT '',
    `password_hash` varchar(100) DEFAULT '',
    `avatar_file_name` varchar(100) DEFAULT '',
    `role` varchar(100) DEFAULT '',
    `token` varchar(100) DEFAULT '',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = UTF8 ROW_FORMAT = DYNAMIC;

CREATE TABLE `campaigns` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) unsigned DEFAULT NULL,
    `name` varchar(100) DEFAULT '',
    `short_description` varchar(250) DEFAULT '',
    `description` text DEFAULT '',
    `goal_amount` int DEFAULT '0',
    `current_amount` int DEFAULT '0',
    `perks` text DEFAULT '',
    `backer_count` int DEFAULT '0',
    `slug` varchar(100) DEFAULT '',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `fk_campaign_1_idx` (`user_id`) USING BTREE,
    CONSTRAINT `fk_campaign_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE = InnoDB DEFAULT CHARSET = UTF8 ROW_FORMAT = DYNAMIC;

CREATE TABLE `campaign_images` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `campaign_id` bigint(20) unsigned DEFAULT NULL,
    `file_name` varchar(100) DEFAULT '',
    `is_primary` tinyint(1) NOT NULL DEFAULT '0',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `fk_campaign_image_1_idx` (`campaign_id`) USING BTREE,
    CONSTRAINT `fk_campaign_image_1` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE = InnoDB DEFAULT CHARSET = UTF8 ROW_FORMAT = DYNAMIC;

CREATE TABLE `transactions` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `campaign_id` bigint(20) unsigned DEFAULT NULL,
    `user_id` bigint(20) unsigned DEFAULT NULL,
    `amount` int DEFAULT '0',
    `status` varchar(100) DEFAULT '',
    `code` varchar(100) DEFAULT '',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `fk_transaction_1_idx` (`campaign_id`) USING BTREE,
    KEY `fk_transaction_2_idx` (`user_id`) USING BTREE,
    CONSTRAINT `fk_transaction_1` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
    CONSTRAINT `fk_transaction_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE = InnoDB DEFAULT CHARSET = UTF8 ROW_FORMAT = DYNAMIC;