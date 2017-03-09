# -*- coding: utf-8 -*-

import json

# string -> dict
def load(s):
    return json.loads(s)

# dict/list->string； 对key排序 sort_keys=True
def dump(j,pretty=False):
    if not pretty:
        return json.dumps(j)
    return json.dumps(j, indent=4, separators=(',', ': '))

if __name__ == "__main__":
    dit = load('{"a":"aa","b":"123","rr":{"a1":"11","b1":"22"}}')
    print dit
    print dump(dit)
    print dump(dit,True)
    print dump([1,2,3,{"a":1}],True)