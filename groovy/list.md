

```
1.定义list：
def numbers = [11,12,13,14]

2.数字列表：  i..j，表示 [i,i+1,,j]，  但是次类列表 不可变，无法修改值。

3.获取元素：
numbers[i]   获取第i位置，i从0开始
 
numbers[i,j]   获取列表，[ numbers[i], numbers[j] ]

numbers[i..j]  获取列表,  [ numbers[i]-- numbers[j] ]

4.列表长度：
numbers.size()

5.列表合并：
numbers = numbers + [16,17]

6.元素操作：
numbers.add(n)，   //添加元素到末尾
numbers.add(index,n)  // 在index位置之前的插入元素n。

numbers[i]=n， 修改元素， 如果i大于numbers.size()，则numbers自动扩容到i+1长度，中间的值为null。

groovy支持负索引：
println collect[-1]      //索引其倒数第1个元素
println collect[-2]      //索引其倒数第2个元素

Collection支持集合运算：
collect=collect+5        //在集合中添加元素5
collect=collect-'a'         //在集合中减去元素a(全部a)

7.join：
numbers.join(",")

8.
[2,5,7].pop()              //Result: 7 
[2,5,7].plus([3,6])        //Result: [2, 5, 7, 3, 6] 
[2,5,7,2].minus(2)         //Result: [5, 7] 
[2,5,7].remove(1)          //Result: 5; list = [2, 7] 
[2,7,5].reverse()          //Result: [5, 7, 2] 
[2,7,5].sort()             //Result: [2, 5, 7] 

9.
def lt = []
lt << 1 // lt.add()
lt << 2
lt.each {it-> println it}


```