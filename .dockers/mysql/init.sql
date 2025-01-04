
CREATE DATABASE IF NOT EXISTS `go_course_users`;

CREATE TABLE IF NOT EXISTS `go_course_users`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(40),
  `last_name` VARCHAR(40),
  `email` VARCHAR(40) NOT NULL,
  PRIMARY KEY (`id`)
);