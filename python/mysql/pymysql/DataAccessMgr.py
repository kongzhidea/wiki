# -*- coding: utf-8 -*-
import pymysql.cursors
import logging
import sys
import ConfigParser

'''
pymysql.cursors.DictCursor 返回dict格式
connect中cursorclass不传  返回元祖格式

后续的查询中更改返回格式
cursor = connection.cursor(pymysql.cursors.DictCursor)

切换数据库
conn.select_db(database)
'''

'''
如果sql中带参数，可以先将字符串格式化，再execute， 推荐！！

也可以直接将sql和args都传过来，再execute，args可以是列表，如果只有1个参数，可以不用列表。
占位符最好使用%s， %s会在参数外加上引号，当做字符串处理。
如 execute("update  user  set name=%s where id = %s",["kk",1]))
'''

'''
字符串格式化：

print "%s %d"%("kong",123)  ## 参数格式要完全对应
print '{name} {age}'.format(age=22, name='管理员')

print '{0} {1} {2}'.format('a', 'b', 'c')
print '{} {} {}'.format('a', 'b', 'c')
'''

'''
调用样例:

from DataAccessMgr import DataAccessMgr

client = DataAccessMgr(host=mysql_ip,user=mysql_username,passwd=mysql_password,database=mysql_database)
print client.insert("insert into user(path,gmt_create) values('%s',now())" %("/k6"))
client.show(client.queryAllObject("select * from user"))
'''
class DataAccessMgr:
    def __init__(self, host=None, user=None, passwd=None, database=None, port=3306, charset="utf8",logger=None, showSql=True):
        self.conn = pymysql.connect(host=host,
                             user=user,
                             password=passwd,
                             database=database,
                             port=port,
                             charset=charset)
        if logger == None and showSql:
            self.logger = initlog("DataAccessMgr")
        else:
            self.logger = None

        return

    def _showdict(self, dit):
        print "{",
        for key, value in dit.items():
            print key, ":", value, ",",
        print "}"

    # 遍历返回结果
    def show(self, result=None):
        print type(result)

        if isinstance(result, list):
            print "length =", len(result)

            alldict = True
            for dit in result:
                if not isinstance(dit, dict):
                    alldict = False
                    break
            if alldict:
                for dit in result:
                    self._showdict(dit)
            else:
                print "[",
                for item in result:
                    print item, ",",
                print "]"
            return

        if isinstance(result, dict):
            self._showdict(result)
            return

        print result

    def log(self, sql, args=None):
        if self.logger:
            self.logger.info(sql)

            if args:
                self.logger.info("args=" + str(args))

    def close(self):
        self.conn.close()

    # useGeneratedKeys=True 返回自增id
    def insert(self,sql,useGeneratedKeys=False):
        if self.logger:
            self.logger.info(sql)

        cursor = self.conn.cursor()
        cursor.execute(sql)

        lastId = None
        if useGeneratedKeys:
            lastId = self.conn.insert_id()

        self.conn.commit()

        return lastId

    # 返回list(dict)类型
    def queryAllObject(self,sql, args=None):
        self.log(sql, args)

        cursor = self.conn.cursor(pymysql.cursors.DictCursor)
        cursor.execute(sql, args)
        rows = cursor.fetchall()
        return rows

    # 返回dict类型
    def queryObject(self,sql, args=None):
        self.log(sql, args)

        cursor = self.conn.cursor(pymysql.cursors.DictCursor)
        cursor.execute(sql, args)
        rows = cursor.fetchone()
        return rows

    # 返回单独的值
    def queryValue(self,sql, args=None):
        self.log(sql, args)

        cursor = self.conn.cursor()
        cursor.execute(sql, args)
        rows = cursor.fetchall()

        length = len(rows)
        if length != 1:
            raise RuntimeError("返回结果长度必须为1，当前长度为" + str(length))

        item = rows[0]

        length = len(item)
        if length != 1:
            raise RuntimeError("查询参数长度必须为1，当前长度为" + str(length))

        return item[0]

    # 返回list<单独值>
    def queryAllValue(self,sql, args=None):
        self.log(sql, args)

        cursor = self.conn.cursor()
        cursor.execute(sql, args)
        rows = cursor.fetchall()

        if len(rows) > 0:
            length = len(rows[0])
            if length != 1:
                raise RuntimeError("查询参数长度必须为1，当前长度为" + str(length))

        result = []
        for item in rows:
            length = len(item)
            if length != 1:
                raise RuntimeError("查询参数长度必须为1，当前长度为" + str(length))

            result.append(item[0])

        return result

    # 返回受影响条数
    def execute(self,sql,args=None):
        if self.logger:
            self.logger.info(sql)

        cursor = self.conn.cursor()
        c = cursor.execute(sql, args)

        self.conn.commit()
        return c

'''
从配置文件解析， mysql.ini

client = buildDataAccessMgr("db.ini")

[db]
mysql_ip = ***
mysql_username = ***
mysql_password = ***
mysql_database = ***
'''
def buildDataAccessMgr(filePath):
    cf = ConfigParser.ConfigParser()
    cf.read(filePath)

    mysql_ip = cf.get("db", "mysql_ip")
    mysql_username = cf.get("db", "mysql_username")
    mysql_password = cf.get("db", "mysql_password")
    mysql_database = cf.get("db", "mysql_database")

    return DataAccessMgr(host=mysql_ip,user=mysql_username,passwd=mysql_password,database=mysql_database)

def initlog(name=__file__, logfile = None):
    # 生成一个日志对象
    logger = logging.getLogger(name)

    # 成一个格式器，用于规范日志的输出格式。如果没有这行代码，那么缺省的
    # 格式就是："%(message)s"。也就是写日志时，信息是什么日志中就是什么，
    # 没有日期，没有信息级别等信息。logging支持许多种替换值，详细请看
    # Formatter的文档说明。这里有三项：时间，信息级别，日志信息

    # initlog()中的name属性对应formatter中的$(name)s
    formatter = logging.Formatter('%(asctime)s %(levelname)s %(filename)s:%(lineno)s - %(message)s')

    # 生成一个Handler。logging支持许多Handler，
    # 象FileHandler, SocketHandler, SMTPHandler等，我由于要写
    # 输出到文件 文件就使用了FileHandler。
    # logfile是一个全局变量，它就是一个文件名，如：'crawl.log'
    if logfile:
        hdlr = logging.FileHandler(logfile)

        # 将格式器设置到处理器上
        hdlr.setFormatter(formatter)

        # 将处理器加到日志对象上
        logger.addHandler(hdlr)

    #输出到控制台
    # 创建一个控制台的handler
    c_handler = logging.StreamHandler(sys.stdout);
    #c_handler.setLevel(logging.WARNING);
    c_handler.setFormatter(formatter)


    logger.addHandler(c_handler)
    # 设置日志信息输出的级别。logging提供多种级别的日志信息，如：NOTSET,
    # DEBUG, INFO, WARNING, ERROR, CRITICAL等。每个级别都对应一个数值。
    # 如果不执行此句，缺省为30(WARNING)。可以执行：logging.getLevelName
    # (logger.getEffectiveLevel())来查看缺省的日志级别。日志对象对于不同
    # 的级别信息提供不同的函数进行输出，如：info(), error(), debug()等。当
    # 写入日志时，小于指定级别的信息将被忽略。因此为了输出想要的日志级别一定
    # 要设置好此参数。这里我设为NOTSET（值为0），也就是想输出所有信息
    logger.setLevel(logging.INFO)   ## info级别
    return logger
