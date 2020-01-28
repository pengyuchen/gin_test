-- MySql 

create database if not exists gin_exer;
use gin_exer;
create table person (
	id int(11) not null auto_increment primary key,
	firstname varchar(40) not null default "",
	lastname  varchar(40) not null default ""
) engine=innodb default charset=utf8;


-- CREATE TABLE `users` (
--   `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
--   `name` varchar(64) DEFAULT NULL,
--   `telephone` varchar(12) DEFAULT '',
--   PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;
