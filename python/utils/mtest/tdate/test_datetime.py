# -*- coding: utf-8 -*-

import time
import datetime

#今天日期
print datetime.date.today()  #2013-05-22

#通过时间戳返回日期
print datetime.date.fromtimestamp(time.time())  #2013-05-22


#自定义日期
now = datetime.date(2000,4,1)

print now  #2000-04-01

print now.year,now.month,now.day  #2000 4 1

#替换日期元素
tommorrow = now.replace(day=2)
print tommorrow  #2000-04-02

#返回日期的struct_time对象
print now.timetuple()  #time.struct_time(tm_year=2000, tm_mon=4, tm_mday=1, tm_hour=0, tm_min=0, tm_sec=0, tm_wday=5, tm_yday=92, tm_isdst=-1)

#返回星期几
print now.weekday()  #星期一=0
print now.isoweekday()  #星期一=1



#返回日期元祖
print now.isocalendar()  #(2000, 13, 6)

#返回格式化日期
print now.isoformat()  #2000-04-01


#日期之间进行操作
now = datetime.date.today()
tommorrow = now.replace(day = now.day+1)
detdata = tommorrow - now
print now,tommorrow,detdata  #2013-05-22 2013-05-23 1 day, 0:00:00
print now-tommorrow  #-1 day, 0:00:00
print now+detdata #2013-05-23
print now < tommorrow  #True



#日期加减天数
start = datetime.date(2013,5,22)
end = start + datetime.timedelta(days = 10)
print start,end #2013-05-22 2013-06-01


#详细时间
dt = datetime.datetime.now()
print dt  #2013-05-22 09:59:24.484000
print '(%Y-%m-%d %H:%M:%S %f): ',dt.strftime('%Y-%m-%d %H:%M:%S %f') #(%Y-%m-%d %H:%M:%S %f):  2013-05-22 09:59:24 484000

print datetime.datetime.strptime("2013-05-22 09:59:24","%Y-%m-%d %H:%M:%S")
