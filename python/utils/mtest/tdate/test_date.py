# -*- coding: utf-8 -*-

import time
from mdate import DateUtils

ticks = time.time()
print "当前时间戳为:", ticks,type(ticks)  #float类型，精确到秒

print "当前时间戳为2:",DateUtils.getUnixTime()

print "当前时间戳3：",time.mktime(time.localtime(time.time()))
print "当前时间戳4：", DateUtils.getUnixTime(time.localtime(time.time()))

localtime = time.localtime(1)
print "本地时间为 :", localtime

print DateUtils.localTime()

'''
ctime 时间戳转时间str，格式为：Thu Mar 16 19:18:39 2017
asctime struct_time转时间str，格式：Thu Mar 16 19:18:39 2017
'''

slocaltime = time.asctime( time.localtime(time.time()) )
print "本地时间为 :", slocaltime,type(slocaltime)  #str   Thu Mar 16 19:18:39 2017

slocaltime = time.ctime( time.time() )
print "本地时间2为 :", slocaltime,type(slocaltime)  #str   Thu Mar 16 19:18:39 2017

print time.strftime("%Y-%m-%d %H:%M:%S",localtime)

print DateUtils.defaultTime(localtime)
print DateUtils.defaultTime()

print DateUtils.defaultDate(localtime)
print DateUtils.defaultDate()

print DateUtils.format(time.localtime(), "%Y-%m-%d %H:%M:%S")

print time.strptime("2017-03-1 19:22:57","%Y-%m-%d %H:%M:%S")
print DateUtils.parseDate("2017-03-1")
print DateUtils.parseTime("2017-03-16 19:22:57")


print DateUtils.defaultDate(DateUtils.localTime())