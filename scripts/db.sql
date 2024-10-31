CREATE TABLE `user` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `username` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `role` bigint NOT NULL,
    `age` bigint NOT NULL,
    `degree` bigint NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_user_username` (`username`),
    UNIQUE KEY `uni_user_email` (`email`),
    KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `job` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `title` varchar(255) NOT NULL,
    `description` text NOT NULL,
    `location` varchar(255) NOT NULL,
    `company` varchar(255) NOT NULL,
    `salary` varchar(255) NOT NULL,
    `demand` text NOT NULL,
    `job_type` varchar(255) NOT NULL,
    `owner` bigint NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_job_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `user_job` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) DEFAULT NULL,
                            `updated_at` datetime(3) DEFAULT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            `user_id` bigint NOT NULL,
                            `job_id` bigint NOT NULL,
                            `progress` bigint NOT NULL,
                            PRIMARY KEY (`id`,`user_id`,`job_id`),
                            KEY `idx_user_job_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `resume` (
                          `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                          `created_at` datetime(3) DEFAULT NULL,
                          `updated_at` datetime(3) DEFAULT NULL,
                          `deleted_at` datetime(3) DEFAULT NULL,
                          `user_id` bigint NOT NULL,
                          `name` varchar(255) NOT NULL,
                          `gender` bigint NOT NULL,
                          `phone` varchar(255) DEFAULT NULL,
                          `email` varchar(255) DEFAULT NULL,
                          `wechat` varchar(255) DEFAULT NULL,
                          `state` bigint NOT NULL,
                          `description` text,
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `uni_resume_user_id` (`user_id`),
                          KEY `idx_resume_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `resume_education` (
                                    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                    `created_at` datetime(3) DEFAULT NULL,
                                    `updated_at` datetime(3) DEFAULT NULL,
                                    `deleted_at` datetime(3) DEFAULT NULL,
                                    `resume_id` bigint NOT NULL,
                                    `school` varchar(255) NOT NULL,
                                    `major` varchar(255) NOT NULL,
                                    `degree` bigint NOT NULL,
                                    `start_time` date NOT NULL,
                                    `end_time` date NOT NULL,
                                    PRIMARY KEY (`id`),
                                    KEY `idx_resume_education_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `resume_experience` (
                                     `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                     `created_at` datetime(3) DEFAULT NULL,
                                     `updated_at` datetime(3) DEFAULT NULL,
                                     `deleted_at` datetime(3) DEFAULT NULL,
                                     `resume_id` bigint NOT NULL,
                                     `company` varchar(255) NOT NULL,
                                     `position` varchar(255) NOT NULL,
                                     `start_time` date NOT NULL,
                                     `end_time` date NOT NULL,
                                     PRIMARY KEY (`id`),
                                     KEY `idx_resume_experience_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `resume_project` (
                                  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                  `created_at` datetime(3) DEFAULT NULL,
                                  `updated_at` datetime(3) DEFAULT NULL,
                                  `deleted_at` datetime(3) DEFAULT NULL,
                                  `resume_id` bigint NOT NULL,
                                  `name` varchar(255) NOT NULL,
                                  `description` text,
                                  `start_time` date NOT NULL,
                                  `end_time` date NOT NULL,
                                  PRIMARY KEY (`id`),
                                  KEY `idx_resume_project_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

