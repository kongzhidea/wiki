# mysql-python,见MysqlUtil.py
* 安装easy_install模块:    安装第三方模块->easy_install
* easy_install mysql-python
* 如果上一步报错，则
* apt-get install python-dev   
* sudo apt-get install libmysqlclient-dev
* 再安装mysql-python模块
* 执行更新操作后需要commit


# 推荐PyMySql
* https://github.com/kongzhidea/PyMySQL
* python版本2.7
* sudo python setup.py install
* 执行更新操作后需要commit
```
样例1：
# -*- coding: utf8 -*-
#!/usr/bin/env python
import pymysql
conn = pymysql.connect(host='localhost', port=3306, user='root', passwd='', db='mysql')
cur = conn.cursor()
cur.execute("SELECT Host,User FROM user")
print(cur.description)
print()
for row in cur:
    print(row)
cur.close()
conn.close()


样例2：

# -*- coding: utf8 -*-

import pymysql.cursors

connection = pymysql.connect(host=mysql_ip,
                             user=mysql_username,
                             password=mysql_password,
                             db=mysql_database,
                             charset="utf8",
                             cursorclass=pymysql.cursors.DictCursor)

with connection.cursor() as cursor:
    # Read a single record
    sql = "SELECT * from user"
    cursor.execute(sql)
    result = cursor.fetchone()  #取1条
    print(result)
    # for row in cursor:    # 取多条
    #     print row
connection.close()

```
