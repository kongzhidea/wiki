# -*- coding: utf-8 -*-
import urllib
from urlparse import parse_qs
from urlparse import urlparse

def getDomain(url):
    parsed = urlparse(url)
    return parsed.scheme + "://" + parsed.netloc

# 返回列表，例如str=kk&str=1->['kk', '1']
def getParameter(url,key):
    parsed = urlparse(url)
    # parse_qsl返回[(key,value),(key,value)] 格式
    # parse_qs返回{key:[values]} 格式
    qs = parse_qs(parsed.query)
    return qs.get(key)


if __name__ == "__main__":
    print urllib.urlencode({"a":1,"b":2})  ## dict=> string，query形式

    #urlparse(url)=ParseResult(scheme='http', netloc='www.kk.com', path='/path', params='param', query='query=arg&str=kk', fragment='frag')
    print getDomain("http://www.baidu.com/index.php?s=word")

    print getParameter("http://www.baidu.com/path?query=arg&str=kk&str=1","str")

