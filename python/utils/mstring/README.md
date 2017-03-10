
1.文件头编码:<br/>
# -*- coding: utf-8 -*- <br/>

2.python中  unicode字符串不能和普通字符串 join，例如：u"aa" + "空" 会报错， <br/>
可以加上以下代码解决此问题。<br/>
```
import sys
reload(sys)
sys.setdefaultencoding('utf8')
```
