#!/usr/bin/env python
#coding:utf8


import smtplib
from email.mime.text import MIMEText
def sendmail():
    sender='sender'
    receivers=['receivers']
    subject="邮件主题_test"
    username="sender@..com"
    password="pwd"
    content="邮件正文"
    msg=MIMEText(content,'plain','utf-8')
    msg['Subject']=subject
    msg['to']=";".join(receivers)
    msg['from']=sender

    smtp=smtplib.SMTP_SSL('smtp.xx.com:465')
    smtp.login(username,password)
    smtp.sendmail(sender,receivers,msg.as_string())
    smtp.quit()

sendmail()