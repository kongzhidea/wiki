


## 第三方json库
* [ujson](https://github.com/esnme/ultrajson)
    * pip install ujson

```
>>> import ujson
>>> ujson.dumps([{"key": "value"}, 81, True])
'[{"key":"value"},81,true]'
>>> ujson.loads("""[{"key": "value"}, 81, true]""")
[{u'key': u'value'}, 81, True]
```
