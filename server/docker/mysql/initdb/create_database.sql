CREATE USER IF NOT EXISTS tweet@`%` IDENTIFIED BY 'passwd'; -- noqa: RF05

CREATE DATABASE IF NOT EXISTS `tweet` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci; -- noqa: LT05

GRANT ALL ON `tweet`.* TO tweet@'%'; -- noqa: RF05
