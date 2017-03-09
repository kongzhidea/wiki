# -*- coding: utf-8 -*-
import ConfigParser

if __name__ == "__main__":
    # 必须要有section
    cf = ConfigParser.ConfigParser()
    filename = "conf.ini"
    cf.read(filename)

    # 返回所有的section
    s = cf.sections()  # ['db', 'concurrent']
    print 'section:', s

    o = cf.options("db")  # ['db_host', 'db_port', 'db_user', 'db_pass']
    print 'options:', o

    v = cf.items("db")  # [('db_host', '127.0.0.1'), ('db_port', '3306'), ('db_user', 'root'), ('db_pass', 'password')]
    print 'db:', v

    print '-' * 60

    # 可以按照类型读取出来
    db_host = cf.get("db", "db_host")
    db_port = cf.getint("db", "db_port")
    db_user = cf.get("db", "db_user")
    db_pass = cf.get("db", "db_pass")

    # 返回的是整型的
    threads = cf.getint("concurrent", "thread")
    processors = cf.getint("concurrent", "processor")

    print "db_host:", db_host
    print "db_port:", db_port
    print "db_user:", db_user
    print "db_pass:", db_pass

    print "thread:", threads
    print "processor:", processors