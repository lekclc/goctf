# set

## mysql

在第一次允许后执行
在mysql容器中
```
mysql -uroot -p
GRANT ALL ON *.* TO 'root'@'%';
flush privileges;
```