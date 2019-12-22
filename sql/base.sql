CREATE TABLE `users`
(
    `id`                   bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `chat_id`              char(25)            NOT NULL,
    `gender`               tinyint(1) unsigned          DEFAULT NULL,
    `last_activity`        datetime            NOT NULL,
    `match_chat_id`        char(25)                     DEFAULT NULL,
    `gender_need_match`    tinyint(1)                   DEFAULT NULL,
    `available`            tinyint(1)          NOT NULL DEFAULT '0',
    `register_date`        datetime            NOT NULL,
    `previous_match`       char(25)                     DEFAULT NULL,
    `safe_mode`            tinyint(1)          NOT NULL DEFAULT '1',
    `banned_until`         datetime                     DEFAULT NULL,
    `waiting_timestamp`    bigint(20) unsigned          DEFAULT NULL,
    `last_gender_changing` datetime                     DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `chat_id` (`chat_id`),
    UNIQUE KEY `match_chat_id` (`match_chat_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

CREATE TABLE `reports`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`     bigint(20) unsigned NOT NULL,
    `reporter_id` bigint(20) unsigned NOT NULL,
    `created_at`  datetime            NOT NULL,
    `report`      text                NOT NULL,
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`),
    KEY `reporter_id` (`reporter_id`),
    CONSTRAINT `reports_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `reports_ibfk_2` FOREIGN KEY (`reporter_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

create index idx_waiting on users (waiting_timestamp);