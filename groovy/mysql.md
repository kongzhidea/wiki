

```

1、select
不用关闭 Connection，也不用关闭 ResultSet

import groovy.sql.Sql

def sql = Sql.newInstance(url, username,
        password, driverClassName)
sql.eachRow("select * from user") {
    println it  // it.property
}

2.insert
def id = 120
def name = "xxx"
sql.execute("insert into user (id, name) values (${id}, ${name})")

3.update
sql.execute("update user set name=${name} where id = ${id}")

4.delete
sql.execute("delete from user where id = ${id}")

5.PreparedStatement
sql.execute("insert into user (id, name) values (?,?)", id, name)

```