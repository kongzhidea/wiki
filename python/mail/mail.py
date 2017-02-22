#!/usr/bin/env python
# -*- coding: utf-8 -*-
#导入smtplib和MIMEText
import smtplib
from email.mime.text import MIMEText

#要发给谁
mail_to=["mail_to@qq.com"]

def send_mail(to_list,sub,content):
    #设置服务器，用户名、口令以及邮箱的后缀
    mail_host="smtp.163.com"
    mail_user="username"
    mail_pass="bloghotword"
    mail_postfix="163.com"
    me=mail_user+"<"+mail_user+"@"+mail_postfix+">"
    msg =MIMEText(content)
    msg['Subject']=sub
    msg['From']= me
    msg['To']= ".".join(to_list)
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
