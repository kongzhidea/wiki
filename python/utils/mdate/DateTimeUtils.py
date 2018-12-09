# -*- coding: utf-8 -*-

import datetime
import time

'''
构建datetime和date格式， datetime：2017-03-16 20:26:51.892537， date：2017-03-16
dateC = datetime.datetime(2013, 9, 5, 11, 00, 00)
dateD = datetime.date(2013, 9, 5)

datetime.datetime.today()
datetime.date.today()


now =datetime.datetime.today()
print now.year,now.month,now.day,now.hour,now.minute,now.second

'''


# date为datetime或date类型，返回时间戳，long型，精确到秒，date为null则表示当前时间戳
def getUnixTime(date=None):
    if date:
        return long(time.mktime(date.timetuple()))
    else:
        return long(time.time())


# date为datetime或date类型，返回struct_time类型
def getTimeTuple(date):
    return date.timetuple()


# date为datetime或date类型，返回日期str
def defaultDate(date=None):
    if date:
        return date.strftime("%Y-%m-%d")
    else:
        return datetime.date.today().strftime("%Y-%m-%d")


# date为datetime或date类型，返回时间str
def defaultTime(date=None):
    if date:
        return date.strftime("%Y-%m-%d %H:%M:%S")
    else:
        return datetime.date.today().strftime("%Y-%m-%d %H:%M:%S")


# date为datetime或date类型，返回str
def format(date, format):
    return date.strftime(format)


# 返回datetime类型
def parseDate(strdate):
    return datetime.datetime.strptime(strdate, "%Y-%m-%d")


# 返回datetime类型
def parseTime(strdate):
    return datetime.datetime.strptime(strdate, "%Y-%m-%d %H:%M:%S")


# 返回datetime类型
def parse(strdate, format):
    return datetime.datetime.strptime(strdate, format)


# date为datetime或date类型
def addDay(date, delta):
    return date + datetime.timedelta(days=delta)


if __name__ == "__main__":
    dateC = datetime.datetime(2013, 9, 5, 11, 00, 00)
    dateD = datetime.date(2013, 9, 5)

    print "时间戳1", getUnixTime(dateC), type(dateC)
    print "时间戳2", getUnixTime(dateD), type(dateD)
    print "时间戳3", getUnixTime()

    print defaultDate(dateC)
    print defaultDate(dateD)
    print defaultDate()

    print defaultTime(dateC)
    print defaultTime(dateD)
    print defaultTime()

    print "..", format(dateC, "%Y-%m-%d %H:%M")

    print datetime.datetime.today()
    print datetime.date.today()

    print getTimeTuple(datetime.datetime.today())

    now = datetime.datetime.today()

    print now.year, now.month, now.day, now.hour, now.minute, now.second

    print parseDate("2017-03-16"), type(parseDate("2017-03-16"))
    print parseTime("2013-09-05 11:00:00"), type(parseTime("2013-09-05 11:00:00"))

    print defaultTime(addDay(now, -20))
