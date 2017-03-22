# -*- coding: utf-8 -*-

import MySQLdb as mdb
import sys
import ConfigParser
import os


class MysqlUtil:
    def __init__(self,host=None,user=None,passwd=None,database=None,port=3306,charset="utf8"):
        self.con = mdb.connect(host=host,user=user,passwd=passwd,port=port,charset=charset)
        self.con.select_db(database)

    def commit(self):
        self.con.commit()

    def close(self):
        self.con.close()

    #如果是修改操作，需要execute后需要执行commit
    def execute(self,sql):
        cur = self.con.cursor()
        cur.execute(sql)

    def queryAll(self,sql,showsql = False):
        if showsql:
            print sql
        # cur = self.con.cursor()  #返回元祖
        #取数据,返回dict
        cur = self.con.cursor(mdb.cursors.DictCursor)
        cur.execute(sql)
        #多行数据  使用fetchall
        rows = cur.fetchall()
        return rows

    def queryOne(self,sql,showsql = False):
        if showsql:
            print sql
        #cur = self.con.cursor() #返回元祖
        #取数据，返回dict
        cur = self.con.cursor(mdb.cursors.DictCursor)
        cur.execute(sql)
        #多行数据  使用fetchall
        row = cur.fetchone()
        return row
