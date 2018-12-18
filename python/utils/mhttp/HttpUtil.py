# -*- coding: utf-8 -*-
import urllib2
import urllib
import requests

# easy_install requests

def doGet(url):
    response = urllib2.urlopen(url)
    html = response.read()
    return html

def doGet2(url):
    req = urllib2.Request(url)
    response = urllib2.urlopen(req)
    html = response.read()
    return html

#param dict格式
def doPost(url,param = None):
    # 进行url编码，转换为str1=1&str2=2的url参数形式，同时进行html转义
    if param == None:
        param = {}
    data = urllib.urlencode(param)

    req = urllib2.Request(url, data)
    response = urllib2.urlopen(req)
    html = response.read()
    return html

#原生post格式,data string格式
def sendPostRequest(url,data):
    response = requests.post(url, data=data)
    # response.text 返回unicode类型，response.content 返回str类型
    return response.text

if __name__ == "__main__":
    print doGet2("http://localhost:9999/json2")

    print "_" * 60
    print doPost("http://localhost:9999/json2")
    print doPost("http://localhost:9999/json2",{"id":1})

    print "_" * 60
    print sendPostRequest("http://localhost:9999/postdata", "sfsdfs")

    print "_" * 60
    ## get with cookie
    req = urllib2.Request("http://localhost:9999/json2")
    req.add_unredirected_header("Cookie", "token=1")
    req.add_unredirected_header("user-agent", "python_test")
    print urllib2.urlopen(req).read()

    print "_" * 60
    # ## post2 with cookie
    headers = {'Cookie': 'token=1;t=2'}
    print requests.post("http://localhost:9999/postdata", data="data",headers=headers).text

