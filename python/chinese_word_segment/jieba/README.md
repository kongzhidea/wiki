#### wiki：http://www.cnblogs.com/kaituorensheng/p/3595879.html

```
组件只提供jieba.cut 方法用于分词
cut方法接受两个输入参数：
　　第一个参数为需要分词的字符串
　　cut_all参数用来控制分词模式
待分词的字符串可以是gbk字符串、utf-8字符串或者unicode
jieba.cut返回的结构是一个可迭代的generator，可以使用for循环来获得分词后得到的每一个词语(unicode)，也可以用list(jieba.cut(...))转化为list

```

### github：[https://github.com/kongzhidea/jieba](https://github.com/kongzhidea/jieba)


##### 安装：sudo  pip install jieba

```
# -*- coding: utf-8 -*-

import jieba

if __name__ == "__main__":
    # seg_list = jieba.cut("我来到北京清华大学", cut_all=True)
    # print "Full Mode:", ' '.join(seg_list)

    seg_list = jieba.cut("我来到北京清华大学")
    print "Default Mode:", ' '.join(seg_list)
```