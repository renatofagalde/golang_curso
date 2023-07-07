### Docker compose
run docker for mysql: ``docker-compose -f docker-compose-mysql-only.yml up -d``

### tabela
criacao da tabela usuários:
```
create table usuarios( 
    id int auto_increment primary key,
    nome varchar(100) not null,
    email varchar(100) not null
) engine =innodb;
```


````go mod init banco-de-dados````

### Download driver mysql
```` go get github.com/go-sql-driver/mysql````