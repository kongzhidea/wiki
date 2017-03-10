#!/usr/bin/env python
#coding:utf8
import smtplib
from email.mime.text import MIMEText

def sendmail(receivers,subject,content):
    username="sender@..com"
    password="pwd"
    msg=MIMEText(content,'plain','utf-8')
    #msg["Content-Type"] = 'text/html; charset=utf-8'  #设置富文本
    msg['Subject']=subject
    msg['to']=";".join(receivers)
    msg['from']=username

    smtp=smtplib.SMTP_SSL('smtp.xx.com:465')
    smtp.login(username,password)
    smtp.sendmail(username,receivers,msg.as_string())
    smtp.quit()

if __name__ == "__main__":
    sendmail(['receivers'],"邮件主题_test","邮件正文")
