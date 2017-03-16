# -*- coding: utf-8 -*-
import time


# 返回unxi时间戳，long型，精确到秒，date为struct_time格式，不传则表示当前时间
def getUnixTime(date=None):
    if date:
        return long(time.mktime(date))
    else:
        return long(time.time())


'''
时间戳转 struct_time

时间元组，月份1-12，time.struct_time
tm_year, tm_mon, tm_mday, tm_hour, tm_min, tm_sec, tm_wday(一周的第几日，从0开始:周一), tm_yday(一年的第几日，从1开始)
'''


# 时间戳转struct_time seconds如果不传则表示当前时间：time.localtime()当前时间
def localTime(seconds=None):
    return time.localtime(seconds)


'''
python中时间日期格式化符号：
%y 两位数的年份表示（00-99）
%Y 四位数的年份表示（000-9999）
%m 月份（01-12）
%d 月内中的一天（0-31）
%H 24小时制小时数（0-23）
%I 12小时制小时数（01-12）
%M 分钟数（00=59）
%S 秒（00-59）
%a 本地简化星期名称
%A 本地完整星期名称
%b 本地简化的月份名称
%B 本地完整的月份名称
%c 本地相应的日期表示和时间表示
%j 年内的一天（001-366）
%p 本地A.M.或P.M.的等价符
%U 一年中的星期数（00-53）星期天为星期的开始
%w 星期（0-6），星期天为星期的开始
%W 一年中的星期数（00-53）星期一为星期的开始
%x 本地相应的日期表示
%X 本地相应的时间表示
%Z 当前时区的名称
%% %号本身
'''


# date为struct_time类型，不传表示当前时间，返回日期str
def defaultTime(date=None):
    if date:
        return time.strftime("%Y-%m-%d %H:%M:%S", date)
    else:
        return time.strftime("%Y-%m-%d %H:%M:%S")

# date为struct_time类型，不传表示当前时间，返回时间str
def defaultDate(date=None):
    if date:
        return time.strftime("%Y-%m-%d", date)
    else:
        return time.strftime("%Y-%m-%d")

# date为struct_time类型，不传表示当前时间，返回str
def format(date,format):
    return time.strftime(format, date)

# 字符串转 struct_time， strdate如：2017-03-16
def parseDate(strdate):
    return time.strptime(strdate, "%Y-%m-%d")

# 字符串转 struct_time， strdate如：2017-03-16 19:26:07
def parseTime(strdate):
    return time.strptime(strdate, "%Y-%m-%d %H:%M:%S")

# 字符串转 struct_time
def parse(strdate,format):
    return time.strptime(strdate,format)