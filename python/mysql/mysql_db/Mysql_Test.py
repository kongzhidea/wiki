# -*- coding: utf-8 -*-

import sys
import MySQLdb as mdb
import ConfigParser
import os
import time,datetime
import MysqlUtil

def config():
    cf = ConfigParser.ConfigParser()
    cf.read(sys.argv[1])

    mysql_ip=cf.get("db", "mysql_ip")
    mysql_username=cf.get("db", "mysql_username")
    mysql_password=cf.get("db", "mysql_password")
    mysql_database=cf.get("db", "mysql_database")
    return (mysql_ip,mysql_username,mysql_password,mysql_database)


def main():
    mysql_ip,mysql_username,mysql_password,mysql_database = config()
    mysql = MysqlUtil.MysqlUtil(host=mysql_ip,user=mysql_username,passwd=mysql_password,database=mysql_database)
    try:
        sql = ""
        rets = mysql.queryAll(sql,True)
    except mdb.Error, e:
        print "Error %d: %s" % (e.args[0],e.args[1])
        sys.exit(1)
    finally:
        if mysql:
            mysql.close()

if __name__ == "__main__":
    main()
