
### [21分钟 MySQL 入门教](http://www.cnblogs.com/mr-wid/archive/2013/05/09/3068229.html)

## mysql启动和关闭
```
window下:
	停止服务:
		bin>net stop mysql
		bin>mysqladmin -u root shutdown
	启动服务:
		bin>net start mysql
		计算机管理->服务->启动mysql服务

如果下载的是xampp的nysql，则直接输入mysqld就启动了mysql服务。

bin>mysqladmin -u root shutdown语句中的root不能换成其他用户。
net start mysql和net stop mysql中的mysql 指的是mysql服务，而不是特定的数据库名字。

linux下:
	1.启动
	   sudo service mysql start
	   sudo /etc/init.d/mysql start
	2、停止
	　　mysqladmin -u root -p shutdown
	   sudo /etc/init.d/mysql stop
```
 

## mysql入门级别命令

### 连接mysql
* mysql -h主机名 -u用户名 -p密码 [数据库]
* phpmyadmin
* navicate 
* sequel pro (适合mac)

### 查看所有数据库
* show databases;

### 切换数据库
* use database

### 查看所有表
* show tables;

### sql中语句以分号结尾

### 展示表结构
* desc tblname

### 展示建表 语句
* shwo create table tblname;

### 查看索引
* show index from tblname;
* show keys from tblname;

### 查看所有列
* show columns from tblname;  

### 数据库字符集，建议 utf-8
* 查看msyql编码：show variables like 'character%';
* 设置mysql编码为utf-8

```
set names "utf8";

它相当于下面的三句指令：
SET character_set_client = utf8;
SET character_set_results = utf8;
SET character_set_connection = utf8;
```

### 修改表名
* alter table tblname rename to tblname1;

### 添加列，comment标示注释
* alter table 表名 add column 列名 varchar(30) comment '';

### 删除列
* alter table 表名 drop column 列名;

### 修改列名，后面需要加上类型和comment。
* alter table 表名 change column1 column1 int;

### 修改列属性
* alter table 表名 modify name varchar(22);

### 加索引：普通索引，唯一索引，主键，多列索引
* alter table tablename add index indexname (columnlist) ;
* alter table tablename add unique (columnlist) ;
* alter table tablename add primary key (columnlist) ;
* alter table tablename add index indexname ( column1, column2, column3 )

### 删除索引
* drop index indexname on tablename ;
* alter table tablename drop index indexname ;
* alter table tablename drop primary key ; 
* 在前面的两条语句中，都删除了tablename中的索引indexname。而在最后一条语句中，只在删除PRIMARY KEY索引中使用，因为一个表只可能有一个PRIMARY KEY索引，因此不需要指定索引名。

### 表改名
* rename table tblname1 to tblname2 ;

### 删除表
* drop table  tblname;

### 删除表中所有数据
* truncate table tblname;
* truncate删除表中的所有数据，不能与where一起使用，不可以rollback，只删数据，不删表结构定义，truncate比delete语句的速度要快得多。

### 克隆表
* 复制结构：create table user2 like user;
* 复制数据：insert into user2 select * from user;

### select时候  最后加上 \G，可以将每行展示每一条数据，适合列比较多的表。

### mysql 重新设定auto_increment
```
查询这个表的auto_increment值：
select AUTO_INCREMENT from information_schema.tables where TABLE_NAME="表名"

或者: show create table tableName

修改表的auto_increment值：
ALTER TABLE tableName auto_increment=number ;
```

### 刷新数据库
* flush privileges;

### 显示数据库版本
* select version();

### 显示当前时间
* select current_date;    // yyyy-MM-dd
* select now();

### 显示数据库状态
* status;
* \s

### sql样例
```
CREATE DATABASE 数据库名 DEFAULT CHARACTER SET utf8 COLLATE utf8_bin

CREATE TABLE 表名 (
  id bigint(16) NOT NULL AUTO_INCREMENT,
  account_id int(11) NOT NULL,
  latest_back_img varchar(256) COLLATE utf8_bin NOT NULL,
  latest_title varchar(256) COLLATE utf8_bin NOT NULL,
  exts varchar(3000) COLLATE utf8_bin NOT NULL,
  update_time datetime NOT NULL ,
  PRIMARY KEY (id),
  KEY account_id (account_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin


alter table account_worldnews add column eid bigint(16) default 0
alter table account_worldnews add index index_name(eid)

ALTER TABLE image_text_items CHANGE author image_url VARCHAR( 1024 ) COLLATE utf8_bin DEFAULT NULL


CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(32) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `account` (`account`),
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin
```


### 导出数据库
* mysqldump -h主机名 -u用户名 -p密码 [数据库] > database.sql
* mysqldump  后面还可以加上tableName，只导出该表。加上--no-data参数后 只导出结构，不导出数据。  
* mysql -h主机名 -u用户名 -p密码 [数据库] -e 'sql' > sql文件
* mysql -h主机名 -u用户名 -p密码 [数据库]  > sql文件，  然后输入sql，ctrl+d退出输入。
    * 如果想获得表格式输出，加-t 参数。
    * 有的可能默认是表格，加-t参数 输出tab格式。-H/--html输出html格式，-X/--xml输出xml格式    -ss/--skip-column-names 禁止输出头部信息,  -ss和 |cat -n配合使用可以 输出行号

### 导入sql语句
* source sql文件（推荐使用绝对地址）
* mysql -h主机名 -u用户名 -p密码 [数据库] < sql文件
* cat sql文件 | mysql -h主机名 -u用户名 -p密码 [数据库]
* echo 'sql' | mysql -h主机名 -u用户名 -p密码 [数据库]
* mysql -h主机名 -u用户名 -p密码 [数据库] -e 'sql' 

### 创建视图
```
不指定类型：
create view teams  as select  distinct termno from result

指定类型：
create ALGORITHM=merge view teams as select  distinct termno from result

ALGORITHM有三个参数分别是：merge、TEMPTABLE、UNDEFINED
```

## mysql 终端快捷操作
```
快捷键
Ctrl+A  行首
Ctrl+E  行尾
Ctrl+D 删除光标所在字符

输入多行时候，输入\c可以去掉这多行的输入。

垂直展示每行 在行尾加上\G可以 .

NULL值判断，不能用=和！=，应该用 is null和is not null
```




