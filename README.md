# mysql_fdw with dolthub/go-mysql-server

## usage

```code
CREATE EXTENSION mysql_fdw;
CREATE SERVER mysql_server  
FOREIGN DATA WRAPPER mysql_fdw
OPTIONS (host '<replace with your server ip>', port '3306');  
CREATE USER MAPPING FOR postgres 
SERVER mysql_server  
OPTIONS (username 'root', password 'dalong');  
CREATE FOREIGN TABLE mytest (name text, email text) server mysql_server options (dbname 'test', table_name 'mytable'); 
```

## query

> in pg server

```code
select * from mytest
```