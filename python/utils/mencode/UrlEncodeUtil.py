# -*- coding: utf-8 -*-

from urllib import urlencode
from urllib import quote
from urllib import unquote

#encode dict， dict转成query格式，并且转义。
def encodeDict(dict):
    return urlencode(dict)

#encode string
def quoteString(s):
    return quote(s)

#decode string
def unquoteString(s):
    return unquote(s)

if __name__ == "__main__":
    print encodeDict({"a":1,"b":"孔&1"})  #a=1&b=%E5%AD%94%261
    print quoteString("孔&1")  #%E5%AD%94%261
    print unquoteString("%E5%AD%94%261")  #孔&1
    print unquoteString("a=1&b=%E5%AD%94%261")  #a=1&b=孔&1