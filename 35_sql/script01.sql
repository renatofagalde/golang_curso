create table usuarios(
                         id int auto_increment primary key,
                         nome varchar(100) not null,
                         email varchar(100) not null
) engine =innodb;

create user 'aplicacao'@'localhost' identified by 'app@@123';
grant select, delete, update, execute, insert, drop, alter, create on dev.* to 'aplicacao'@'localhost';
