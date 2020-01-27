-- MySql 

-- create database if not exists gin_exer;
-- use gin_exer;
-- create table `person` (
-- 	`id` int(11) not null auto_increment primary key,
-- 	`firstname` varchar(40) not null default "",
-- 	`lastname`  varchar(40) not null default ""
-- ) engine=innodb default charset=utf8;


CREATE TABLE `person` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(40) NOT NULL DEFAULT '',
  `last_name` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

