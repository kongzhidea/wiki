## package用法

### 样例1：
```
import "fmt"

使用时候  fmt.Println(1)
```

### 样例2，引入多个package
```
import (
    "fmt"
    "os"
)
```

### GOPATH，import包
* GOPATH目录中，工程代码放到src目录下。
* import时候，路径填写 被依赖go文件的路径，从src开始（不包含src，不包含文件名）
* 如，文件结构为  src/utils/list/GoArrayList.go，则 import "utils/list"
* go文件 package名称可以和目录名称不同，调用的时候  packageName.FuncName()

```
export GOROOT="/usr/local/go"
export GOPATH=$GOPATH:"/Users/kongzhihui/workspace/go"
export PATH=$PATH:$GOROOT/bin

GOROOT为系统自带包，GOPATH为自定义依赖包。
```

### import包后，会执行其init方法。

### Go中如果函数名的首字母大写，表示该函数是公有的，可以被其他程序调用，如果首字母小写，该函数就是是私有的


### 样例3，点（.）操作，调用该包中函数时可以省略前缀包名
```
import . "package1"

import (
    . "package1"
    . "package2"
    . "package3"
    ...
)

如  import . "fmt"
调用： Println(1)
```

### 样例4，别名操作
```
import alias package

import (
    p1 "package1"
    p2 "package2"
    p3 "package3"
    ...
)

如：import sys "fmt"

调用：sys.Println(1)
```


### 样例5：下划线
* 下划线（_）操作的含义是：导入该包，但不导入整个包，而是执行该包中的init函数，因此无法通过包名来调用包中的其他函数。

