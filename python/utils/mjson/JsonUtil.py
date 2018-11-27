# -*- coding: utf-8 -*-

import json


# string -> dict
def loads(s):
    return json.loads(s)


# load和dump 针对文件句柄，如  json.load(open(filename))


# dict/list->string； 对key排序 sort_keys=True
# pretty:true 格式化输出， ensure_ascii：False 直接打印中文
# 没有default参数，只能适用于dict和list等基本类型，不能适用与Class。default可以适用与Class转json，可以自定义函数。
def dumps(j, pretty=False, ensure_ascii=False, default=None):
    print type(j)
    if default is None:
        default = lambda obj: obj.__dict__
    if not pretty:
        return json.dumps(j, ensure_ascii=ensure_ascii, default=default)
    return json.dumps(j, indent=4, separators=(',', ': '), ensure_ascii=ensure_ascii, default=default)


if __name__ == "__main__":
    dit = loads('{"a":"aa","b":"123","rr":{"a1":"11","b1":"22"}}')
    print dit
    print dumps(dit)
    print dumps(dit, True)
    print dumps([1, 2, 3, {"a": "孔智慧"}])  # [1, 2, 3, {"a": "孔智慧"}]
    print dumps([1, 2, 3, {"a": "孔智慧"}], ensure_ascii=True)  # [1, 2, 3, {"a": "\u5b54\u667a\u6167"}]
    print dumps([1, 2, 3, {"a": "孔智慧"}], pretty=True)
