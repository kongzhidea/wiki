#!/usr/bin/env python
# -*- coding: utf-8 -*-
#导入smtplib和MIMEText
import smtplib
from email.mime.text import MIMEText
from email.MIMEMultipart import MIMEMultipart
from email.Header import Header
import datetime

#要发给谁
mail_to=["mail_to@qq.com"]

def send_mail(to_list,sub,content):
    #设置服务器，用户名、口令以及邮箱的后缀
    mail_host="smtp.163.com"
    mail_user="username"
    mail_pass="password"
    mail_postfix="163.com"
    me=mail_user+"<"+mail_user+"@"+mail_postfix+">"
    #创建一个带附件的实例
    msg = MIMEMultipart()
    txt=MIMEText(content)
    txt.set_charset("gb2312")
    msg.attach(txt)
    #msg.attach(MIMEText(body))
    #构造附件
    att = MIMEText(open('/data/data.csv', 'rb').read(), 'base64', 'gb2312')
    att["Content-Type"] = 'application/octet-stream'
    att["Content-Disposition"] = 'attachment; filename="data.csv"'
    msg.attach(att)

    #加邮件头
    msg['to'] = ".".join(to_list)
    msg['from'] = me
    msg['subject'] = Header('REPORT (' + str(datetime.date.today()) + ')', 'utf-8')
    try:
        s = smtplib.SMTP()
        s.connect(mail_host)
        s.login(mail_user,mail_pass)
        s.sendmail(me, to_list, msg.as_string())
        s.close()
        return True
    except Exception, e:
        print str(e)
        return False
if __name__ =='__main__':
    if send_mail(mail_to,"hello","this is python sent"):
        print "发送成功"
    else:
        print "发送失败"
