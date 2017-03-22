
```
Java中的方法返回写为return; 或者return obj; 
在Groovy的方法中return是可选的，参数前面的def可以忽略。

groovy中，void类型 返回值为null。
 groovy中，调用方法时候，括号可以省略。

定义方法如下：
def function(arg){

}

如果要将方法设置为static类型，则 static def function(arg){}

返回值可以与java混写：
String function(arg){
}

方法中可以设置 参数默认值：
static void repeat(var, x = 10) {}
```